package cmd

import (
	"fmt"

	"dominguezdev.com/cli/auth"
	"dominguezdev.com/cli/tui"
	"github.com/spf13/cobra"
)

// Top-level variables for our username and password being passed by the user
var (
	noPrompt bool
	username string
	password string
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with your username and password",
	Run: func(cmd *cobra.Command, args []string) {
		if noPrompt {
			err := auth.Login(username, password)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			tui.RunTUI()
		}
	},
}

func init() {
	// This AddCommand adds the loginCmd to our root set of commands.
	// rootCmd is available because this file and it are part of the same
	// cmd package.
	rootCmd.AddCommand(loginCmd)

	// This sets the flags our command uses and marks which are req
	loginCmd.Flags().BoolVarP(&noPrompt, "no-prompt", "n", false, "Use command-line flags for login")
	loginCmd.Flags().StringVarP(&username, "username", "u", "", "Username")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "Password")
	loginCmd.MarkFlagRequired("username")
	loginCmd.MarkFlagRequired("password")
}
