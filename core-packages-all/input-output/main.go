package main

import (
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReader("Learning is Fun!")
	buff := make([]byte, r.Len())

	for {
		n, err := r.Read(buff)
		fmt.Println(string(buff[:n]), err)
		if err != nil {
			break
		}
	}

	// if _, err := io.Copy(os.Stdout, r); err != nil {
	// 	log.Fatal(err)
	// }
}
