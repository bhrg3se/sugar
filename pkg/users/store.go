package users

import (
	"database/sql"
	"sugar/models"
)

func FetchUser(db *sql.DB, userID string) (models.User, error) {
	var user models.User
	err := db.QueryRow(`SELECT id,email,is_admin,first_name,last_name FROM users where id=$1`, userID).
		Scan(&user.ID, &user.Email, &user.IsAdmin, &user.FirstName, &user.LastName)
	return user, err
}

func addUser(db *sql.DB, user *models.User) error {
	_, err := db.Exec(`INSERT INTO users (id, email,password,is_admin,first_name,last_name) VALUES ($1,$2,$3,$4,$5,$6)`,
		user.ID, user.Email, user.Password, user.IsAdmin, user.FirstName, user.LastName)
	return err
}

func removeUser(db *sql.DB, userID string) error {
	_, err := db.Exec(`DELETE FROM users WHERE id=$1`, userID)
	return err
}
