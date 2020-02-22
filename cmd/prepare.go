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
	"fmt"
	"github.com/sachaos/atcoder/files"
	"github.com/sachaos/atcoder/lib"
	"github.com/sachaos/atcoder/preparer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// prepareCmd represents the prepare command
var prepareCmd = &cobra.Command{
	Use:   "prepare CONTEST_ID",
	Short: "prepare for contest by fetching examples and generate source code from template",
	Aliases: []string{"p"},
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		username := viper.GetString("username")
		password := viper.GetString("password")
		language := viper.GetString("language")

		contestId := args[0]

		if !validateLanguage(language) {
			fmt.Println("Please specify supported language. Refer `atcoder languages`.")
			return fmt.Errorf("invalid language")
		}

		fmt.Printf("Using template of %s\n", language)

		ac, err := lib.NewAtCoder()
		if err != nil {
			return err
		}

		err = ac.Login(username, password)
		if err != nil {
			return err
		}

		contest, err := ac.FetchContest(contestId)
		if err != nil {
			return err
		}

		return preparer.Prepare(contest, ".", files.Environments[language])
	},
}

func init() {
	rootCmd.AddCommand(prepareCmd)
	prepareCmd.Flags().StringP("language", "l", "", "language")
	viper.BindPFlag("language", prepareCmd.Flags().Lookup("language"))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// prepareCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// prepareCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
