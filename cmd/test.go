package cmd

import (
	"github.com/sachaos/atcoder/lib/tester"
	"github.com/spf13/cobra"
)

var timeout int

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test DIRECTORY",
	Short: "Run test",
	Aliases: []string{"t"},
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := args[0]

		_, err := tester.RunTest(dir, !noDocker, timeout)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().BoolVar(&noDocker, "no-docker", false, "no docker")
	testCmd.Flags().IntVarP(&timeout, "timeout", "t", 3, "timeout (in second)")
}
