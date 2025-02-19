package main

import (
	"fmt"
	"sort"
)

func main() {
	vars := []int{5, 3, 2, 4, 1}
	vars_str := []string{"learning", "golang", "on", "kodekloud"}

	sort.Ints(vars)
	sort.Strings(vars_str)

	fmt.Println(vars)
	fmt.Println(vars_str)

}
