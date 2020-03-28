package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var confName = ".ac-deck.toml"
var debugFlag bool
var version string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "acd",
	Short: "AC Deck: Unofficial CLI for AtCoder users",
	Version: version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	i := indexOfSubCommand(rootCmd, os.Args)
	args := os.Args

	if i >= 0 {
		subCmd := args[i]
		args = append(args[:i], args[i+1:]...)
		args = append(args, "")
		copy(args[2:], args[1:])
		args[1] = subCmd
	}

	rootCmd.SetArgs(args[1:])
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func indexOfSubCommand(cmd *cobra.Command, args []string) int {
	for i, arg := range args {
		for _, subCmd := range cmd.Commands() {
			if arg == subCmd.Name() {
				return i
			}

			for _, alias := range subCmd.Aliases {
				if arg == alias {
					return i
				}
			}
		}
	}

	return -1
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/%s)", confName))
	rootCmd.PersistentFlags().BoolVar(&debugFlag, "debug",  false, "Debug mode")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if debugFlag {
		logrus.SetOutput(os.Stderr)
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetOutput(ioutil.Discard)
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(confName)
	}

	viper.SetConfigType("toml")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil && !isConfigCommand(os.Args) {
		fmt.Println(err)
		fmt.Println("Please run `acd config`.")
		os.Exit(1)
	}
}

func isConfigCommand(args []string) bool {
	for _, arg := range args {
		if arg == "config" {
			return true
		}
	}

	return false
}

