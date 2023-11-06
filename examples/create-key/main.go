package main

import (
	"fmt"
	"log"

	"github.com/Sagleft/license-engine"
)

func main() {
	keypair, err := license.CreateNewKeypair()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("private key:", keypair.Private)
	fmt.Println()
	fmt.Println("public key:", keypair.Public)
	fmt.Println()
}
