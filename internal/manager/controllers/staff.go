package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	app2 "shequn1/internal/foundation/app"
	"shequn1/internal/foundation/database/mongo"
	"shequn1/internal/foundation/middlewares"
	"shequn1/internal/foundation/password"
	"shequn1/internal/foundation/validator"
	"shequn1/internal/store/entities"
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
		app2.NewResponse(app2.Success, err.ErrorsInfo).End(ctx)
		return
	}

	var staff entities.Staff
	notFound := mongo.Collection(staff).Where(bson.M{"username": loginForm.Username}).FindOne(&staff)
	if notFound != nil {
		app2.NewResponse(app2.Success, nil, notFound.Error()).End(ctx)
		return
	}

	if !password.Verify(loginForm.Password, staff.Password) {
		app2.NewResponse(app2.Success, nil, "password error").End(ctx)
		return
	}
	staff.Logged(ctx.DefaultQuery("platform", "view"))
	token, err := middlewares.NewToken(staff)
	if err != nil {
		app2.NewResponse(app2.Success, nil, err.Error()).End(ctx)
		return
	}

	app2.NewResponse(app2.Success, gin.H{"token": token}).End(ctx)
}

// StaffInfo 获取当前登录用户信息
func StaffInfo(ctx *gin.Context) {
	staff, exists := ctx.Get(middlewares.AuthKey)
	if exists {
		app2.NewResponse(app2.Success, staff).End(ctx)
		return
	}
	app2.NewResponse(app2.Success, nil).End(ctx)
}
