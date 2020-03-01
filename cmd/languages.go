package cmd

import (
	"github.com/olekukonko/tablewriter"
	"github.com/sachaos/atcoder/lib/environment"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// languagesCmd represents the languages command
var languagesCmd = &cobra.Command{
	Use:   "languages",
	Short: "list supported languages",
	Run: func(cmd *cobra.Command, args []string) {
		w := tablewriter.NewWriter(os.Stdout)
		w.SetHeader([]string{"key", "alias", "name", "image", "note"})
		selector := environment.DefaultEnvironmentSelector

		for _, key := range selector.Keys() {
			env := selector.Select(key)
			w.Append([]string{env.Key, strings.Join(selector.Aliases(key), ","), env.Language, env.DockerImage, env.Note})
		}
		w.Render()
	},
}

func init() {
	rootCmd.AddCommand(languagesCmd)
}
