package tester

import (
	"context"
	"fmt"
	"github.com/sachaos/ac-deck/lib/atcoder"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/docker/docker/client"
	"github.com/gookit/color"
	"github.com/sirupsen/logrus"

	"github.com/sachaos/ac-deck/lib/files"
)

type Result struct {
	Actual   io.Reader
	Log      io.Reader
	ExitCode int
}

type Tester interface {
	Run(ctx context.Context, r io.Reader, w io.Writer, ew io.Writer) error
	Test(ctx context.Context, example *atcoder.Example) (*Result, error)
	Clean(ctx context.Context) error
}

func Run(dir string, r io.Reader, w io.Writer, ew io.Writer) error {
	ctx := context.Background()
	conf, err := files.LoadConf(dir)
	if err != nil {
		return err
	}

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}

	tester, err := NewContainerTester(ctx, cli, conf, dir)
	if err != nil {
		return err
	}

	if err = tester.Run(ctx, r, w, ew); err != nil {
		return err
	}


	if err = tester.Clean(ctx); err != nil {
		return err
	}

	return nil
}

func RunTest(dir string, onContainer bool, timeout int, verbose bool) (bool, error) {
	conf, err := files.LoadConf(dir)
	if err != nil {
		return false, err
	}

	examples, err := files.LoadTestData(dir)
	if err != nil {
		return false, err
	}

	var tester Tester
	logrus.Debug("preparing Docker client")
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil || !onContainer {
		fmt.Println("Running test on Native mode")
		tester, err = NewNativeTester(dir, conf)
		if err != nil {
			return false, err
		}
	} else {
		fmt.Println("Running test on Container mode")
		tester, err = NewContainerTester(context.Background(), cli, conf, dir)
		if err != nil {
			return false, err
		}
	}

	defer tester.Clean(context.Background())

	all := true
	for index, example := range examples {
		ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		start := time.Now()
		result, err := tester.Test(ctx, example)
		if err != nil {
			return false, err
		}
		end := time.Now()

		ok, err := judgeResult(index, example, result, end.Sub(start), verbose)
		if err != nil {
			return false, err
		}

		if !ok {
			all = false
		}
	}

	return all, nil
}

func judgeResult(index int, example *atcoder.Example, result *Result, duration time.Duration, verbose bool) (bool, error) {
	actual, err := ioutil.ReadAll(result.Actual)
	if err != nil {
		return false, err
	}

	actualStr := strings.TrimSpace(string(actual))

	fmt.Printf("\n")
	fmt.Printf(color.Bold.Sprintf("Case %d: ", index+1))
	passed := judgeEquality(example.Exp, actualStr) && result.ExitCode == 0
	if passed {
		color.Green.Printf("AC\n")
	} else {
		color.Red.Printf("WA\n")
	}

	if !passed || verbose {
		fmt.Printf("Input:\n")
		fmt.Println(example.In)
		fmt.Printf("\nExpected:\n")
		fmt.Println(example.Exp)
		fmt.Printf("\nActually:\n")
		fmt.Println(actualStr)
		fmt.Printf("\nExit with: %d\n", result.ExitCode)
	}

	fmt.Printf("Time: %s\n", duration)

	errOutput, err := ioutil.ReadAll(result.Log)
	if err != nil {
		return false, err
	}
	if len(errOutput) != 0 {
		fmt.Printf("\nLog:\n")
		os.Stdout.Write(errOutput)
	}

	return passed, nil
}

func judgeEquality(example string, actual string) bool {
	if example == actual {
		return true
	}

	if !strings.Contains(example, ".") {
		return false
	}

	af, err := strconv.ParseFloat(actual, 64)
	if err != nil {
		return false
	}

	ef, err := strconv.ParseFloat(example, 64)
	if err != nil {
		return false
	}

	return (af - ef) / ef < 0.00001
}

