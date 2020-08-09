package main

import "fmt"

func main() {
	arr22 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	r1 := TwoSum(arr22, 9)
	fmt.Println(r1)

	s1 := "()[]{}"
	r2 := isValid(s1)
	fmt.Println(r2)

}
