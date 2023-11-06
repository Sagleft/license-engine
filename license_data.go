package license

import (
	"errors"
	"fmt"
	"time"
)

type LicenseData struct {
	Data            any       `json:"data"`
	AppKeyHash      string    `json:"appKeyHash"`
	RegisteredUntil time.Time `json:"until"`
}

func (data LicenseData) CheckTime() error {
	h := newTimeHandler()
	currentTime, err := h.getCurrentTime()
	if err != nil {
		return fmt.Errorf("get current time: %w", err)
	}

	if currentTime == nil {
		return errors.New("failed to get current time: result is nil")
	}

	if data.RegisteredUntil.Before(*currentTime) {
		return fmt.Errorf(
			"license expired on: %s",
			data.RegisteredUntil.Format(dateFormat),
		)
	}
	return nil
}
