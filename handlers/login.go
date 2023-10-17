package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username" form:"username" db:"username"`
	Password string `json:"Password" form:"password" db:"password"`
}

type Session struct {
	Id       string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}

func Login(db *sqlx.DB, ctx *fiber.Ctx) error {
	payload := new(User)
	err := ctx.BodyParser(payload)
	if err != nil {
		println(err.Error())
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
	u := new(User)
	err = db.Get(
		u,
		`select username, password from users where username=$1`,
		payload.Username,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.Render("login", fiber.Map{
				"BadRequest": true,
			})
		}
		return ctx.Render("login", fiber.Map{
			"InternalError": true,
		})
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(payload.Password))
	if err != nil {
		return ctx.Render("login", fiber.Map{
			"BadRequest": true,
		})
	}
	session_id := uuid.NewString()
	_, err = db.NamedExec(
		`insert into sessions(id, username) values(:id, :username)`,
		Session{
			Id:       session_id,
			Username: u.Username,
		},
	)
	if err != nil {
		return ctx.Render("login", fiber.Map{
			"InternalError": true,
		})
	}
  cookie := new(fiber.Cookie)
  cookie.Name = "session_id"
  cookie.Value = session_id
  cookie.Path = "/"
  ctx.Cookie(cookie)
  return ctx.Render("console", fiber.Map{})
}
