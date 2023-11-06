package main

import (
	"log"
	"time"

	"github.com/Sagleft/license-engine"
)

const licFilePath = "license.dat"

func main() {
	privateKeyEncoded := "FD7YCAYBAEFXA22DN5XHIYLJNZSXEAP7QIAACAQBANIHKYQBBIAACAKEAH7YIAAAAAFP7AYFAEBP7BQAAAAP7GP7QIAWCBBZQ3NVCEF2WG7UNFPXIFRWGDKTDBABQNPCZCHCIRQ4RHLT5UU4DWB5RJHW2YLOCEG75FPKOGBBLLSESOT6IAO73VY4WJYH4NMB5KH7FQKLAUN7PKSCBTI5YTKUGVSJ4SD5YOYM3GLKPJ2DWQA3AFADK424AEYQEFDDRBDIGSMMJSJFS57CQMGNGXLYO2XVOHSHTCCZCAWEZJ5OSEXBK2EQ7ETRDJPC7OAQVL43IYML4MAA===="

	lic := license.LicenseData{
		Data:            "some user data",
		RegisteredUntil: time.Now().Add(time.Minute * 2),
	}

	if err := license.GrantUserAccessAndSave(
		lic,
		privateKeyEncoded,
		licFilePath,
	); err != nil {
		log.Fatalln(err)
	}
}
