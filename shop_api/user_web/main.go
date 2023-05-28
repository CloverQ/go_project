package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"shop_api/user_web/global"
	"shop_api/user_web/initialize"
	myvalidator "shop_api/user_web/validator"
)

func main() {
	// 1. 初始化logger
	initialize.InitLogger()
	// 2. 初始化配置文件
	initialize.InitConfig()
	// 3. 初始化 routers
	Router := initialize.Routers()
	// 4. 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}
	// 5. 初始化srv的连接
	initialize.InitSrvConn()
	//注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}
	zap.S().Debugf("启动服务器，端口: %d", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败: ", err.Error())
	}
}
