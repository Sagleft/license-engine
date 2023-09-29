package license

import (
	"bytes"
	"encoding/base32"
	"encoding/gob"

	"github.com/hyperboloide/lk"
)

func createLicenseFromStruct(lc *lk.License) License {
	return &defaultLicense{licenseData: lc}
}

func toB32String(obj interface{}) (string, error) {
	b, err := toBytes(obj)
	if err != nil {
		return "", err
	}

	return base32.StdEncoding.EncodeToString(b), nil
}

func toBytes(obj interface{}) ([]byte, error) {
	var buffBin bytes.Buffer

	encoderBin := gob.NewEncoder(&buffBin)
	if err := encoderBin.Encode(obj); err != nil {
		return nil, err
	}

	return buffBin.Bytes(), nil
}

func EncodeKeyToBase32(key string) (string, error) {
	return toB32String(key)
}
