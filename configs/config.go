package configs

import (
	_ "embed"
	"fmt"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

//go:embed openapi.yaml
var OpenAPI []byte

var conf *ApplicationConfig

type ApplicationConfig struct {
	Server Server   `mapstructure:"server"`
	DB     DBConfig `mapstructure:"db"`
}

type Server struct {
	HTTP    ServerConfig `mapstructure:"http"`
	GRPC    ServerConfig `mapstructure:"grpc"`
	OpenAPI ServerConfig `mapstructure:"open_api"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
}

type ServerConfig struct {
	Addr    string `mapstructure:"addr"`
	Timeout int    `mapstructure:"timeout"`
}

func NewConfig() *ApplicationConfig {
	once := new(sync.Once)
	once.Do(func() {
		data := ApplicationConfig{}
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("../../configs")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("failed to read config file: %s", err))
		}

		if err := viper.Unmarshal(&data); err != nil {
			panic(fmt.Errorf("vailed to load config %s", err))
		}

		conf = &data
	})
	return conf
}
