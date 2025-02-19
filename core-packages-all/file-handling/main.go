// Getting Metadata, create new file, read/write file

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path := filepath.Join("dir1", "dir2", "test.txt")
	fmt.Println(path)
	fmt.Println(filepath.Dir(path))
	fmt.Println(filepath.Base(path))
	fmt.Println(filepath.IsAbs(path))
	fmt.Println(filepath.Ext(path))

	fileInfo, err := os.Stat("/<os_path>/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.IsDir())

	data, err := os.ReadFile("/<os_path>/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	file, err := os.Open("/<os_path>/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	b := make([]byte, 4)
	for {
		n, err := file.Read(b)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(b[:n]))

	}

	newFile, err := os.OpenFile("<os_path>/test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer newFile.Close()

	_, err = newFile.WriteString("Hope you had a good day!")
	if err != nil {
		fmt.Println(err)
	}

}
