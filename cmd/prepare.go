package cmd

import (
	"fmt"
	"github.com/sachaos/ac-deck/lib/environment"
	"github.com/sachaos/ac-deck/lib/preparer"
	"github.com/sachaos/ac-deck/lib/atcoder"
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
		templatePath := viper.GetString("template")

		contestId := args[0]

		if !validateLanguage(language) {
			fmt.Println("Please specify supported language. Refer `acd languages`.")
			return fmt.Errorf("invalid language")
		}

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

		var selector *environment.EnvironmentSelector
		if contest.LangVer == atcoder.LangOld {
			selector = environment.DefaultOldEnvironmentSelector
		} else {
			selector = environment.DefaultEnvironmentSelector
		}

		return preparer.Prepare(contest, ".", selector.Select(language), templatePath)
	},
}

func init() {
	rootCmd.AddCommand(prepareCmd)
	prepareCmd.Flags().StringP("language", "l", "", "language")
	viper.BindPFlag("language", prepareCmd.Flags().Lookup("language"))

	prepareCmd.Flags().StringP("template", "t", "", "template file path")
	viper.BindPFlag("template", prepareCmd.Flags().Lookup("template"))
}
