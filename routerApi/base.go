package router
import(

    "user-web/model"
	"github.com/gin-gonic/gin"
	//"go.uber.org/zap"
)

func InitBaseRouter(Router *gin.RouterGroup){
	BaseRouter:=Router.Group("base")
	{
       BaseRouter.GET("captcha",model.GetCaptcha)
	   //BaseRouter.POST("send_sms",api.SendSms)
	}
     
}