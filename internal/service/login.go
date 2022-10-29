package service

import (
	"errors"
	"github.com/gogf/gf/util/gconv"
	"shequn1/foundation/middlewares"
	"shequn1/foundation/password"
	"shequn1/internal/entities"
)

func Auth(loginForm entities.LoginForm) (error, string) {
	err, user := GetUserPwd(loginForm.Username)
	if err != nil {
		return err, ""
	}
	if !password.Verify(loginForm.Password, user.Pwd) {
		return errors.New("password error"), ""
	}
	userAuth := entities.GetAuth()
	gconv.Struct(user, userAuth)
	token, err := middlewares.NewToken(userAuth)
	if err != nil {
		return err, ""
	}
	return nil, token
}
func GetUserPwd(indentifed string) (error, entities.User) {
	return nil, entities.User{}
}
