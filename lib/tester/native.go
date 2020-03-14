package tester

import (
	"bytes"
	"context"
	"fmt"
	"github.com/sachaos/atcoder/lib/atcoder"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/sachaos/atcoder/lib/files"
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

func (t *NativeTester) Run(ctx context.Context, index int, example *atcoder.Example) (*Result, error) {
	result := Result{
		Log: &bytes.Buffer{},
		Actual: &bytes.Buffer{},
	}

	cmd := strings.Split(t.conf.Environment.Cmd, " ")

	err := runTestOnNative(ctx, cmd, t.dir, example, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
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

func runTestOnNative(ctx context.Context, args []string, dir string, c *atcoder.Example, r *Result) error {
	cmd := exec.CommandContext(ctx, args[0], args[1:]...)
	cmd.Dir = dir
	cmd.Stdin = strings.NewReader(c.In)
	cmd.Stderr = r.Log
	cmd.Stdout = r.Actual

	err := cmd.Run()
	if err != nil {
		io.Copy(os.Stderr, r.Log)
		io.Copy(os.Stderr, r.Actual)
		return err
	}

	r.ExitCode = cmd.ProcessState.ExitCode()

	return nil
}
