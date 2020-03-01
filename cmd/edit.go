package cmd

import (
	"github.com/sachaos/atcoder/lib/files"
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
}
