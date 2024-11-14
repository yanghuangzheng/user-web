package initialize

import (

    "github.com/gin-gonic/gin"
    userrouter "user-web/routerApi"
)

// 初始化路由
var ApiGroup *gin.RouterGroup

func Routers() *gin.Engine {
    // 创建默认的 Gin 引擎
    Router := gin.Default()
    //配置跨域
    //Router.Use(middlewares.Cors())
    // 创建 API 组
    ApiGroup = Router.Group("/user/v1")
    // 初始化用户路由
    userrouter.InitUserRouter(ApiGroup)
    userrouter.InitBaseRouter(ApiGroup)
    // 返回配置好的 Gin 引擎
    return Router
}