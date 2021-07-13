package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 1, 3, 6, 2, 4}
	sort.Ints(s)
	fmt.Println(s)
}
