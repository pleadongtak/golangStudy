package main

import "fmt"

func changeArray(array [5]int) {
	array[2] = 500
}
func main() {
	/*array := [5]int{1,2,3,4,5}
	changeArray(array)

	array[0]=100
	array[1]=200
	fmt.Println("Array  ?", array)
	slice := array[1:2]
	fmt.Println("Array  ?", array)
	fmt.Println("slice? ",slice,len(slice),cap(slice))

	array[1]=100
	fmt.Println("Array  ?", array)
	*/

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))

	for i, v := range slice1 {
		slice2[i] = v
	}
	slice1[1] = 100
	fmt.Println(slice1)
	fmt.Println(slice2)

}
