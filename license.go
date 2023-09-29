package license

import (
	"encoding/json"
	"fmt"
	"time"

	swissknife "github.com/Sagleft/swiss-knife"
	"github.com/hyperboloide/lk"
)

const dateFormat = "2006-01-02"

type License interface {
	Encode() (string, error)
	Save(filepath string) error
	Validate(appPublicKey string) error
}

type defaultLicense struct {
	licenseData *lk.License
}

func (l *defaultLicense) Save(filepath string) error {
	licData, err := l.Encode()
	if err != nil {
		return fmt.Errorf("encode: %w", err)
	}

	if err := swissknife.SaveStringToFile(filepath, licData); err != nil {
		return fmt.Errorf("save: %w", err)
	}
	return nil
}

func (l *defaultLicense) Encode() (string, error) {
	// the b32 representation of our license
	// this is what you give to your customer.
	licenseB32, err := l.licenseData.ToB32String()
	if err != nil {
		return "", fmt.Errorf("encode license file: %w", err)
	}

	return licenseB32, nil
}

func (l *defaultLicense) Validate(appPublicKey string) error {
	publicKey, err := lk.PublicKeyFromB32String(appPublicKey)
	if err != nil {
		return fmt.Errorf("decode public key: %w", err)
	}

	if ok, err := l.licenseData.Verify(publicKey); err != nil {
		return fmt.Errorf("verify license: %w", err)
	} else if !ok {
		return fmt.Errorf("invalid license signature")
	}

	var uData UserData
	if err := json.Unmarshal(l.licenseData.Data, &uData); err != nil {
		return fmt.Errorf("decode license: %w", err)
	}

	if uData.RegisteredUntil.Before(time.Now()) {
		return fmt.Errorf(
			"license expired on: %s",
			uData.RegisteredUntil.Format(dateFormat),
		)
	}
	return nil
}
