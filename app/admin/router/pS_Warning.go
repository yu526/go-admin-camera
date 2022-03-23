package router
import (
	"go-admin/app/admin/apis"

	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerPS_WarningRouter)
}

// 需认证的路由代码
func registerPS_WarningRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api:= apis.RealGps{}
	r := v1.Group("")
	{
		r.GET("/realWarning", api.GetPS_Warning)
	}
}