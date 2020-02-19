package tester

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/sirupsen/logrus"

	"github.com/sachaos/atcoder/files"
	"github.com/sachaos/atcoder/lib"
)

func testOnContainer(cli *client.Client, dir string, conf *files.Conf, examples []*lib.Example, results []*Result) error {
	if conf.Environment.DockerImage == "" {
		return fmt.Errorf("empty DockerImage")
	}

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	config := &container.Config{
		Tty:        true,
		Cmd:        []string{"/bin/sh"},
		Image:      conf.Environment.DockerImage,
		WorkingDir: "/src",
	}
	hostConfig := &container.HostConfig{
		Binds: []string{path.Join(pwd, dir) + ":/src"},
	}
	networkingConfig := &network.NetworkingConfig{}
	pull, err := cli.ImagePull(context.Background(), conf.Environment.DockerImage, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	io.Copy(ioutil.Discard, pull)
	defer pull.Close()

	create, err := cli.ContainerCreate(context.Background(), config, hostConfig, networkingConfig, "")
	if err != nil {
		return err
	}

	containerId := create.ID
	err = cli.ContainerStart(context.Background(), containerId, types.ContainerStartOptions{})
	if err != nil {
		return err
	}

	defer func() {
		err := cli.ContainerRemove(context.Background(), containerId, types.ContainerRemoveOptions{
			Force:         true,
		})
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	if conf.Environment.BuildCmdOnDocker != "" {
		result, err := Exec(context.Background(), cli, containerId, strings.Split(conf.Environment.BuildCmdOnDocker, " "))
		if err != nil {
			return err
		}

		if result.ExitCode != 0 {
			fmt.Println(result.Stdout)
			fmt.Println(result.Stderr)
			return fmt.Errorf("exit %d", result.ExitCode)
		}
	}

	for i := range examples {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		r, err := ExecWithStdin(ctx, cli, containerId, strings.Split(conf.Environment.Cmd, " "), examples[i].In)
		if err != nil {
			return err
		}

		results[i].Actual = bytes.NewBufferString(r.Stdout)
		results[i].Log = bytes.NewBufferString(r.Stderr)
	}

	return nil
}

type ExecResult struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

func Exec(ctx context.Context, cli *client.Client, name string, cmd []string) (*ExecResult, error) {
	return ExecWithStdin(ctx, cli, name, cmd, "")
}

// Ref: https://stackoverflow.com/questions/52774830/docker-exec-command-from-golang-api
func ExecWithStdin(ctx context.Context, cli *client.Client, name string, cmd []string, stdin string) (*ExecResult, error) {
	logrus.Infof("running: %+v", cmd)
	execConf := types.ExecConfig{
		Cmd: cmd,
		AttachStdin:  true,
		AttachStderr: true,
		AttachStdout: true,
		Tty: false,
		Detach: false,
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

	logrus.Debug("Sending input data")
	if stdin != "" {
		_, err = res.Conn.Write([]byte(stdin + "\n"))
		if err != nil {
			return nil, err
		}
	}

	var outBuf, errBuf bytes.Buffer
	outputDone := make(chan error)

	go func() {
		logrus.Debug("Copy from output")
		_, err = stdcopy.StdCopy(&outBuf, &errBuf, res.Reader)
		logrus.Debug("Copy end")
		outputDone <- err
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

	logrus.Debug("Load from outBuf")
	stdout, err := ioutil.ReadAll(&outBuf)
	if err != nil {
		return nil, err
	}
	logrus.Debug("Load from errBuf")
	stderr, err := ioutil.ReadAll(&errBuf)
	if err != nil {
		return nil, err
	}

	inspect, err := cli.ContainerExecInspect(ctx, exec.ID)
	if err != nil {
		return nil, err
	}

	logrus.Debug(string(stdout))
	logrus.Debug(string(stderr))

	return &ExecResult{
		Stdout:   string(stdout),
		Stderr:   string(stderr),
		ExitCode: inspect.ExitCode,
	}, nil
}
