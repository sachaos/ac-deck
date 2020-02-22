package cmd

import (
	"github.com/sachaos/atcoder/tester"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test DIRECTORY",
	Short: "Run test",
	Aliases: []string{"t"},
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := args[0]

		_, err := tester.RunTest(dir, !noDocker)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().BoolVar(&noDocker, "no-docker", false, "no docker")
}
