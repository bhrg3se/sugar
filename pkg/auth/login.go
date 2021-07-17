package auth

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"sugar/store"
	"sugar/utils"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		utils.ErrorResponse(writer, "invalid json data", 400)
		return
	}

	user, err := getAuthDetail(store.State.DB, req.Email)
	if err != nil {
		logrus.Debug(err)
		utils.ErrorResponse(writer, "user not found", 404)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		utils.ErrorResponse(writer, "incorrect password", 401)
		return
	}

	token, err := utils.SignToken(user.ID)
	if err != nil {
		logrus.Debug(err)
		utils.ErrorResponse(writer, "could not generate token", 500)
		return
	}

	utils.SuccessResponse(writer, token, 200)
}
