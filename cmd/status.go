package cmd

import (
	"github.com/olekukonko/tablewriter"
	"github.com/sachaos/ac-deck/lib/atcoder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"sort"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display status of submissions",
	RunE: func(cmd *cobra.Command, args []string) error {
		username := viper.GetString("username")
		password := viper.GetString("password")

		contestId := args[0]

		ac, err := atcoder.NewAtCoder()
		if err != nil {
			return err
		}

		err = ac.Login(username, password)
		if err != nil {
			return err
		}

		statuses, err := ac.Status(contestId)
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
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
