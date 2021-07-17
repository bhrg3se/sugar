package doors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sugar/store"
)

func callAcmeAPI(deviceID, action string) (*acmeApiResponse, error) {

	url := fmt.Sprintf("%s/devices/%s/%s",
		store.State.Config.Acme.ApiURL,
		deviceID,
		action,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create http request: %v", err)
	}
	req.Header.Add("Authorization", "Bearer "+store.State.Config.Acme.ApiKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not complete http request: %v", err)
	}
	defer resp.Body.Close()

	var response acmeApiResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not parse http response data: %v", err)
	}

	return &response, nil
}
