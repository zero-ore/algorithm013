package main

//盛最多的水 https://leetcode-cn.com/problems/container-with-most-water/submissions/

/**
 * 解法2
 * 1、i指向第一个元素，j指向第二个元素
 * 2、比较i、j的值，取最小值
 * 3、计算两根柱子围成的长方形面积，并取最大值
 */
func maxArea2(height []int) int {
	marea := 0
	i := 0
	j := len(height) - 1
	minHeight := 0
	for i < j {
		if height[i] < height[j] {
			minHeight = height[i]
			i = i + 1
		} else {
			minHeight = height[j]
			j = j - 1
		}
		area := (j - i + 1) * minHeight
		marea = Max(area, marea)
	}
	return marea
}

/**
 * 解法1：嵌套循环，数组下标i与每个元素a组成（x，y）坐标，分别穷举计算面积
 */
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	marea := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			area := (j - i) * Min(height[i], height[j])
			marea = Max(area, marea)
		}
	}
	return marea
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
