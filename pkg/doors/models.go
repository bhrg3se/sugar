package doors

type acmeApiResponse struct {
	DeviceID string `json:"device_id"`
	Locked   bool   `json:"locked"`
	Status   string `json:"status"`
}
