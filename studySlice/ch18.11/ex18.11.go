package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}

	slice = append(slice, 0)
	idx := 2 //추가하려는 위치

	//맨 뒤부터 추가하려는 위치까지 값을하나씩 옮겨준다.

	for i := len(slice) - 2; i >= idx; i-- {
		slice[i+1] = slice[i]
	}
	fmt.Println(slice[0], slice[1], slice[2], slice[3], slice[6], "s2")
	//값변경
	slice[idx] = 100
	fmt.Println(slice)
}
