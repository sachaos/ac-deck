package cmd

import (
	"fmt"
	"os"
	"path"
	"syscall"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure atcoder-cli",
	RunE: func(cmd *cobra.Command, args []string) error {
		var username, password, language string
		fmt.Printf("username: ")
		fmt.Scanf("%s", &username)
		fmt.Printf("password: ")
		bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
		password = string(bytePassword)
		fmt.Printf("\n")
		fmt.Printf("language (list supported languages by `atcoder languages`): ")
		fmt.Scanf("%s", &language)

		viper.Set("username", username)
		viper.Set("password", password)
		viper.Set("language", language)

		if !validateLanguage(language) {
			fmt.Println("Please specify supported language. Refer `atcoder languages`.")
			return fmt.Errorf("invalid language")
		}

		dir, err := homedir.Dir()
		if err != nil {
			return err
		}

		file, err := os.OpenFile(path.Join(dir, confName), os.O_CREATE|os.O_RDWR, 0600)
		if err != nil {
			return err
		}
		file.Close()

		return viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
