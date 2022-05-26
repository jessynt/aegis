package cmd

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4/database"
	migrateClickhouse "github.com/golang-migrate/migrate/v4/database/clickhouse"
	_ "github.com/kshvakov/clickhouse"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"aegis/migrations/warehouse"
	"aegis/pkg/migrated"
	migratedcli "aegis/pkg/migrated/cobracli"
	"aegis/pkg/migrated/gobindata"
)

func MustMakeMigrationWarehouseCommand(config *viper.Viper) *cobra.Command {
	migrator := &migrated.MigratorConfig{
		Logger:           &logger{},
		MigrationBaseDir: "./migrations/warehouse",
		CreateDatabase: func() (string, database.Driver, error) {
			db, err := sql.Open("clickhouse", fmt.Sprintf(
				"tcp://%s:%s?database=%s&debug=%s",
				config.GetString("conn.host"),
				config.GetString("conn.port"),
				config.GetString("conn.dbname"),
				config.GetString("conn.debug"),
			))
			if err != nil {
				panic(err)
			}
			driver, err := migrateClickhouse.WithInstance(db, &migrateClickhouse.Config{
				MigrationsTable: "warehouse_migrations",
			})
			if err != nil {
				panic(err)
			}
			return "mysql", driver, nil
		},
		CreateSource: gobindata.NewSourceFrom(
			warehouse.AssetNames,
			warehouse.Asset,
		),
	}
	return migratedcli.NewMigrationCommand(migrator, &cobra.Command{
		Use:   "migration:warehouse (create|up|down)",
		Short: "warehouse database migration",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	})
}
