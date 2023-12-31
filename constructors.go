package license

import (
	"encoding/json"
	"fmt"

	swissknife "github.com/Sagleft/swiss-knife"
	"github.com/hyperboloide/lk"
)

// GrantUserAccess - register user access and create license.
//
// secretKey - base32 encoded private key generated by `lkgen gen` note that you might
// prefer reading it from a file, and that it should stay secret
// (ie: dont distribute it with your app)!
func GrantUserAccess(data LicenseData, licensePrivateKey string) (License, error) {
	pk, err := lk.PrivateKeyFromB32String(licensePrivateKey)
	if err != nil {
		return nil, fmt.Errorf("decode app key: %w", err)
	}

	licBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("encode license to json: %w", err)
	}

	// generate license with the private key and the document
	lc, err := lk.NewLicense(pk, licBytes)
	if err != nil {
		return nil, fmt.Errorf("create license: %w", err)
	}

	return createLicenseFromStruct(lc), nil
}

// GrantUserAccessAndSave - register user access and create license file
func GrantUserAccessAndSave(
	data LicenseData,
	licensePrivateKey string,
	licenseFilePath string,
) error {
	l, err := GrantUserAccess(data, licensePrivateKey)
	if err != nil {
		return fmt.Errorf("grant access: %w", err)
	}

	if err := l.Save(licenseFilePath); err != nil {
		return fmt.Errorf("save license: %w", err)
	}
	return nil
}

func LoadFromFile(filepath string) (License, error) {
	licData, err := swissknife.ReadFileToString(filepath)
	if err != nil {
		return nil, fmt.Errorf("read license file: %w", err)
	}

	lc, err := lk.LicenseFromB32String(licData)
	if err != nil {
		return nil, fmt.Errorf("parse license: %w", err)
	}

	return createLicenseFromStruct(lc), nil
}

// ValidateUserLicenseFile - load & validate user license file
func ValidateUserLicenseFile(
	licenseFilePath string,
	licensePublicKey string,
	appKeyHash string,
) error {
	l, err := LoadFromFile(licenseFilePath)
	if err != nil {
		return fmt.Errorf("load: %w", err)
	}

	if err := l.Validate(appKeyHash, licensePublicKey); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}
