package main

func maxArea2(height []int) int {
	marea := 0
	i := 0
	j := len(height) - 1
	minHeight := 0
	for i < j {
		if height[i] < height[j] {
			i = i + 1
			minHeight = height[i]
		} else {
			j = j - 1
			minHeight = height[j]
		}
		area := (j - i + 1) * minHeight
		marea = Max(area, marea)
	}
	return marea
}

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
