package main

//爬楼梯：https://leetcode-cn.com/problems/climbing-stairs/submissions/

/**
 * 爬楼梯,类似于斐波那契，求斐波那契的第n项值
 */
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	f1, f2, f3 := 1, 2, 0
	for i := 3; i <= n; i++ {
		f3 = f1 + f2
		f1, f2 = f2, f3
	}
	return f3
}
