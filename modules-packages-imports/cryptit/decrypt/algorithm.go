package decrypt

func NimbusReverse(str string) string {
	decryptedStr := ""
	for _, c := range str {
		asciiCode := int(c)
		character := string(asciiCode - 3)
		decryptedStr += character
	}
	return decryptedStr
}
