package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"shop_api/user_web/global"
	"shop_api/user_web/proto"
)

func InitSrvConn() {
	// 从注册中心获取到用户服务的信息（ip port）
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)

	userSrvHost := ""
	userSrvPost := 0
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	//data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \"%s\"", global.ServerConfig.UserSrvInfo.Name))
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSrvInfo.Name))
	if err != nil {
		panic(err)
	}
	for _, val := range data {
		userSrvHost = val.Address
		userSrvPost = val.Port
		break
	}
	if userSrvHost == "" {
		zap.S().Fatal("[InitSrvConn]连接【用户服务】失败")
		return
	}
	// 拨号连接用户grpc服务
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPost), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList]连接【用户服务】失败", "msg", err.Error())
	}
	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}
