package tester

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/docker/docker/client"
	"github.com/gookit/color"

	"github.com/sachaos/atcoder/files"
	"github.com/sachaos/atcoder/lib"
)

type Result struct {
	Actual io.ReadWriter
	Log    io.ReadWriter
}


type Tester interface {
	Run(ctx context.Context, index int, example *lib.Example) (*Result, error)
	Clean(ctx context.Context) error
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
		ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
		result, err := tester.Run(ctx, index, example)
		if err != nil {
			return false, err
		}

		ok, err := judgeResult(index, example, result)
		if err != nil {
			return false, err
		}

		if !ok {
			all = false
		}
	}

	return all, nil
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
