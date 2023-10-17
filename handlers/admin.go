package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func GoToConsole(db *sqlx.DB, ctx *fiber.Ctx) error {
	session_id := ctx.Cookies("session_id", "")
	if session_id == "" {
		return ctx.Redirect("/admin/login")
	}
	session := new(Session)
	err := db.Get(
		session,
		"select * from sessions where id=$1",
		session_id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.Redirect("/admin/login")
		}
		//Redirect to error page
		return ctx.Redirect("/error")
	}
  return ctx.Render("home", fiber.Map{
    "username": session.Username,
    "location": "home",
  }, "layout")
}
