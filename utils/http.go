package utils

import (
	"encoding/json"
	"net/http"
)

func SuccessResponse(writer http.ResponseWriter, data interface{}, status int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	resp := apiResponse{
		Status:  true,
		Message: "",
		Body:    data,
	}

	marshalledResp, _ := json.Marshal(resp)
	writer.Write(marshalledResp)
}

func ErrorResponse(writer http.ResponseWriter, message string, status int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	data := apiResponse{Status: false, Message: message}
	marshalledData, _ := json.Marshal(data)
	writer.Write(marshalledData)
}

type apiResponse struct {
	Status  bool        `json:"status" omitempty`
	Message string      `json:"message" omitempty`
	Body    interface{} `json:"body" omitempty`
}
