package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"Password"`
}

func Login(db *sqlx.DB, ctx *fiber.Ctx) (err error) {
	u := new(UserInfo)
	if err = ctx.BodyParser(u); err != nil {
		return ctx.SendStatus(fiber.ErrBadRequest.Code)
	}
	return 
}
