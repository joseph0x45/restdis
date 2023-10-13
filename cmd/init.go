package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
  "restdis/config"
)

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
    _, err = os.Create(db_location)
    if err != nil {
      println("Something went wrong : ", err.Error())
      os.Exit(1)
    }
		println("Database initialized at ", db_location)
    config.Init(db_location)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
