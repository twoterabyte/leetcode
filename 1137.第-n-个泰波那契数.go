/*
 * @lc app=leetcode.cn id=1137 lang=golang
 *
 * [1137] 第 N 个泰波那契数
 */

// @lc code=start
func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	t0, t1, t2 := 0, 1, 1
	for i := 3; i <= n; i++ {
		t := t0 + t1 + t2
		t0, t1, t2 = t1, t2, t
	}
	return t2
}

// @lc code=end

