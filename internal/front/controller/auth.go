package controller

import (
	"github.com/gin-gonic/gin"
	"shequn1/foundation/app"
	"shequn1/foundation/validator"
	"shequn1/internal/entities"
	"shequn1/internal/service"
)

// Login 登录示例
func Login(ctx *gin.Context) {
	var loginForm entities.LoginForm
	if err := validator.Bind(ctx, &loginForm); !err.IsValid() {
		app.NewResponse(app.Success, err.ErrorsInfo).End(ctx)
		return
	}
	err, token := service.Auth(loginForm)
	if err != nil {
		app.NewResponse(app.Success, err.Error()).End(ctx)
		return
	}
	app.NewResponse(app.Success, gin.H{"token": token}).End(ctx)
}
