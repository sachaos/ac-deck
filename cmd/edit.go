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
	"github.com/sachaos/atcoder/files"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path"
	"strings"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit source code",
	Aliases: []string{"e"},
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := args[0]

		conf, err := files.LoadConf(dir)
		if err != nil {
			return err
		}

		filePath := path.Join(dir, conf.Environment.SrcName)

		return runEditor(filePath)
	},
}

func runEditor(filePath string) error {
	editor := os.Getenv("EDITOR")

	if editor == "" {
		editor = "vim"
	}

	splitted := strings.Split(editor, " ")
	cname := splitted[0]
	args := splitted[1:]
	args = append(args, filePath)

	cmd := exec.Command(cname, args[:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
