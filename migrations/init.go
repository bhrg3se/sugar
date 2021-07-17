package migrations

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func InitRootUser(db *sql.DB) {
	var rootExists bool
	err := db.QueryRow(`SELECT (EXISTS ( SELECT id FROM users WHERE id='1'))`).Scan(&rootExists)
	if err != nil {
		logrus.Fatal(err)
	}
	pass, err := bcrypt.GenerateFromPassword([]byte("root"), bcrypt.DefaultCost)
	if err != nil {
		logrus.Fatal(err)
	}
	if !rootExists {
		_, err = db.Exec(`INSERT INTO users (id,first_name,last_name,email,password,is_admin) VALUES ('1','root','root','root',$1,true)`, string(pass))
		if err != nil {
			logrus.Fatal(err)
		}
	}
}
