package cmd

import (
	"fmt"
	"github.com/sachaos/ac-deck/lib/atcoder"
	"github.com/sachaos/ac-deck/lib/files"
	"github.com/sachaos/ac-deck/lib/status"
	"github.com/sachaos/ac-deck/lib/tester"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path"
)

var skipTest bool
var noDocker bool

// submitCmd represents the submit command
var submitCmd = &cobra.Command{
	Use:     "submit DIRECTORY",
	Short:   "Submit to AtCoder",
	Aliases: []string{"s"},
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		username := viper.GetString("username")
		password := viper.GetString("password")

		dir := args[0]

		if !skipTest {
			allPassed, err := tester.RunTest(dir, !noDocker, timeout)
			if err != nil {
				return err
			}

			fmt.Println()

			if !allPassed {
				fmt.Printf("Submit was canceled because test failed. Please use --skip-test if you want to submit anyway.")
				return nil
			}
		}

		ac, err := atcoder.NewAtCoder()
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

		stat := status.NewStatus(ac)
		err = stat.WaitFor(conf.AtCoder.ContestID)
		if err != nil {
			return err
		}

		err = stat.Render(conf.AtCoder.ContestID)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(submitCmd)
	submitCmd.Flags().BoolVarP(&skipTest, "skip-test", "s", false, "skip test")
	submitCmd.Flags().BoolVar(&noDocker, "no-docker", false, "no docker")
	submitCmd.Flags().IntVarP(&timeout, "timeout", "t", 3, "timeout (in second)")
}
