package tester

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gookit/color"
	"github.com/sachaos/atcoder/files"
	"github.com/sachaos/atcoder/lib"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func RunTest(dir string) error {
	conf, err := files.LoadConf(dir)
	if err != nil {
		return err
	}

	examples, err := files.LoadTestData(dir)
	if err != nil {
		return err
	}

	cmd := strings.Split(conf.Environment.Cmd, " ")

	for i, example := range examples {
		err := runTest(i, cmd, dir, example)
		if err != nil {
			return err
		}
	}

	return nil
}

// TODO: refactoring
func runTest(index int, args []string, dir string, c *lib.Example) error {
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
			return err
		}
		os.Stderr.Write(errOutput)
		return err
	}

	output, err := ioutil.ReadAll(out)
	if err != nil {
		return err
	}

	fmt.Printf("Case %d: ", index)
	if string(output) == c.Exp {
		color.Green.Printf("OK\n")
	} else {
		color.Red.Printf("NG\n")
		fmt.Printf("  Input:\n")
		fmt.Printf("    ")
		fmt.Println(c.In)
		fmt.Printf("  Expected:\n")
		fmt.Printf("    \"%s\"\n", c.Exp)
		fmt.Printf("  Actually:\n")
		fmt.Printf("    \"%s\"\n", string(output))
	}

	errOutput, err := ioutil.ReadAll(stderr)
	if err != nil {
		return err
	}
	if len(errOutput) != 0 {
		fmt.Printf("Log:\n")
		os.Stderr.Write(errOutput)
	}

	return nil
}
