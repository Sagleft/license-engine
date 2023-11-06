package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Sagleft/license-engine"
)

const licFilePath = "license.dat"
const appKey = "123-456-789"

func main() {
	privateKeyEncoded := "FD7YCAYBAEFXA22DN5XHIYLJNZSXEAP7QIAACAQBANIHKYQBBIAACAKEAH7YIAAAAAFP7AYFAEBP7BQAAAAP7GP7QIAWCBEMB36G5U7F3XBG2QGTGPDFJQM4XSD7PIUUF4V6YFFATDRATHN4ENNLO4RG2I5XXDV52BKV62N52II7IRNRNLCNUBR23U44XNP4OPSF2B7IOATBCIORSH543CIQLQGKDNKBWEHRJX2DKOJF4IEDF4JK6V6TAEYQEASJRURI4BWECN7QYMYAMBFKWW36YLQ5KS47NRWXPDL67HGFIFPPS7TG7M4B4GILZSRBSUS7N4FRIEAA===="

	publicKeyEncoded := "ASGA57DO2PS53QTNIDJTHRSUYGOLZB7XUKKC6K7MCSQJRYQJTW6CGWVXOITNEO33R265AVK7NG65EEPUIWYWVRG2AY5N2OOLWX6HHZC5A7UHAJQREHIZD66NREIFYDFBWVA3CDYU35BVHES6ECBS6EVPK7JQ===="

	appKeyHash := license.GetAppKeyHash(appKey)
	lic := license.LicenseData{
		Data:            "some user data",
		AppKeyHash:      appKeyHash,
		RegisteredUntil: time.Now().Add(time.Second * 3),
	}

	if err := license.GrantUserAccessAndSave(
		lic,
		privateKeyEncoded,
		licFilePath,
	); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("license file created. wait for 4 seconds..")
	time.Sleep(time.Second * 4)
	fmt.Println("expect error..")

	if err := license.ValidateUserLicenseFile(
		licFilePath,
		publicKeyEncoded,
		appKeyHash,
	); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("OK!")
}
