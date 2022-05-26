package cmd

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4/database"
	migrateMysql "github.com/golang-migrate/migrate/v4/database/mysql"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"aegis/internal/mysql"
	"aegis/migrations/metadata"
	"aegis/pkg/migrated"
	migratedcli "aegis/pkg/migrated/cobracli"
	"aegis/pkg/migrated/gobindata"
)

func MustMakeMigrationMetadataCommand(config *viper.Viper) *cobra.Command {
	migrator := &migrated.MigratorConfig{
		Logger:           &logger{},
		MigrationBaseDir: "./migrations/metadata",
		CreateDatabase: func() (string, database.Driver, error) {
			db, err := sql.Open("mysql", mysql.GetDSN(config))
			if err != nil {
				panic(err)
			}
			driver, err := migrateMysql.WithInstance(db, &migrateMysql.Config{
				MigrationsTable: "metadata_migrations",
			})
			if err != nil {
				panic(err)
			}
			return "mysql", driver, nil
		},
		CreateSource: gobindata.NewSourceFrom(
			metadata.AssetNames,
			metadata.Asset,
		),
	}
	return migratedcli.NewMigrationCommand(migrator, &cobra.Command{
		Use:   "migration:metadata (create|up|down)",
		Short: "metadata database migration",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	})
}
