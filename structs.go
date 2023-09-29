package license

import "time"

type UserData struct {
	Email           string    `json:"email"`
	RegisteredUntil time.Time `json:"until"`
}
