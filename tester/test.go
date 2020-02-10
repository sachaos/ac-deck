package tester

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gookit/color"
	"github.com/sachaos/atcoder/files"
	"github.com/sachaos/atcoder/lib"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func RunTest(dir string) (bool, error) {
	conf, err := files.LoadConf(dir)
	if err != nil {
		return false, err
	}

	examples, err := files.LoadTestData(dir)
	if err != nil {
		return false, err
	}

	defer func() {
		if conf.Environment.CleanCmd != "" {
			logrus.Infof("running: %s", conf.Environment.CleanCmd)
			cleanCmd := strings.Split(conf.Environment.CleanCmd, " ")
			cmd := exec.Command(cleanCmd[0], cleanCmd[1:]...)
			stdErr := bytes.Buffer{}
			cmd.Dir = dir
			cmd.Stderr = &stdErr
			err := cmd.Run()
			if err != nil {
				logrus.Debugf("stderr: %s", stdErr.String())
				fmt.Println(err)
			}
		}
	}()

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
			return false, fmt.Errorf("build cmd: %w", err)
		}
	}

	cmd := strings.Split(conf.Environment.Cmd, " ")

	allPassed := true
	for i, example := range examples {
		passed, err := runTest(i, cmd, dir, example)
		if err != nil {
			return false, err
		}
		if passed == false {
			allPassed = false
		}
	}

	return allPassed, nil
}

// TODO: refactoring
func runTest(index int, args []string, dir string, c *lib.Example) (bool, error) {
	out := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	cmd := exec.CommandContext(context.Background(), args[0], args[1:]...)
	cmd.Dir = dir
	cmd.Stdin = strings.NewReader(c.In)
	cmd.Stdout = out
	cmd.Stderr = stderr

	err := cmd.Run()
	if err != nil {
		errOutput, err := ioutil.ReadAll(stderr)
		if err != nil {
			return false, err
		}
		os.Stderr.Write(errOutput)
		return false, err
	}

	output, err := ioutil.ReadAll(out)
	if err != nil {
		return false, err
	}

	outputStr := strings.TrimSpace(string(output))

	fmt.Printf("Case %d: ", index)
	passed := outputStr == c.Exp
	if passed {
		color.Green.Printf("OK\n")
	} else {
		color.Red.Printf("NG\n")
		fmt.Printf("  Input:\n")
		fmt.Printf("    ")
		fmt.Println(c.In)
		fmt.Printf("  Expected:\n")
		fmt.Printf("    \"%s\"\n", c.Exp)
		fmt.Printf("  Actually:\n")
		fmt.Printf("    \"%s\"\n", outputStr)
	}

	errOutput, err := ioutil.ReadAll(stderr)
	if err != nil {
		return false, err
	}
	if len(errOutput) != 0 {
		fmt.Printf("Log:\n")
		os.Stderr.Write(errOutput)
	}

	return passed, nil
}
