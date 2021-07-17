package auth

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"sugar/store"
	"sugar/utils"
)

func SetPassword(writer http.ResponseWriter, request *http.Request) {

	token := request.URL.Query().Get("token")
	user, err := utils.VerifyAndParseToken(token)
	if err != nil {
		logrus.Debug(err)
		utils.ErrorResponse(writer, "could not verify token", 401)
		return
	}

	var req struct {
		Password string `json:"password"`
	}
	err = json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponse(writer, "json parse error", 400)
		return
	}
	passBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("could not generate password hash: %v", err)
		utils.ErrorResponse(writer, "could not sign up", 500)
		return
	}
	req.Password = string(passBytes)

	err = setPassword(store.State.DB, user.UserID, req.Password)
	if err != nil {
		utils.ErrorResponse(writer, "could not save data", 500)
		return
	}
	utils.SuccessResponse(writer, nil, 200)
}
