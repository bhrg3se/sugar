package doors

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"sugar/models"
	"sugar/store"
	"sugar/utils"
)

func AddDoor(writer http.ResponseWriter, request *http.Request) {
	var door models.Door
	err := json.NewDecoder(request.Body).Decode(&door)
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponse(writer, "invalid json data", 400)
		return
	}
	door.ID = uuid.New().String()
	err = addDoor(store.State.DB, door)
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponse(writer, "could not add door", 500)
		return
	}

	utils.SuccessResponse(writer, door, 200)
}

func RemoveDoor(writer http.ResponseWriter, request *http.Request) {
	var door models.Door
	err := json.NewDecoder(request.Body).Decode(&door)
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponse(writer, "invalid json data", 400)
		return
	}

	err = removeDoor(store.State.DB, door.ID)
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponse(writer, "could not remove door", 500)
		return
	}

	utils.SuccessResponse(writer, nil, 200)
}

func MyDoorsList(writer http.ResponseWriter, request *http.Request) {
	user := request.Context().Value("user").(models.User)

	doors, err := getAccessibleDoors(store.State.DB, user.ID)
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponse(writer, "could not remove door", 500)
		return
	}

	utils.SuccessResponse(writer, doors, 200)
}
