package main

/**
 * 移动零
 */
func moveZeroes(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}
	i := 0
	j := 0
	for ; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
		j++
	}
}
