package model
import(
    "net/http"
	"strings"
     

     "user-web/global"

    "github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/go-playground/validator/v10"
)

func removetop(fie map[string]string)map[string]string{
	rsp :=map[string]string{}
	for fies,err:=range fie{
		rsp[fies[strings.Index(fies,".")+1:]]=err
	}
	return rsp
}

//错误处理
func HandleValidatorError(ctx *gin.Context,err error){
    ctx.JSON(http.StatusBadRequest, gin.H{"error": removetop(err.(validator.ValidationErrors).Translate(global.Trans))})       
    }
	// 将 gRPC 错误转换为 HTTP 响应
func HandleGrpcErrortoHttp(err error, c *gin.Context) {
    if err != nil {
        if e, ok := status.FromError(err); ok {
            switch e.Code() {
            case codes.NotFound:
                c.JSON(http.StatusNotFound, gin.H{
                    "msg": e.Message(),
                })
            case codes.Internal:
                c.JSON(http.StatusInternalServerError, gin.H{
                    "msg": "内部错误",
                })
            case codes.InvalidArgument:
                c.JSON(http.StatusBadRequest, gin.H{
                    "msg": "参数错误",
                })
			case codes.Unavailable:
                c.JSON(http.StatusBadRequest, gin.H{
                    "msg": "用户服务不可用",
                })
            default:
                c.JSON(http.StatusInternalServerError, gin.H{
                    "msg": "其他错误"+e.Message(),
                })
            }
            return
        }
    }
    c.JSON(http.StatusInternalServerError, gin.H{
        "msg": "未知错误",
    })
}