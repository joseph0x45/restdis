package types

type User struct {
	ID              string `db:"id"`
	Username        string `db:"username"`
	Password        string `db:"password"`
	CanManageUsers  bool   `db:"can_manage_users"`
	CanManageTokens bool   `db:"can_manage_tokens"`
	IsActive        bool   `db:"is_active"`
}
