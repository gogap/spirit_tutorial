package controllers

import (
	"github.com/gogap/errors"
)

var (
	ErrUserNotExist  = errors.TN(TodoErrorNamespace, 40000, "user not exist: {{.username}}")
	ErrWrongPassword = errors.TN(TodoErrorNamespace, 40001, "wrong password !")
)

type Auth struct {
	users map[string]string
}

func NewAuth() *Auth {
	users := make(map[string]string)
	users["zeal"] = "123456"
	users["gogap"] = "123456"

	return &Auth{
		users: users,
	}
}

func (p *Auth) CheckAuth(username, password string) (err error) {
	if pwd, exist := p.users[username]; !exist {
		err = ErrUserNotExist.New(errors.Params{"username": username})
		return
	} else if pwd != password {
		err = ErrWrongPassword.New()
		return
	}

	return
}
