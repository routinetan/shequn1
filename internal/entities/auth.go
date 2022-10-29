package entities

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	jsoniter "github.com/json-iterator/go"
	"shequn1/foundation/database/orm"
	"shequn1/foundation/middlewares"
	"shequn1/internal/model"
	"time"
)

type LoginForm struct {
	Username string `binding:"required,max=12" form:"username"`
	Password string `binding:"required,max=128" form:"password"`
}

type MysqlAuth struct {
	Id         int
	Indentifed string
	Pwd        string
	LoginAt    map[string]interface{}
}

func GetAuth() middlewares.AuthInterface {
	return &MysqlAuth{LoginAt: make(map[string]interface{})}
}

func (mysqlAuth *MysqlAuth) TableName() string {
	return "user"
}

func (mysqlAuth *MysqlAuth) GetUser() interface{} {
	return mysqlAuth
}

func (mysqlAuth *MysqlAuth) Find(topic interface{}) middlewares.AuthInterface {
	var user *model.User
	id, _ := topic.(int)
	orm.Master().Find(&user, id)
	gconv.Struct(user, &mysqlAuth)
	return mysqlAuth
}

func (mysqlAuth *MysqlAuth) GetCheckData() string {
	checkData, _ := jsoniter.MarshalToString(mysqlAuth.LoginAt)
	return checkData
}

func (mysqlAuth *MysqlAuth) Check(ctx *gin.Context, checkData string) bool {
	var loginAt map[string]int64
	_ = jsoniter.UnmarshalFromString(checkData, &loginAt)
	return mysqlAuth.LoginAt[ctx.DefaultQuery("platform", "web")] == loginAt[ctx.DefaultQuery("platform", "web")]

}

func (mysqlAuth MysqlAuth) ExpiredAt() int64 {
	return time.Now().Add(86400 * time.Second).Unix()
}
