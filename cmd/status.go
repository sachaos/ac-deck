package cmd

import (
	"github.com/sachaos/ac-deck/lib/atcoder"
	"github.com/sachaos/ac-deck/lib/status"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

		status := status.NewStatus(ac)
		err = status.Render(contestId)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
