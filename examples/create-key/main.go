package main

import (
	"fmt"
	"log"

	"github.com/Sagleft/license-engine"
)

func main() {
	salt := "test-123456789"
	machinePrivateKey, err := license.CreateMachinePrivateKey(salt)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("use machine private key:", machinePrivateKey)
}
