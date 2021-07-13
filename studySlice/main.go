package main

import "fmt"

func changeArray(array2 [5]int) {
	array2[2] = 200
}

func changeSlice(slice []int) {
	slice[2] = 200
}
func main() {

	/*var slice =[]int{1,2,3}
	slice2 :=append(slice,4)

	fmt.Println(slice)
	fmt.Println(slice2)*/

	/*	var slice []int

		for i := 0; i <= 10; i++ {
			slice = append(slice, i)
		}
		slice2 := append(slice, 11, 12)
		fmt.Println(slice2)*/

	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	for i := 0; i <= 4; i++ {
		fmt.Print(array[i])
	}
	fmt.Println()
	for i := 0; i <= 6; i++ {
		slice = append(slice, i)
		fmt.Print(slice)
	}
	fmt.Println()
	for i := 0; i <= 6; i++ {
		slice = append(slice, i)
		fmt.Print(slice[i])
	}

	fmt.Println()
	changeArray(array)
	changeSlice(slice)
	fmt.Println("array? = ", array)
	fmt.Println("slice?=", slice)

	fmt.Println()
	sliceTest1 := make([]int, 3, 5)
	sliceTest2 := append(sliceTest1, 4, 5)
	fmt.Println("SliceTest1 :", sliceTest1, len(sliceTest1), cap(sliceTest1))
	fmt.Println("SliceTest2 :", sliceTest2, len(sliceTest2), cap(sliceTest2))

	sliceTest1[1] = 100 //slicetest2 까지 값이 변경 .

	fmt.Println("After change second element ")
	fmt.Println("SliceTest1 :", sliceTest1, len(sliceTest1), cap(sliceTest1))
	fmt.Println("SliceTest2 :", sliceTest2, len(sliceTest2), cap(sliceTest2))

	sliceTest1 = append(sliceTest1, 500)
	fmt.Println("After append 500 ")
	fmt.Println("SliceTest1 :", sliceTest1, len(sliceTest1), cap(sliceTest1))
	fmt.Println("SliceTest2 :", sliceTest2, len(sliceTest2), cap(sliceTest2))

}
