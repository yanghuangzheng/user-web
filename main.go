package main
import(

	"user-web/initialize"

	"go.uber.org/zap"
)
func main()  {
    initialize.InitLogger()//zap初始化	
	router:=initialize.Routers()//初始化Router
	initialize.InitSrvConnng()//初始化连接
	if err:=initialize.Inittrans("zh");err!=nil{
		panic(err)
	}//初始化翻译
	zap.S().Info("启动服务器，端口:8888")
	if err:=router.Run("127.0.0.1:8888");err!=nil{
		zap.S().Panic("启动失败",err.Error())
	}else{
		zap.S().Debug("【监视成功:127.0.0.1:8888】")
	}
}