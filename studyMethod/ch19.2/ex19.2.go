package main

import "fmt"

//사용자 정의 별칭타입
type myInt int

//myint 별칭을 리시버로 갖는 메소드
func (a myInt) add(b int) int {
	return int(a) + b
}

func main() {
	var a myInt = 10 //myint 타입변수
	fmt.Println(a.add(30))
	var b int = 20
	fmt.Println(myInt(b).add(50))
}
