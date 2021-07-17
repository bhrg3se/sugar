package users

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"sugar/models"
	"sugar/store"
	"sugar/utils"
)

func GetUser(writer http.ResponseWriter, request *http.Request) {
	userID := request.URL.Query().Get("user")

	user, err := FetchUser(store.State.DB, userID)
	if err != nil {
		utils.ErrorResponse(writer, "could not get user", 500)
		return
	}
	utils.SuccessResponse(writer, user, 200)

}

func AddUser(writer http.ResponseWriter, request *http.Request) {
	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		utils.ErrorResponse(writer, "json parse error", 400)
		return
	}

	user.ID = uuid.New().String()
	token, err := utils.SignToken(user.ID)
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponse(writer, "could not generate token", 500)
		return
	}

	err = addUser(store.State.DB, &user)
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponse(writer, "could not save user", 500)
		return
	}

	//TODO send email to user with "change password" URL

	utils.SuccessResponse(writer, token, 200)
}

func RemoveUser(writer http.ResponseWriter, request *http.Request) {
	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		utils.ErrorResponse(writer, "json parse error", 400)
		return
	}

	err = removeUser(store.State.DB, user.ID)
	if err != nil {
		utils.ErrorResponse(writer, "could not remove user", 500)
		return
	}
	utils.SuccessResponse(writer, "user removed", 200)
}
