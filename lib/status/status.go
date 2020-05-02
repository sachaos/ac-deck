package status

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/olekukonko/tablewriter"
	"github.com/sachaos/ac-deck/lib/atcoder"
	"github.com/tj/go-spin"
	"os"
	"sort"
	"strings"
	"time"
)

type StatusMonitor struct {
	ac *atcoder.AtCoder
}

func NewStatus(ac *atcoder.AtCoder) *StatusMonitor {
	return &StatusMonitor{ac: ac}
}

func (s StatusMonitor) WaitFor(contestId string) error {
	done := make(chan struct{})
	resultC := make(chan string)
	go func() {
		s := spin.New()
		var result string
		for {
			select {
			case r := <-resultC:
				result = r
			case <-done:
				fmt.Printf("\r"+ color.Bold.Sprintf("Judging: ") + "%s DONE\n", result)
				return
			default:
				fmt.Printf("\r"+ color.Bold.Sprintf("Judging: ") + "%s %s", result, s.Next())
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	defer func() {
		done <- struct{}{}
	}()

	for {
		statuses, err := s.ac.Status(contestId)
		if err != nil {
			return err
		}

		result := statuses[0].Result
		resultC <- result
		if strings.Index(result, "/") > 0 || result == "WJ" {
			time.Sleep(2 * time.Second)
		} else {
			return nil
		}
	}
}

func (s StatusMonitor) Render(contestId string) error {
	statuses, err := s.ac.Status(contestId)
	if err != nil {
		return err
	}

	sort.Slice(statuses, func(i, j int) bool {
		return statuses[i].SubmissionDate < statuses[j].SubmissionDate
	})

	render(statuses)

	return nil
}

func render(statuses []*atcoder.Status) error {
	w := tablewriter.NewWriter(os.Stdout)
	w.SetHeader([]string{"submission date", "problem", "language", "score", "code length", "result", "elapsed time", "memory"})

	for _, status := range statuses {
		var resultColor tablewriter.Colors
		switch status.Result {
		case "AC":
			resultColor = tablewriter.Colors{tablewriter.FgHiGreenColor}
		case "WA":
			resultColor = tablewriter.Colors{tablewriter.FgHiRedColor}
		case "TLE":
			resultColor = tablewriter.Colors{tablewriter.FgHiYellowColor}
		case "RE":
			resultColor = tablewriter.Colors{tablewriter.FgHiYellowColor}
		case "CE":
			resultColor = tablewriter.Colors{tablewriter.FgHiYellowColor}
		case "OLE":
			resultColor = tablewriter.Colors{tablewriter.FgHiYellowColor}
		case "MLE":
			resultColor = tablewriter.Colors{tablewriter.FgHiYellowColor}
		case "IE":
			resultColor = tablewriter.Colors{tablewriter.FgHiYellowColor}
		default:
			resultColor = tablewriter.Colors{tablewriter.FgHiBlackColor}
		}
		colors := []tablewriter.Colors{{}, {}, {}, {}, {}, resultColor, {}, {}}
		w.Rich([]string{status.SubmissionDate, status.Problem, status.Language, status.Point, status.CodeLength, status.Result, status.ElapsedTime, status.Memory}, colors)
	}
	w.Render()
	return nil
}
