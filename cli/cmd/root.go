package cmd

import (
	"fmt"
	"os"

	"dominguezdev.com/cli/tui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "ddn-gopher",
	Short: "A CLI tool which access a DDN-powered API",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			tui.RunTUI()
		} else {
			err := cmd.Help()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	// Basic config information
	viper.SetConfigName(".config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.ddn-gopher")
	viper.AutomaticEnv()

	// Attempting to read in an env file
	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())
}

func init() {
	cobra.OnInitialize(initConfig)
}
