package main

import (
	"fmt"
	"log"

	"github.com/Sagleft/license-engine"
)

func main() {
	machinePrivateKey, err := license.CreateNewPrivateKey()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("create private key:", machinePrivateKey)
}
