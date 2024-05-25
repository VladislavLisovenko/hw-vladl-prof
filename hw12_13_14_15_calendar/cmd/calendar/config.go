package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.
type Config struct {
	Logger     LoggerConf
	Storage    StorageConf
	HTTPServer HTTPServerConf
	GRPCServer GRPCServerConf
}

type LoggerConf struct {
	Level string
}

type StorageConf struct {
	Type     string
	Postgres PostgresConf
}

type PostgresConf struct {
	DataSourceName string
}

type HTTPServerConf struct {
	Host string
	Port string
}

type GRPCServerConf struct {
	Host string
	Port string
}

func NewConfig(path string) Config {
	viper.SetConfigFile(path)
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	config := Config{}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return config
}
