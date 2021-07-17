package migrations

import (
	"database/sql"
)

func Migrate(db *sql.DB) error {
	for _, q := range migartion_v1 {
		_, err := db.Exec(q)
		if err != nil {
			return err
		}
	}
	InitRootUser(db)
	return nil
}
