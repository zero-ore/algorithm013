package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 1
	num := nums[0]
	for _, v := range nums {
		if v != num {
			num = v
			nums[count] = v
			count += 1
		}
	}
	return count
}
