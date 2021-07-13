package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Fail")
		return
	}

	defer fmt.Println("반드시호출")
	defer f.Close() //지연수행
	defer fmt.Println("파일닫음")

	fmt.Println("파일에 Hello작성")
	fmt.Fprintln(f, "hello")
}
