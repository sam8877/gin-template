package conf

import (
	"github.com/spf13/viper"
)

type ConfStruct struct {
	Auth *AuthStruct `mapstructure:"auth"`
}

type AuthStruct struct {
	TokenSecretKey      string `mapstructure:"secret_key"`
	TokenExpireSeconds  uint   `mapstructure:"expire_seconds"`
	TokenRefreshSeconds uint   `mapstructure:"refresh_seconds"`
}

var Instance = &ConfStruct{}

func GetAuthConf() *AuthStruct {
	return Instance.Auth
}

func init() {
	v := viper.New()
	v.SetConfigName("conf")
	v.SetConfigType("yaml")
	v.AddConfigPath("./")
	v.ReadInConfig()
	v.Unmarshal(Instance)
}
