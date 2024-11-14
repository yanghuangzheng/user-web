package router
import(

  "user-web/model"
  "user-web/middlewares"

  "github.com/gin-gonic/gin"
  "go.uber.org/zap"
)

func InitUserRouter(Router *gin.RouterGroup){
	//给用户设置路由
     UserRouter:=Router.Group("/user")
	 zap.S().Info("配置用户相关的url")
	 {
		UserRouter.GET("list",middlewares.JWTlc(),middlewares.CheckRole(),model.GetUserList)
		 UserRouter.POST("log",model.Logger)
		 UserRouter.POST("id",middlewares.JWTlc(),model.SelectById)
		 UserRouter.POST("id",model.CreateUser)
		 UserRouter.POST("update",middlewares.JWTlc(),model.UpdateUser)
		 UserRouter.POST("mo",middlewares.JWTlc(),model.MobileInfo)	 
	 }
}