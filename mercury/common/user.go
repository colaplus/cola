package common

const (
	UserSexMan = 1
	UserSexWomen = 2
)

type UserInfo struct {
	UserId   uint64 `json:"user_id" db:"user_id"`
	Username string `json:"user" db:"username"`
	Nickname string `json:"nickname" db:"nickname"`
	Sex      int    `json:"sex" db:"sex"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
