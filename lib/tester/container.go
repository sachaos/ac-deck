package tester

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"github.com/docker/cli/cli/streams"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/stdcopy"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/sachaos/ac-deck/lib/atcoder"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/sachaos/ac-deck/lib/files"
)

type ContainerTester struct {
	cli         *client.Client
	conf        *files.Conf
	dir         string
	containerId string
}

func NewContainerTester(ctx context.Context, cli *client.Client, conf *files.Conf, dir string) (*ContainerTester, error) {
	logrus.Debug("Start container")
	containerId, err := startContainer(ctx, cli, conf, dir)
	if err != nil {
		return nil, err
	}
	logrus.Debug("Finish start container")

	if conf.Environment.BuildCmdOnDocker != "" {
		logrus.Debug("Building binary")
		result, err := Exec(ctx, cli, containerId, strings.Split(conf.Environment.BuildCmdOnDocker, " "))
		if err != nil {
			return nil, err
		}

		if result.ExitCode != 0 {
			fmt.Println(result.Stdout)
			fmt.Println(result.Stderr)
			return nil, fmt.Errorf("exit %d", result.ExitCode)
		}
	}

	return &ContainerTester{
		cli:         cli,
		conf:        conf,
		dir:         dir,
		containerId: containerId,
	}, nil
}

func (t *ContainerTester) Run(ctx context.Context, stdin io.Reader, w io.Writer, ew io.Writer) error {
	contentBytes, err := ioutil.ReadAll(stdin)
	if err != nil {
		return err
	}
	content := string(contentBytes) + "\n"

	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	tw.WriteHeader(&tar.Header{
		Name: "example",
		Mode: 0666,
		Size: int64(len(content)),
	})
	tw.Write([]byte(content))
	tw.Close()

	err = t.cli.CopyToContainer(ctx, t.containerId, "/", &buf, types.CopyToContainerOptions{})
	if err != nil {
		return err
	}

	logrus.Debug("Run test")
	cmd := []string{"sh", "-c", fmt.Sprintf("cat /example | %s", t.conf.Environment.Cmd)}
	r, err := ExecWithStdin(ctx, t.cli, t.containerId, cmd, strings.NewReader(""))
	if err != nil {
		return err
	}

	io.Copy(w, r.Stdout)
	io.Copy(ew, r.Stderr)

	return nil
}

func (t *ContainerTester) Test(ctx context.Context, example *atcoder.Example) (*Result, error) {
	content := example.In + "\n"

	logrus.Debug("Running ContainerTester.Test")
	logrus.Debug("Copy input file to container")
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	tw.WriteHeader(&tar.Header{
		Name: "example",
		Mode: 0666,
		Size: int64(len(content)),
	})
	tw.Write([]byte(content))
	tw.Close()

	err := t.cli.CopyToContainer(ctx, t.containerId, "/", &buf, types.CopyToContainerOptions{})
	if err != nil {
		return nil, err
	}

	logrus.Debug("Run test")
	cmd := []string{"sh", "-c", fmt.Sprintf("cat /example | %s", t.conf.Environment.Cmd)}
	r, err := ExecWithStdin(ctx, t.cli, t.containerId, cmd, strings.NewReader(""))
	if err != nil {
		return nil, err
	}

	return &Result{
		Actual:   r.Stdout,
		Log:      r.Stderr,
		ExitCode: r.ExitCode,
	}, nil
}

func (t *ContainerTester) Clean(ctx context.Context) error {
	return t.cli.ContainerRemove(ctx, t.containerId, types.ContainerRemoveOptions{
		Force: true,
	})
}

func startContainer(ctx context.Context, cli *client.Client, conf *files.Conf, dir string) (string, error) {
	if conf.Environment.DockerImage == "" {
		return "", fmt.Errorf("empty DockerImage")
	}

	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	config := &container.Config{
		Tty:        true,
		Cmd:        []string{"/bin/sh"},
		Image:      conf.Environment.DockerImage,
		WorkingDir: conf.Environment.GetWorkingDir(),
	}
	hostConfig := &container.HostConfig{
		Binds: []string{path.Join(pwd, dir) + ":" + conf.Environment.GetSrcDir()},
	}

	err = PrepareImage(cli, ctx, conf.Environment.DockerImage)
	if err != nil {
		return "", err
	}

	networkingConfig := &network.NetworkingConfig{}
	create, err := cli.ContainerCreate(ctx, config, hostConfig, networkingConfig, &v1.Platform{
		Architecture: "amd64",
		OS:           "linux",
	}, "")
	if err != nil {
		return "", err
	}

	containerId := create.ID
	err = cli.ContainerStart(ctx, containerId, types.ContainerStartOptions{})
	if err != nil {
		return "", err
	}

	return containerId, nil
}

func PrepareImage(cli *client.Client, ctx context.Context, imageName string) error {
	_, _, err := cli.ImageInspectWithRaw(ctx, imageName)
	if client.IsErrNotFound(err) {
		fmt.Printf("Image not found: %s\n", imageName)
		pull, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
		if err != nil {
			return err
		}
		defer pull.Close()

		fmt.Printf("Pulling image: %s\n", imageName)
		out := streams.NewOut(os.Stdout)
		err = jsonmessage.DisplayJSONMessagesToStream(pull, out, nil)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

type ExecResult struct {
	Stdout   io.Reader
	Stderr   io.Reader
	ExitCode int
}

func Exec(ctx context.Context, cli *client.Client, name string, cmd []string) (*ExecResult, error) {
	return ExecWithStdin(ctx, cli, name, cmd, strings.NewReader(""))
}

// Ref: https://stackoverflow.com/questions/52774830/docker-exec-command-from-golang-api
func ExecWithStdin(ctx context.Context, cli *client.Client, name string, cmd []string, stdin io.Reader) (*ExecResult, error) {
	logrus.Infof("running: %+v", cmd)
	execConf := types.ExecConfig{
		Cmd:          cmd,
		AttachStdin:  true,
		AttachStderr: true,
		AttachStdout: true,
		Tty:          false,
		Detach:       false,
	}

	logrus.Debug("Running ContainerExecCreate")
	exec, err := cli.ContainerExecCreate(ctx, name, execConf)
	if err != nil {
		return nil, err
	}

	logrus.Debug("Running ContainerExecAttach")
	res, err := cli.ContainerExecAttach(ctx, exec.ID, types.ExecStartCheck{
		Tty: execConf.Tty,
	})
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var outBuf, errBuf bytes.Buffer
	outputDone := make(chan error)

	go func() {
		logrus.Debug("Copy from output")
		_, err = stdcopy.StdCopy(&outBuf, &errBuf, res.Reader)
		logrus.Debug("Copy end")
		outputDone <- err
	}()

	go func() {
		logrus.Debug("Sending input data")
		_, err = io.Copy(res.Conn, stdin)
		logrus.Debug("Finish sending input data")
	}()

	select {
	case err := <-outputDone:
		if err != nil {
			return nil, err
		}
		break

	case <-ctx.Done():
		return nil, ctx.Err()
	}

	inspect, err := cli.ContainerExecInspect(ctx, exec.ID)
	if err != nil {
		return nil, err
	}

	return &ExecResult{
		Stdout:   &outBuf,
		Stderr:   &errBuf,
		ExitCode: inspect.ExitCode,
	}, nil
}
