/*
	cryptography:
		secure communication in presence of third parties
		use encrypt/decrypt tech so digital data is not exploited to unwanted/harmful entity
*/

package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func encodeWithMD5(str string) string {
	var hashCode = md5.Sum([]byte(str))
	return hex.EncodeToString(hashCode[:])
}

func main() {
	var password string
	fmt.Println("Enter Password: ")
	fmt.Scanln(&password)
	fmt.Println("Password Encrypted: ", encodeWithMD5(password))
}
