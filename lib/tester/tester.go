package tester

import (
	"context"
	"fmt"
	"github.com/sachaos/atcoder/lib/atcoder"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/docker/docker/client"
	"github.com/gookit/color"
	"github.com/sirupsen/logrus"

	"github.com/sachaos/atcoder/lib/files"
)

type Result struct {
	Actual io.ReadWriter
	Log    io.ReadWriter
}

type Tester interface {
	Run(ctx context.Context, index int, example *atcoder.Example) (*Result, error)
	Clean(ctx context.Context) error
}

func RunTest(dir string, onContainer bool, timeout int) (bool, error) {
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
		result, err := tester.Run(ctx, index, example)
		if err != nil {
			return false, err
		}
		end := time.Now()

		ok, err := judgeResult(index, example, result, end.Sub(start))
		if err != nil {
			return false, err
		}

		if !ok {
			all = false
		}
	}

	return all, nil
}

func judgeResult(index int, example *atcoder.Example, result *Result, duration time.Duration) (bool, error) {
	actual, err := ioutil.ReadAll(result.Actual)
	if err != nil {
		return false, err
	}

	actualStr := strings.TrimSpace(string(actual))

	fmt.Printf("\n")
	fmt.Printf(color.Bold.Sprintf("Case %d: ", index+1))
	passed := actualStr == example.Exp
	if passed {
		color.Green.Printf("AC\n")
	} else {
		color.Red.Printf("WA\n")
		fmt.Printf("Input:\n")
		fmt.Println(example.In)
		fmt.Printf("\nExpected:\n")
		fmt.Println(example.Exp)
		fmt.Printf("\nActually:\n")
		fmt.Println(actualStr)
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
