package cmd

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
	"github.com/sachaos/atcoder/files"
	"github.com/sachaos/atcoder/tester"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install [LANGUAGE_NAME]",
	Short: "Install language Docker image",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		language := args[0]

		if !validateLanguage(language) {
			fmt.Println("Please specify supported language. Refer `atcoder languages`.")
			return fmt.Errorf("invalid language")
		}

		environment := files.Environments[language]

		cli, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			return err
		}

		err = tester.PrepareImage(cli, context.Background(), environment.DockerImage)
		if err != nil {
			return err
		}

		fmt.Println("Preparation completed")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}