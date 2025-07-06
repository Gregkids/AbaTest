package models

type DeviceReq struct {
	DeviceId       string `json:"device_id"`
	DeviceName     string `json:"device_name"`
	DeviceLocation string `json:"location"`
	DeviceStatus   bool   `json:"status"`
}

type DeviceData struct {
	DeviceId       string `json:"device_id"`
	DeviceName     string `json:"device_name"`
	DeviceLocation string `json:"location"`
	DeviceStatus   bool   `json:"status"`
	UpdatedAt      string `json:"updated_at"`
}
