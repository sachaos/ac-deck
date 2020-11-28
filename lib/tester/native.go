package tester

import (
	"bytes"
	"context"
	"fmt"
	"github.com/sachaos/ac-deck/lib/atcoder"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/sachaos/ac-deck/lib/files"
)

type NativeTester struct {
	dir  string
	conf *files.Conf
}

func NewNativeTester(dir string, conf *files.Conf) (*NativeTester, error) {
	if conf.Environment.BuildCmd != "" {
		logrus.Infof("running: %s", conf.Environment.BuildCmd)
		buildCmd := strings.Split(conf.Environment.BuildCmd, " ")
		cmd := exec.Command(buildCmd[0], buildCmd[1:]...)
		stdErr := bytes.Buffer{}
		cmd.Dir = dir
		cmd.Stderr = &stdErr
		err := cmd.Run()
		if err != nil {
			logrus.Debugf("stderr: %s", stdErr.String())
			return nil, fmt.Errorf("build cmd: %w", err)
		}
	}

	return &NativeTester{dir: dir, conf: conf}, nil
}

func (t *NativeTester) Run(ctx context.Context, r io.Reader, w io.Writer, ew io.Writer) error {
	panic("run is not implemented in Native mode")
}

func (t *NativeTester) Test(ctx context.Context, index int, example *atcoder.Example) (*Result, error) {
	cmd := strings.Split(t.conf.Environment.Cmd, " ")

	result, err := runTestOnNative(ctx, cmd, t.dir, example)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *NativeTester) Clean(ctx context.Context) error {
	if t.conf.Environment.CleanCmd == "" {
		return nil
	}

	logrus.Infof("running: %s", t.conf.Environment.CleanCmd)
	cleanCmd := strings.Split(t.conf.Environment.CleanCmd, " ")
	cmd := exec.Command(cleanCmd[0], cleanCmd[1:]...)
	stdErr := bytes.Buffer{}
	cmd.Dir = t.dir
	cmd.Stderr = &stdErr

	err := cmd.Run()
	if err != nil {
		logrus.Debugf("stderr: %s", stdErr.String())
		return err
	}

	return nil
}

func runTestOnNative(ctx context.Context, args []string, dir string, c *atcoder.Example) (*Result, error) {
	var buf bytes.Buffer
	var ebuf bytes.Buffer

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)
	cmd.Dir = dir
	cmd.Stdin = strings.NewReader(c.In)
	cmd.Stdout = &buf
	cmd.Stderr = &ebuf

	r := Result{
		Actual:   &buf,
		Log:      &ebuf,
		ExitCode: 0,
	}

	err := cmd.Run()
	if err != nil {
		io.Copy(os.Stderr, r.Log)
		io.Copy(os.Stderr, r.Actual)

		return nil, err
	}

	io.Copy(&buf, r.Actual)
	io.Copy(&ebuf, r.Log)

	r.ExitCode = cmd.ProcessState.ExitCode()

	return &r, nil
}
