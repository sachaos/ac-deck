package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path"
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
		fmt.Scanf("%s", &password)
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
