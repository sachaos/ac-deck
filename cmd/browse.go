package cmd

import (
	"github.com/pkg/browser"
	"github.com/sachaos/atcoder/lib/files"
	"github.com/spf13/cobra"
)

// browseCmd represents the browse command
var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "browse AtCoder task page",
	Aliases: []string{"b"},
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := args[0]

		conf, err := files.LoadConf(dir)
		if err != nil {
			return err
		}

		return browser.OpenURL(conf.AtCoder.TaskURL)
	},
}

func init() {
	rootCmd.AddCommand(browseCmd)
}
