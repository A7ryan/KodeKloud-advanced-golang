package main

import (
	"fmt"
)

func process(i int) error {
	if i%2 == 0 {
		// return errors.New("only odd numbers allowed")
		return fmt.Errorf("only odd numbers allowed, got: %d", i)
	}
	return nil
}

func checkErr(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println("Operation Successfuly..")
}

func main() {
	err := process(3)
	checkErr(err)
	err = process(2)
	checkErr(err)
}
