package status

import (
	"github.com/olekukonko/tablewriter"
	"github.com/sachaos/ac-deck/lib/atcoder"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

type Status struct {
	ac *atcoder.AtCoder
}

func NewStatus(ac *atcoder.AtCoder) *Status {
	return &Status{ac: ac}
}

func (s Status) WaitFor(contestId string) error {
	for {
		statuses, err := s.ac.Status(contestId)
		if err != nil {
			return err
		}

		result := statuses[0].Result
		log.Printf("result: %s", result)
		if strings.Index(result, "/") > 0 || result == "WJ" {
			time.Sleep(2 * time.Second)
		} else {
			return nil
		}
	}
}

func (s Status) Render(contestId string) error {
	statuses, err := s.ac.Status(contestId)
	if err != nil {
		return err
	}

	w := tablewriter.NewWriter(os.Stdout)
	w.SetHeader([]string{"submission date", "problem", "language", "score", "code length", "result", "elapsed time", "memory"})
	sort.Slice(statuses, func(i, j int) bool {
		return statuses[i].SubmissionDate < statuses[j].SubmissionDate
	})
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
