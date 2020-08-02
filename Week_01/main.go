package main

import "fmt"

func main() {
	arr1 := []int{0, 2, 3, 2, 0, 8, 9}
	moveZeroes(arr1)
	fmt.Println(arr1)

	arr2 := []int{0, 0, 0, 2, 0, 8, 9}
	moveZeroes(arr2)
	fmt.Println(arr2)
}
