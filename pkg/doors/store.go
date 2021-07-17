package doors

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"sugar/models"
)

func addDoor(db *sql.DB, door models.Door) error {
	_, err := db.Exec(`INSERT INTO doors (id,acme_device_id,name) VALUES ($1,$2,$3)`, door.ID, door.AcmeDeviceID, door.Name)
	return err
}
func removeDoor(db *sql.DB, doorID string) error {
	_, err := db.Exec(`DELETE FROM doors WHERE id=$1`, doorID)
	return err
}

func checkPermission(db *sql.DB, userID, doorID string) (bool, error) {
	var ok bool
	err := db.QueryRow(`SELECT EXISTS (SELECT id FROM permissions WHERE user_id=$1 AND door_id=$2)`, userID, doorID).
		Scan(&ok)
	return ok, err
}

func getAcmeIDIfAuthorised(db *sql.DB, userID, doorID string) (string, error) {
	var acmeDeviceID string
	err := db.QueryRow(`
SELECT (CASE WHEN  (EXISTS (SELECT id FROM permissions WHERE user_id=$1 AND door_id=$2))
THEN
(SELECT acme_device_id FROM doors WHERE id=$2)
ELSE ''
END)
`, userID, doorID).
		Scan(&acmeDeviceID)
	return acmeDeviceID, err
}

func getAccessibleDoors(db *sql.DB, userID string) ([]models.Door, error) {
	var doors = make([]models.Door, 0)

	rows, err := db.Query(`SELECT d.id,d.name FROM doors d JOIN permissions p ON d.id=p.door_id WHERE p.user_id=$1`, userID)
	if err != nil {
		return doors, err
	}
	defer rows.Close()

	for rows.Next() {
		var door models.Door
		err = rows.Scan(&door.ID, &door.Name)
		if err != nil {
			logrus.Error(err)
			continue
		}
		doors = append(doors, door)
	}

	return doors, err
}
