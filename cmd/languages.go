/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// languagesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// languagesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
