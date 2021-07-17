package doors

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"sugar/models"
	"sugar/store"
	"sugar/utils"
)

func Action(writer http.ResponseWriter, request *http.Request) {
	user := request.Context().Value("user").(models.User)

	var req lockRequest
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponse(writer, "invalid json data", 400)
		return
	}

	// action must be either "lock","unlock" or "check"
	switch req.Action {
	case "lock":
		break
	case "unlock":
		break
	case "check":
		req.Action = ""
		break
	default:
		logrus.Debug(err)
		utils.ErrorResponse(writer, "invalid action", 400)
		return
	}

	//check if user has permission to access the door
	acmeDeviceID, err := getAcmeIDIfAuthorised(store.State.DB, user.ID, req.DoorID)
	if err != nil {
		logrus.Errorf("could not check permission: %v", err)
		utils.ErrorResponse(writer, "could not check permission", 500)
		return
	}

	if acmeDeviceID == "" {
		utils.ErrorResponse(writer, "not authorised to access this door", 403)
		return
	}

	//call acme api to access the door
	statusResp, err := callAcmeAPI(acmeDeviceID, req.Action)
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponse(writer, "could not access the door", 500)
		return
	}

	//let's not expose acme device ID just in case
	statusResp.DeviceID = req.DoorID
	utils.SuccessResponse(writer, statusResp, 200)

}

type lockRequest struct {
	DoorID string `json:"door_id"`
	Action string `json:"action"`
}
