package license

import (
	"bytes"
	"encoding/base32"
	"encoding/gob"
	"fmt"
	"net"

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

func GetMACAddress() (string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return "", fmt.Errorf("get net interfaces: %w", err)
	}

	for _, iface := range ifas {
		if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 {
			return iface.HardwareAddr.String(), nil
		}
	}
	return "00:00:00:00:00:00", nil
}

// CreateMachinePrivateKey - create a key unique for this machine
func CreateMachinePrivateKey(salt string) (string, error) {
	mac, err := GetMACAddress()
	if err != nil {
		return "", fmt.Errorf("get net address: %w", err)
	}

	privateKey, err := EncodeKeyToBase32(fmt.Sprintf(
		"%s-%s", mac, salt,
	))
	if err != nil {
		return "", fmt.Errorf("encode private key: %w", err)
	}

	return privateKey, nil
}
