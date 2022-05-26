package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func GetDSN(config *viper.Viper) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=UTC",
		config.GetString("conn.username"),
		config.GetString("conn.password"),
		config.GetString("conn.host"),
		config.GetString("conn.port"),
		config.GetString("conn.dbname"),
	) // "user:password@tcp(127.0.0.1:3306)/hello?charset=utf8mb4&parseTime=true&loc=UTC"
}

func CreateConnx(config *viper.Viper) (*sqlx.DB, error) {
	return sqlx.Open("mysql", GetDSN(config))
}
