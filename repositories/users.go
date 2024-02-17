package repositories

import (
	"database/sql"
	"fmt"
	"github.com/blockloop/scan/v2"
	"restdis/types"
)

type Users struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *Users {
	return &Users{
		db: db,
	}
}

func (u *Users) Insert(newUser *types.User) error {
	_, err := u.db.Exec(
		`
      insert into users (
        username, 
        password, 
        can_manage_users, 
        can_manage_tokens, 
        is_active
      )
      values (
        ?, ?, ?, ?, ?
      )
    `,
		newUser.Username,
		newUser.Password,
		newUser.CanManageUsers,
		newUser.CanManageTokens,
		newUser.IsActive,
	)
	if err != nil {
		return fmt.Errorf("Error while inserting new user: %w", err)
	}
	return nil
}

func (u *Users) GetByUsername(username string) (*types.User, error) {
	rows, err := u.db.Query("select * from users where username=? limit 1", username)
	if err != nil {
		return nil, fmt.Errorf("Error while querying db for user: %w", err)
	}
	defer rows.Close()
	dbUser := new(types.User)
	err = scan.Row(dbUser, rows)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.ErrUserNotFound
		}
		return nil, fmt.Errorf("Error while scanning rows: %w", err)
	}
	return dbUser, nil
}

func (u *Users) GetAll() (*[]types.User, error) {
	return nil, nil
}

func (u *Users) Delete() error {
	return nil
}

func (u *Users) ToggleActive() error {
	return nil
}
