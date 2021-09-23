package db

import (
	"database/sql"
	"mercury/common"
	"mercury/util"
)

const PasswordSalt = "asdfHOSDHFkjdaslkjfoOHIOISDF14654"

func Register(user *common.UserInfo) (err error) {
	var count int64
	sqlstr := "select count(user_id) from user where username=?"
	err = DB.Get(&count, sqlstr, user.Username)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	if count > 0 {
		err = ErrUserExists
		return
	}

	passwd := user.Password + PasswordSalt
	dbPassword := util.Md5([]byte(passwd))

	sqlstr = "insert into user(user_id, username, password, email, sex, nickname) values(?,?,?,?,?,?)"
	_, err = DB.Exec(sqlstr, user.UserId, user.Username, dbPassword, user.Email, user.Sex, user.Nickname)
	return
}

func Login(user *common.UserInfo) (err error) {
	originPassword := user.Password
	sqlstr := "select username,password,user_id from user where username=?"
	err = DB.Get(user, sqlstr, user.Username)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	if err == sql.ErrNoRows {
		err = ErrUserNotExists
		return
	}

	passwd := originPassword + PasswordSalt
	originPasswordSalt := util.Md5([]byte(passwd))
	if originPasswordSalt != user.Password {
		err = ErrUserPasswordWrong
	}

	return
}
