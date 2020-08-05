package main

import "fmt"

func main() {
	arr2 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	r1 := TwoSum(arr2, 9)
	fmt.Println(r1)

	r2 := TwoSum(arr2, 9)
	fmt.Println(r2)

}
