package cobracli

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"aegis/pkg/migrated"
)

func NewMigrationCommand(
	config *migrated.MigratorConfig,
	c *cobra.Command,
) *cobra.Command {
	logger := config.Logger

	abortWithError := func(err error) {
		logger.Errorf("migrate: %+v", err)
		os.Exit(1)
	}

	c.AddCommand(&cobra.Command{
		Use:   "create NAME",
		Short: "create a migration",
		Args:  cobra.MinimumNArgs(1),
		Run: func(c *cobra.Command, args []string) {
			migrationBaseDir := config.MustGetMigrationBaseDir()
			name := args[0]

			base, err := migrated.CreateMigration(migrationBaseDir, time.Now().Unix(), name)
			if err != nil {
				abortWithError(errors.Wrap(err, "migration.CreateMigration"))
			}
			logger.Infof("created %s\n", base)
		},
	})

	addCommand := &cobra.Command{
		Use:   "up",
		Short: "apply migrations",
		Run: func(c *cobra.Command, args []string) {
			limit, err := c.LocalFlags().GetInt("limit")
			if err != nil {
				abortWithError(errors.WithStack(err))
			}

			err = migrated.UpMigration(config.MustNewMigrator(), limit)
			switch err {
			case nil:
			case migrate.ErrNoChange:
				logger.Warnf("no changes\n")
			case os.ErrNotExist:
				logger.Warnf("no more migrations can be applied\n")
			default:
				abortWithError(errors.Wrap(err, "up"))
			}
			logger.Infof("migration(s) applied\n")
		},
	}
	addCommand.Flags().IntP("limit", "n", -1, "apply limit, defaults to apply all")
	c.AddCommand(addCommand)

	c.AddCommand(&cobra.Command{
		Use:   "down",
		Short: "rollback one migration",
		Run: func(c *cobra.Command, args []string) {
			err := migrated.DownMigration(config.MustNewMigrator())
			switch err {
			case nil:
			case os.ErrNotExist:
				logger.Warnf("no more migrations can be rollbacked\n")
			default:
				abortWithError(errors.Wrap(err, "down"))
			}
			logger.Infof("migration rollbacked\n")
		},
	})

	c.AddCommand(&cobra.Command{
		Use:   "force",
		Short: "force to migration version",
		Args:  cobra.MinimumNArgs(1),
		Run: func(c *cobra.Command, args []string) {
			version, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				abortWithError(errors.WithStack(err))
			}

			err = migrated.ForceMigration(config.MustNewMigrator(), int(version))
			if err != nil {
				abortWithError(errors.Wrap(err, "migration.ForceMigration"))
			}
			logger.Warnf("forced to version %d\n", version)
		},
	})

	return c
}
