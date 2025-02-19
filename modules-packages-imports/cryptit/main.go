package main

import (
	"fmt"

	"github.com/Priyanka-yadavv/cryptit/decrypt"
	"github.com/Priyanka-yadavv/cryptit/encrypt"
)

func main() {
	encryptData := encrypt.Nimbus("KodeKloud")
	fmt.Println("Encrypted Data: ", encryptData)
	fmt.Println("Decrypted Data: ", decrypt.NimbusReverse(encryptData))

}
