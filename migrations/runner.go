package migrations

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/mizanalyst/mizanalyst/mizanlyst_logger"

	"gorm.io/gorm"
)

const migrationsTable = "schema_migrations"

// RunMigrations reads all .sql files from the migrations directory, skips those
// already recorded in schema_migrations, and executes the rest in sorted order.
// Each migration runs inside its own transaction.
func RunMigrations(db *gorm.DB) error {
	// Ensure tracking table exists
	if err := ensureMigrationsTable(db); err != nil {
		return fmt.Errorf("failed to create migrations tracking table: %w", err)
	}

	// Discover SQL files
	files, err := discoverSQLFiles("migrations")
	if err != nil {
		return fmt.Errorf("failed to discover migration files: %w", err)
	}

	if len(files) == 0 {
		mizanlyst_logger.Log("No migration files found")
		return nil
	}

	// Fetch already-applied migrations
	applied, err := getApplied(db, migrationsTable)
	if err != nil {
		return fmt.Errorf("failed to fetch applied migrations: %w", err)
	}

	for _, file := range files {
		name := filepath.Base(file)

		if applied[name] {
			mizanlyst_logger.Log("Migration already applied, skipping: %s", name)
			continue
		}

		content, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", name, err)
		}

		if err := db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Exec(string(content)).Error; err != nil {
				return fmt.Errorf("migration %s failed: %w", name, err)
			}

			if err := tx.Exec(
				fmt.Sprintf("INSERT INTO %s (name) VALUES (?)", migrationsTable),
				name,
			).Error; err != nil {
				return fmt.Errorf("failed to record migration %s: %w", name, err)
			}

			return nil
		}); err != nil {
			return err
		}

		mizanlyst_logger.Log("Migration applied successfully: %s", name)
	}

	return nil
}

// ensureMigrationsTable creates the tracking table if it doesn't exist.
func ensureMigrationsTable(db *gorm.DB) error {
	sql := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id         BIGSERIAL PRIMARY KEY,
			name       VARCHAR(255) NOT NULL UNIQUE,
			applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);
	`, migrationsTable)

	return db.Exec(sql).Error
}

// discoverSQLFiles returns sorted .sql file paths from the given directory.
func discoverSQLFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if strings.HasSuffix(entry.Name(), ".sql") {
			files = append(files, filepath.Join(dir, entry.Name()))
		}
	}

	sort.Strings(files)
	return files, nil
}

// getApplied returns a set of already-applied file names from the tracking table.
func getApplied(db *gorm.DB, table string) (map[string]bool, error) {
	var names []string
	if err := db.Raw(fmt.Sprintf("SELECT name FROM %s", table)).Scan(&names).Error; err != nil {
		return nil, err
	}

	applied := make(map[string]bool, len(names))
	for _, n := range names {
		applied[n] = true
	}

	return applied, nil
}
