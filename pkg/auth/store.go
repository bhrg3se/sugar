package auth

import (
	"database/sql"
	"sugar/models"
)

func getAuthDetail(db *sql.DB, email string) (models.User, error) {
	var user models.User
	err := db.QueryRow(`SELECT id,email,password FROM users where email=$1`, email).
		Scan(&user.ID, &user.Email, &user.Password)
	return user, err
}

func setPassword(db *sql.DB, userID, password string) error {
	_, err := db.Exec(`UPDATE users SET password=$1 WHERE id=$2`, userID, password)
	return err
}
