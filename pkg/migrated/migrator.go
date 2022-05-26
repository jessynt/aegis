package migrated

import (
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/source"
)

type LoggerT interface {
	Infof(format string, a ...interface{})
	Warnf(format string, a ...interface{})
	Errorf(format string, a ...interface{})
}

type MigratorConfig struct {
	Logger           LoggerT
	MigrationBaseDir string
	CreateDatabase   func() (string, database.Driver, error)
	CreateSource     func() (string, source.Driver, error)
}

func (c *MigratorConfig) MustNewMigrator() *migrate.Migrate {
	sourceName, source, err := c.CreateSource()
	c.abortIfError(err)
	driverName, driver, err := c.CreateDatabase()
	c.abortIfError(err)
	migrator, err := migrate.NewWithInstance(
		sourceName, source,
		driverName, driver,
	)
	c.abortIfError(err)

	return migrator
}

func (c *MigratorConfig) MustGetMigrationBaseDir() string {
	migrationBaseDir, err := filepath.Abs(c.MigrationBaseDir)
	c.abortIfError(err)
	return migrationBaseDir
}

func (c *MigratorConfig) abortIfError(err error) {
	if err == nil {
		return
	}
	panic(err)
}
