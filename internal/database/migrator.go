package database

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
)

// RunMigrations ejecuta todos los archivos .sql en orden
func RunMigrations(db *DB, migrationsPath string) error {
	var files []string

	err := filepath.WalkDir(migrationsPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(path) == ".sql" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	sort.Strings(files)

	for _, file := range files {
		if err := runMigrationFile(db, file); err != nil {
			return fmt.Errorf("error en migración %s: %w", file, err)
		}
	}

	return nil
}

func runMigrationFile(db *DB, path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	_, err = db.Conn.Exec(string(content))
	return err
}