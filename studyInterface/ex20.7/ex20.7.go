package main

import "fmt"

type Stringer interface { //인터페이스
	String() string
}
type Student struct {
	Age int
}

func (s *Student) String() string {
	return fmt.Sprintf("Student age %d", s.Age)
}

func PrintAge(stringer Stringer) {
	s := stringer.(*Student)
	fmt.Printf("age:%d", s.Age)
}

func main() {
	s := &Student{15}

	PrintAge(s)
}
