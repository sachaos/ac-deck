package tester

import (
	"bytes"
	"context"
	"fmt"
	"github.com/docker/docker/client"
	"github.com/gookit/color"
	"github.com/sachaos/atcoder/files"
	"github.com/sachaos/atcoder/lib"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Result struct {
	Actual io.ReadWriter
	Log    io.ReadWriter
}

func RunTest(dir string, onContainer bool) (bool, error) {
	conf, err := files.LoadConf(dir)
	if err != nil {
		return false, err
	}

	examples, err := files.LoadTestData(dir)
	if err != nil {
		return false, err
	}

	results := make([]*Result, len(examples))

	err = test(dir, conf, examples, results, onContainer)
	if err != nil {
		return false, err
	}

	all := true
	for i := range results {
		ok, err := judgeResult(i, examples[i], results[i])
		if err != nil {
			return false, err
		}

		if !ok {
			all = false
		}
	}

	return all, nil
}

// test attempt container mode first, and fallback to native mode
func test(dir string, conf *files.Conf, examples []*lib.Example, results []*Result, onContainer bool) error {
	for i := range results {
		results[i] = &Result{
			Actual: &bytes.Buffer{},
			Log:    &bytes.Buffer{},
		}
	}

	cli, err := client.NewEnvClient()
	logrus.Infof("onContainer: %s", onContainer)
	if err == nil && onContainer {
		fmt.Println("Running test in container mode")
		defer cli.Close()
		err := testOnContainer(cli, dir, conf, examples, results)
		if err == nil {
			return nil
		}

		fmt.Printf("Failed to run in container mode: %s\n", err)
	}

	fmt.Println("Running test in native mode")
	return testOnNative(dir, conf, examples, results)
}

func judgeResult(index int, example *lib.Example, result *Result) (bool, error) {
	actual, err := ioutil.ReadAll(result.Actual)
	if err != nil {
		return false, err
	}

	actualStr := strings.TrimSpace(string(actual))

	fmt.Printf("Case %d: ", index)
	passed := actualStr == example.Exp
	if passed {
		color.Green.Printf("AC\n")
	} else {
		color.Red.Printf("WA\n")
		fmt.Printf("  Input:\n")
		fmt.Printf("    ")
		fmt.Println(example.In)
		fmt.Printf("  Expected:\n")
		fmt.Printf("    \"%s\"\n", example.Exp)
		fmt.Printf("  Actually:\n")
		fmt.Printf("    \"%s\"\n", actualStr)
	}

	errOutput, err := ioutil.ReadAll(result.Log)
	if err != nil {
		return false, err
	}
	if len(errOutput) != 0 {
		fmt.Printf("Log:\n")
		os.Stderr.Write(errOutput)
	}

	return true, nil
}

func testOnNative(dir string, conf *files.Conf, examples []*lib.Example, results []*Result) error {
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
			return fmt.Errorf("build cmd: %w", err)
		}
	}

	cmd := strings.Split(conf.Environment.Cmd, " ")

	for i := range examples {
		err := runTestOnNative(i, cmd, dir, examples[i], results[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func runTestOnNative(index int, args []string, dir string, c *lib.Example, r *Result) error {
	cmd := exec.CommandContext(context.Background(), args[0], args[1:]...)
	cmd.Dir = dir
	cmd.Stdin = strings.NewReader(c.In)
	cmd.Stdout = r.Actual
	cmd.Stderr = r.Log

	err := cmd.Run()
	if err != nil {
		errOutput, err := ioutil.ReadAll(r.Log)
		if err != nil {
			return err
		}
		os.Stderr.Write(errOutput)
		return err
	}

	return nil
}
