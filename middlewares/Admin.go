package middlewares

import(
	"net/http"

	"user-web/form"

	"github.com/gin-gonic/gin"
)
func CheckRole()gin.HandlerFunc{
	return func(ctx*gin.Context){
    claims,_:=ctx.Get("claim")
	User:=claims.(*form.CusClaims)
	if User.Role!=2{
		ctx.JSON(http.StatusForbidden,gin.H{
			"message":"无权限",
		})
		ctx.Abort()
		return
	}
	ctx.Next()
	}
}