package main

import "fmt"

func main() {
	//移动零
	arr1 := []int{0, 2, 3, 2, 0, 8, 9}
	moveZeroes(arr1)
	fmt.Println(arr1)

	//盛最多的水
	arr2 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	m1 := maxArea(arr2)
	fmt.Println(m1)

	m2 := maxArea2(arr2)
	fmt.Println(m2)

	//
	arr3 := []int{1, 2, 6, 3, 5, 4, 8, 3, 7}
	count := removeDuplicates(arr3)
	fmt.Println(count)

}
