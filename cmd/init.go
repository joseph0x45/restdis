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
	Short: "A brief description of your application",
	Long:  `The init command sets up your restdis server by creating a configuration file, a sqlite database.\n Use the -i flag for the interactive version`,
	Run: func(cmd *cobra.Command, args []string) {
		user_homedir, err := os.UserHomeDir()
		if err != nil {
			println("Something went wrong while setting up the database: ", err.Error())
			os.Exit(1)
		}
		if _, err = os.Stat(fmt.Sprintf("%s/restdis", user_homedir)); os.IsNotExist(err) {
      err = os.Mkdir(fmt.Sprintf("%s/restdis", user_homedir), os.ModePerm)

		}
		if err != nil {
			println("Something went wrong while setting up the database: ", err.Error())
			os.Exit(1)
		}
		if err != nil {
			println("Something went wrong while setting up the database: ", err.Error())
			os.Exit(1)
		}
		db_location := fmt.Sprintf("%s/restdis/data.db", user_homedir)
		if _, err := os.Stat(db_location); os.IsNotExist(err) {
			_, err = os.Create(db_location)
			if err != nil {
				println("Something went sd wrong while setting up the database: ", err.Error())
				os.Exit(1)
			}
		}
		println("Restdis database setup at ", db_location)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
