package cmd

import (
	"restdis/db"

	"github.com/gofiber/template/html/v2"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

var launchCmd = &cobra.Command{
	Use:          "launch",
	Short:        "Lanches the Restdis server on port 8080 ",
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		engine := html.New("/home/thewisepigeon/code/restdis/views", ".html")
		app := fiber.New(fiber.Config{Views: engine})

    db := db.ConnectToDB()

		app.Get("/admin/login", func(c *fiber.Ctx) error {
			return c.Render("login", fiber.Map{})
		})

		app.Listen(":8080")
	},
}

func init() {
	rootCmd.AddCommand(launchCmd)
}
