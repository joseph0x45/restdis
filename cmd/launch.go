package cmd

import (
	// "net/http"
	//
	// "github.com/go-chi/chi/v5"
	// "github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
)

var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Lanches the Restdis server on port 8080 ",
  SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(launchCmd)
}
