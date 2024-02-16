package repositories

import (
	"database/sql"
	"fmt"
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

func (u *Users) GetAll() (*[]types.User, error) {
	return nil, nil
}

func (u *Users) Delete() error {
	return nil
}

func (u *Users) ToggleActive() error {
	return nil
}
