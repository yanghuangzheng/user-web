package initialize

import(
	"user-web/global"
	"user-web/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"
    //"github.com/hashicorp/consul/api"
    //_ "github.com/mbobakov/grpc-consul-resolver" // 导入Consul resolver，注册resolver
)
func InitSrvConnng(){
	conn, err := grpc.Dial("127.0.0.1:8088",
        grpc.WithInsecure(), // 使用不安全的连接，即不使用TLS
    )
    if err != nil {
        // 如果连接失败，记录错误并终止程序
        zap.S().Fatal("[srvConn] 连接 [用户服务失败]", zap.Error(err))
    }
    // 创建客户端存根
    userSrvClient := proto.NewUserClient(conn)
    global.UserSrvClient = userSrvClient

    // 在函数结束时关闭连接
    //defer conn.Close()
}
/*func InitSrvConnng(){
	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s",global.ServerConfig.ConsulInfo.Host,global.ServerConfig.ConsulInfo.Port,global.ServerConfig.Name), //tag与名字要对应
		grpc.WithInsecure(), // 使用不安全的连接，即不使用TLS
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), // 设置默认的服务配置，这里使用轮询负载均衡策略
	)
	if err != nil {
		// 如果连接失败，记录错误并终止程序
		zap.S().Fatal("[srvConnng]连接[用户服务失败]")
	}
	
	 //1.后续的用户服务下线2.该端口3改ip
  //已经事先创立好了链接资源后续就不用再次进行tcp的三次握手
  //一个连接多个协程共用 连接池
  userSrvClient:=proto.NewUserClient(userConn)
  global.UserSrvClient=userSrvClient
  // 当main函数结束时，关闭gRPC连接
	defer userConn.Close()
}*/