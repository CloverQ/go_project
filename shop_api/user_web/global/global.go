package global

import (
	ut "github.com/go-playground/universal-translator"
	"shop_api/user_web/config"
)

var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	Trans        ut.Translator
)
