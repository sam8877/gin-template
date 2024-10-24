package conf

import (
	"encoding/json"
	"fmt"
	"gin-template/model"
	"github.com/spf13/viper"
)

var Instance = &model.ConfStruct{}

func GetAuthConf() *model.AuthStruct {
	return Instance.Auth
}

func init() {
	v := viper.New()
	v.SetConfigName("conf")
	v.SetConfigType("yaml")
	v.AddConfigPath("./")
	v.ReadInConfig()
	v.Unmarshal(Instance)

	bts, err := json.Marshal(Instance)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bts))
}
