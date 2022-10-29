package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"shequn1/foundation/app"
	"shequn1/foundation/database/mongo"
	"shequn1/foundation/middlewares"
	"shequn1/foundation/password"
	"shequn1/foundation/validator"
	"shequn1/internal/entities"
)

// LoginForm 登录表单验证结构体
type LoginForm struct {
	Username string `binding:"required,max=12" form:"username"`
	Password string `binding:"required,max=128" form:"password"`
}

// Login 登录示例
func Login(ctx *gin.Context) {
	var loginForm LoginForm
	if err := validator.Bind(ctx, &loginForm); !err.IsValid() {
		app.NewResponse(app.Success, err.ErrorsInfo).End(ctx)
		return
	}

	var staff entities.Staff
	notFound := mongo.Collection(staff).Where(bson.M{"username": loginForm.Username}).FindOne(&staff)
	if notFound != nil {
		app.NewResponse(app.Success, nil, notFound.Error()).End(ctx)
		return
	}

	if !password.Verify(loginForm.Password, staff.Password) {
		app.NewResponse(app.Success, nil, "password error").End(ctx)
		return
	}
	staff.Logged(ctx.DefaultQuery("platform", "web"))
	token, err := middlewares.NewToken(staff)
	if err != nil {
		app.NewResponse(app.Success, nil, err.Error()).End(ctx)
		return
	}

	app.NewResponse(app.Success, gin.H{"token": token}).End(ctx)
}

// StaffInfo 获取当前登录用户信息
func StaffInfo(ctx *gin.Context) {
	staff, exists := ctx.Get(middlewares.AuthKey)
	if exists {
		app.NewResponse(app.Success, staff).End(ctx)
		return
	}
	app.NewResponse(app.Success, nil).End(ctx)
}
