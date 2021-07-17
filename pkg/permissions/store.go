package permissions

import (
	"database/sql"
	"sugar/models"
)

func addPermission(db *sql.DB, permission *models.Permission) error {
	_, err := db.Exec(`INSERT INTO permissions (id, user_id,door_id) VALUES ($1,$2,$3)`,
		permission.ID, permission.UserID, permission.DoorID)
	return err
}

func removePermission(db *sql.DB, userID, doorID string) error {
	_, err := db.Exec(`DELETE FROM permissions WHERE user_id=$1 AND door_id=$2`, userID, doorID)
	return err
}
