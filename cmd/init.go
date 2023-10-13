/*
Copyright © 2023 github.com/TheWisePigeon <zozozozeph@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Sets up restdis sqlite database in the user's home directory",
  SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		user_homedir, err := os.UserHomeDir()
		db_location := fmt.Sprintf("%s/.restdis.db", user_homedir)
		if err != nil {
			println("Something went wrong : ", err.Error())
			os.Exit(1)
		}
		if _, err = os.Stat(db_location); os.IsNotExist(err) {
			_, err = os.Create(db_location)
			if err != nil {
				println("Something went wrong : ", err.Error())
				os.Exit(1)
			}
		}
		println("Database initialized at ", db_location)
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
