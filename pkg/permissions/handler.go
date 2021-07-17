package permissions

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"sugar/models"
	"sugar/store"
	"sugar/utils"
)

func AddPermission(writer http.ResponseWriter, request *http.Request) {
	var permission models.Permission
	err := json.NewDecoder(request.Body).Decode(&permission)
	if err != nil {
		utils.ErrorResponse(writer, "json parse error", 400)
		return
	}

	permission.ID = uuid.New().String()

	err = addPermission(store.State.DB, &permission)
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponse(writer, "could not save permission", 500)
		return
	}

	utils.SuccessResponse(writer, permission, 200)
}

func RemovePermission(writer http.ResponseWriter, request *http.Request) {
	var permission models.Permission
	err := json.NewDecoder(request.Body).Decode(&permission)
	if err != nil {
		utils.ErrorResponse(writer, "json parse error", 400)
		return
	}

	err = removePermission(store.State.DB, permission.UserID, permission.DoorID)
	if err != nil {
		utils.ErrorResponse(writer, "could not remove permission", 500)
		return
	}
	utils.SuccessResponse(writer, nil, 200)
}
