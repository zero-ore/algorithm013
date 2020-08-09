package main

import (
	"sort"
)

/**
 * 三数之和：https://leetcode-cn.com/problems/3sum/
 * 如果和大于 0，那就说明 right 的值太大，需要左移。如果和小于 0，那就说明 left 的值太小，需要右移
 */

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for {
			if j >= k {
				break
			}
			sum := nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
				continue
			}
			if sum < 0 {
				j++
				continue
			}
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			}
		}
	}
	return res
}
