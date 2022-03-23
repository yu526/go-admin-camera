package router

import (
	"go-admin/app/admin/apis"

	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerPS_DeviceStatusRouter)
}

// 需认证的路由代码
func registerPS_DeviceStatusRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api:= apis.RealGps{}
	r := v1.Group("")
	{
		r.GET("/realGps", api.GetPS_DeviceStatus)
	}
}