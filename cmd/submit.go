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
	"github.com/sachaos/atcoder/tester"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path"
)

var skipTest bool
var noDocker bool

// submitCmd represents the submit command
var submitCmd = &cobra.Command{
	Use:   "submit DIRECTORY",
	Short: "Submit to AtCoder",
	Aliases: []string{"s"},
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		username := viper.GetString("username")
		password := viper.GetString("password")

		dir := args[0]

		if !skipTest {
			allPassed, err := tester.RunTest(dir, !noDocker)
			if err != nil {
				return err
			}

			fmt.Println()

			if !allPassed {
				fmt.Printf("Submit was canceled because test failed. Please use --skip-test if you want to submit anyway.")
				return nil
			}
		}

		ac, err := lib.NewAtCoder()
		if err != nil {
			return err
		}

		conf, err := files.LoadConf(dir)
		if err != nil {
			return err
		}

		fpath := path.Join(dir, conf.Environment.SrcName)
		file, err := os.Open(fpath)
		if err != nil {
			return err
		}
		defer file.Close()

		err = ac.Login(username, password)
		if err != nil {
			return err
		}

		err = ac.Submit(conf.AtCoder.ContestID, conf.AtCoder.TaskID, conf.Environment.LanguageCode, file)
		if err != nil {
			return err
		}

		fmt.Println("Submit success!")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(submitCmd)
	submitCmd.Flags().BoolVarP(&skipTest, "skip-test", "s", false, "skip test")
	submitCmd.Flags().BoolVar(&noDocker, "no-docker", false, "no docker")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// submitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// submitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
