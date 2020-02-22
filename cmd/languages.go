package cmd

import (
	"os"
	"sort"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/sachaos/atcoder/files"
)

// languagesCmd represents the languages command
var languagesCmd = &cobra.Command{
	Use:   "languages",
	Short: "list supported languages",
	Run: func(cmd *cobra.Command, args []string) {
		environments := []*files.Environment{}
		for _, env := range files.Environments {
			environments = append(environments, env)
		}

		sort.Slice(environments, func(i, j int) bool {
			return environments[i].LanguageCode < environments[j].LanguageCode
		})

		w := tablewriter.NewWriter(os.Stdout)
		w.SetHeader([]string{"key", "name", "note"})
		for _, env := range environments {
			w.Append([]string{env.Key, env.Language, env.Note})
		}
		w.Render()
	},
}

func init() {
	rootCmd.AddCommand(languagesCmd)
}
