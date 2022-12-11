package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shequn1/internal/entities"
	app2 "shequn1/internal/foundation/app"
	"shequn1/internal/foundation/middlewares"
	"shequn1/internal/foundation/rbac"
)

// CheckPermission 验证用户权限, 这是自定义的结构体，这里只是示例
func CheckPermission(context *gin.Context) {
	staff, exists := context.Get(middlewares.AuthKey)
	if !exists {
		app2.NewResponse(app2.PermissionDenied, nil, app2.PermissionDeniedMessage).End(context, http.StatusForbidden)
		return
	}

	if rbac.HasPermission(staff.(entities.Staff).ID.Hex(), context) {
		context.Next()
		return
	}

	app2.NewResponse(app2.PermissionDenied, nil, app2.PermissionDeniedMessage).End(context, http.StatusForbidden)
	context.Abort()
	return
}
