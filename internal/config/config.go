package config

import (
	"github.com/spf13/viper"
)

var (
	Api        *viper.Viper
	MySQL      *viper.Viper
	ClickHouse *viper.Viper
	Engine     *viper.Viper
	Admin      *viper.Viper
)

func Init(configName string) {
	if configName == "" {
		configName = "aegis"
	}

	viper.SetConfigName(configName)
	viper.AddConfigPath("/etc/aegis")
	viper.AddConfigPath("$HOME/aegis")
	viper.AddConfigPath("../config/")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	Api = viper.Sub("Api")
	if Api == nil {
		Api = viper.New()
	}

	Api.SetDefault("http", map[string]interface{}{
		"port": 3210,
		"host": "127.0.0.1",
	})

	MySQL = viper.Sub("mysql")
	if MySQL == nil {
		MySQL = viper.New()
	}
	MySQL.SetDefault("conn", map[string]interface{}{
		"host":     "127.0.0.1",
		"port":     3306,
		"username": "root",
		"password": "",
		"dbname":   "aegis",
		"charset":  "utf8mb4",
		"driver":   "mysql",
	})

	ClickHouse = viper.Sub("clickhouse")
	if ClickHouse == nil {
		ClickHouse = viper.New()
	}
	ClickHouse.SetDefault("conn", map[string]interface{}{
		"host":   "127.0.0.1",
		"port":   9000,
		"dbname": "aegis",
		"debug":  false,
	})

	Engine = viper.Sub("engine")
	if Engine == nil {
		Engine = viper.New()
	}
	Engine.SetDefault("grpc", map[string]interface{}{
		"port": 3210,
		"host": "127.0.0.1",
	})

	Admin = viper.Sub("admin")
	if Admin == nil {
		Admin = viper.New()
	}
	Admin.SetDefault("http", map[string]interface{}{
		"port": 3220,
		"host": "127.0.0.1",
	})
	Admin.SetDefault("engine", map[string]interface{}{
		"grpc_url": "127.0.0.1:3210",
	})
}

func IsProduction() bool {
	return false
}
