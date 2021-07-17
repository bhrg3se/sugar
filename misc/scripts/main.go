package main

import (
	"encoding/csv"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
	"sugar/server"
	"sugar/store"
)

func main() {
	config := server.ParseConfig("/etc/sugar/config")
	store.State = store.NewRealStore(config)

	f, err := os.Open("doors.csv")
	if err != nil {
		panic(err)
	}
	csvr := csv.NewReader(f)

	//skip first row
	csvr.Read()

	for record, errRead := csvr.Read(); errRead == nil; record, errRead = csvr.Read() {
		doorID := uuid.New().String()
		_, err = store.State.DB.Exec(`INSERT INTO doors (name,id,acme_device_id) VALUES ($1,$2,$3)`, record[0], doorID, record[1])
		if err != nil {
			logrus.Error(err)
		}
	}
	f.Close()

	emptyPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(""), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	f, err = os.Open("residents.csv")
	if err != nil {
		panic(err)
	}
	csvr = csv.NewReader(f)

	//skip first row
	csvr.Read()

	for record, errRead := csvr.Read(); errRead == nil; record, errRead = csvr.Read() {
		userID := uuid.New().String()
		firstName := record[1]
		lastName := record[0]
		email := record[2]
		doors := record[3]

		_, err = store.State.DB.Exec(`INSERT INTO users (id,last_name,first_name,email,is_admin,password) VALUES ($1,$2,$3,$4,'false',$5)`, userID, lastName, firstName, email, string(emptyPasswordBytes))
		if err != nil {
			logrus.Error(err)
		}

		for _, doorName := range strings.Split(doors, ",") {
			permissionID := uuid.New().String()
			_, err = store.State.DB.Exec(`INSERT INTO permissions (id, user_id, door_id)
												SELECT $1,$2, (select id from doors where name=$3);`,
				permissionID, userID, doorName)
			if err != nil {
				logrus.Error(err)
			}
		}

	}
	f.Close()

	f.Seek(0, 0)

	csvr = csv.NewReader(f)
	//
	//for record,errRead:=csvr.Read();errRead!=nil;record,errRead=csvr.Read(){
	//
	//	var doorID string;
	//
	//
	//	id:=uuid.New().String()
	//	_,err=store.State.DB.Exec(`INSERT INTO access (id,user_id,door_id) VALUES ($1,$2)`,id,record[0],record[1],record[2])
	//	if err != nil {
	//		logrus.Error(err)
	//	}
	//}
	//

}
