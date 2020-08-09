package main

// 两数之和 ：https://leetcode-cn.com/problems/two-sum/

/**
 * 暴力解法：遍历每个元素 x，并查找是否存在一个值与 target - x 相等的目标元素
 */
func TwoSum(nums []int, target int) []int {
	for i, v := range nums {
		for k := i + 1; k < len(nums); k++ {
			if nums[k] == target-v {
				return []int{i, k}
			}
		}
	}
	return []int{}
}

/**
 * 解法2：
 * 1、遍历nums中的每个元素，每次遍历中（i），将每个元素值作为key，每个元素的数组下标作为value存在一个map中，判断map中是否【target-元素】值，存在的话，即获取该map元素的值和当前遍历的i
 */
func TwoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if value, exist := m[target-v]; exist {
			return []int{i, value}
		}
		m[v] = i
	}
	return []int{}
}
