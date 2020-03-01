package cmd

import (
	"fmt"
	"github.com/sachaos/atcoder/lib/atcoder"
	"github.com/sachaos/atcoder/lib/environment"
	"github.com/sachaos/atcoder/lib/preparer"
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

		ac, err := atcoder.NewAtCoder()
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

		return preparer.Prepare(contest, ".", environment.DefaultEnvironmentSelector.Select(language))
	},
}

func init() {
	rootCmd.AddCommand(prepareCmd)
	prepareCmd.Flags().StringP("language", "l", "", "language")
	viper.BindPFlag("language", prepareCmd.Flags().Lookup("language"))
}
