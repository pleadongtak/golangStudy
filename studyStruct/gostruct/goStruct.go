package main

import "fmt"

type Duck struct {
}

func (d Duck) sound() {
	fmt.Println("나ㅏ는 오리 꿱")
}

func (d Duck) taste() {
	fmt.Println("맛있는 오리껍질 ")

}

type Fish struct {
}

func (f Fish) sound() {
	fmt.Println("물고기 푸다닥")
}

func (f Fish) taste() {
	fmt.Println("물고기 고기 지글지글 ")
}

type hunting interface { //인터페이스는 메소드로 묶임
	sound()
	taste()
}

func inTheriver(h hunting) {
	h.sound()
	h.taste()
}
func main() {

}
