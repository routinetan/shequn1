package controller

import (
	"github.com/gin-gonic/gin"
	"shequn1/internal/biz"
	app2 "shequn1/internal/foundation/app"
	"shequn1/internal/foundation/validator"
	"shequn1/internal/store/entities"
)

// Login 登录示例
func Login(ctx *gin.Context) {
	var loginForm entities.LoginForm
	if err := validator.Bind(ctx, &loginForm); !err.IsValid() {
		app2.NewResponse(app2.Success, err.ErrorsInfo).End(ctx)
		return
	}
	err, token := biz.Auth(loginForm)
	if err != nil {
		app2.NewResponse(app2.Success, err.Error()).End(ctx)
		return
	}
	app2.NewResponse(app2.Success, gin.H{"token": token}).End(ctx)
}
