package migrated

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
)

func CreateMigration(migrationDir string, timestamp int64, name string) (base string, err error) {
	base = filepath.Join(migrationDir, fmt.Sprintf("%v_%v", timestamp, name))
	if err := os.MkdirAll(migrationDir, os.ModePerm); err != nil {
		return base, err
	}
	if _, err := os.Create(base + ".up.sql"); err != nil {
		return base, err
	}
	if _, err := os.Create(base + ".down.sql"); err != nil {
		return base, err
	}

	return base, nil
}

func UpMigration(m *migrate.Migrate, limit int) error {
	if limit > 0 {
		return m.Steps(limit)
	}

	return m.Up()
}

func DownMigration(m *migrate.Migrate) error {
	return m.Steps(-1)
}

func ForceMigration(m *migrate.Migrate, version int) error {
	return m.Force(version)
}