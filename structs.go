package license

import "time"

type LicenseData struct {
	Data            any       `json:"data"`
	RegisteredUntil time.Time `json:"until"`
}
