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
			w.Append([]string{status.SubmissionDate, status.Problem, status.Language, status.Point, status.CodeLength, status.Result, status.ElapsedTime, status.Memory})
		}
		w.Render()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
