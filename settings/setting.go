package settings

import "gopkg.in/ini.v1"

var Conf = new(AppConfig)

// AppConfig 应用程序配置
type AppConfig struct {
	Release      bool `json:"release"`
	Port         int  `json:"port"`
	*MysqlConfig `ini:"mysql"`
}
type MysqlConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	DB       string `ini:"db"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
}

func Init(file string) error {
	return ini.MapTo(Conf, file)
}
