package cmd

import (
	"github.com/sachaos/ac-deck/lib/files"
	"path"

	"github.com/spf13/cobra"
)

// editTestCmd represents the edit-test command
var editTestCmd = &cobra.Command{
	Use:   "edit-test",
	Short: "Edit test data",
	Aliases: []string{"et"},
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := args[0]

		_, err := files.LoadConf(dir)
		if err != nil {
			return err
		}

		testData := path.Join(dir, files.TESTDATA_NAME)

		return runEditor(testData)
	},
}

func init() {
	rootCmd.AddCommand(editTestCmd)
}
