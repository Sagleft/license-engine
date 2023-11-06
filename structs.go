package license

import "time"

type LicenseData struct {
	Data            any       `json:"data"`
	AppKeyHash      string    `json:"appKeyHash"`
	RegisteredUntil time.Time `json:"until"`
}

type Keypar struct {
	Public  string `json:"public"`
	Private string `json:"secret"`
}
