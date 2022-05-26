package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"

	"aegis/internal/config"
	"aegis/internal/model"
	"aegis/internal/module/collection_dao"
	propertyDao "aegis/internal/module/property_dao"
	"aegis/internal/mysql"
	"aegis/internal/proto"
)

func MustMakeSeedMetadataCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "seed:metadata",
		Short: "run metadata seeder",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			var mysqlConn *sqlx.DB
			{
				var err error
				mysqlConn, err = mysql.CreateConnx(config.MySQL)
				if err != nil {
					panic(err)
				}
			}
			var clickHouseConn *sqlx.DB
			{
				var err error
				clickHouseConn, err = sqlx.Open("clickhouse", fmt.Sprintf(
					"tcp://%s:%s?database=%s&debug=%s",
					config.ClickHouse.GetString("conn.host"),
					config.ClickHouse.GetString("conn.port"),
					config.ClickHouse.GetString("conn.dbname"),
					config.ClickHouse.GetString("conn.debug"),
				))
				if err != nil {
					panic(err)
				}
			}

			seedProperties(ctx, mysqlConn, clickHouseConn)
			seedCollections(ctx, mysqlConn)
		},
	}
}

func seedProperties(ctx context.Context, mysqlConn *sqlx.DB, clickHouseConn *sqlx.DB) {
	var err error
	_, err = mysqlConn.ExecContext(ctx, `truncate table properties`)
	if err != nil {
		abort("seed:properties: %+v", err)
	}

	seeder := func(name, label string, t proto.PropertyType, validateType proto.ValidateType, validateArgs string) {
		err = propertyDao.CreatePropertyMetadata(ctx, mysqlConn, &model.Property{
			Name:         name,
			Label:        label,
			Type:         int32(t),
			ValidateType: int32(validateType),
			ValidateArgs: validateArgs,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		})
		if err != nil {
			panic(err)
		}
	}

	seeder("ModelId", "模型ID", proto.PropertyTypeInteger, proto.ValidateTypeNumber, "")
	seeder("Date", "日期", proto.PropertyTypeDate, proto.ValidateTypeString, "")
	seeder("Time", "时间", proto.PropertyTypeInteger, proto.ValidateTypeNumber, "")
	seeder("Year", "年份", proto.PropertyTypeInteger, proto.ValidateTypeNumber, "")
	seeder("Quarter", "季度", proto.PropertyTypeInteger, proto.ValidateTypeNumber, "")
	seeder("Month", "月份", proto.PropertyTypeInteger, proto.ValidateTypeNumber, "")
	seeder("DayOfMonth", "DayOfMonth", proto.PropertyTypeInteger, proto.ValidateTypeNumber, "")
	seeder("DayOfWeek", "DayOfWeek", proto.PropertyTypeInteger, proto.ValidateTypeNumber, "")
	seeder("UserId", "用户ID", proto.PropertyTypeInteger, proto.ValidateTypeString, "")
	seeder("DeviceId", "模型ID", proto.PropertyTypeString, proto.ValidateTypeString, "")
	seeder("IP", "IP", proto.PropertyTypeString, proto.ValidateTypeString, "")
	seeder("Country", "国家", proto.PropertyTypeString, proto.ValidateTypeString, "")
	seeder("Province", "省份", proto.PropertyTypeString, proto.ValidateTypeString, "")
	seeder("City", "城市", proto.PropertyTypeString, proto.ValidateTypeString, "")
}

func seedCollections(ctx context.Context, mysqlConn *sqlx.DB) {
	var err error
	_, err = mysqlConn.ExecContext(ctx, `truncate table collections`)
	if err != nil {
		abort("seed:collections: %+v", err)
	}

	seeder := func(name, label string, propertyId int64, t proto.CollectionType) {
		err = collection_dao.CreateCollection(ctx, mysqlConn, &model.Collection{
			Name:       name,
			Label:      label,
			PropertyId: propertyId,
			Type:       int32(t),
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		})
		if err != nil {
			panic(err)
		}
	}

	seeder("system_user_id_black_list", "[系统]用户黑名单", 9, proto.CollectionTypeBlackList)
	seeder("system_device_id_black_list", "[系统]设备黑名单", 10, proto.CollectionTypeBlackList)
	seeder("system_ip_black_list", "[系统]IP黑名单", 11, proto.CollectionTypeBlackList)
	seeder("system_user_white_list", "[系统]用户百名单", 9, proto.CollectionTypeWhiteList)
}
