/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 *
 * https://leetcode-cn.com/problems/combinations/description/
 *
 * algorithms
 * Medium (74.42%)
 * Likes:    334
 * Dislikes: 0
 * Total Accepted:    70.5K
 * Total Submissions: 94.6K
 * Testcase Example:  '4\n2'
 *
 * 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
 * 
 * 示例:
 * 
 * 输入: n = 4, k = 2
 * 输出:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 * 
 */

// @lc code=start
func combine(n int, k int) [][]int {
	var result [][]int
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}
	track := make([]int, k)
	var dfs func(start, idxTrack int)
	dfs = func(start, idxTrack int) {
		if idxTrack == k {
			tmp := make([]int, k)
			copy(tmp, track)
			result = append(result, tmp)
			return
		}
		// 为什么这循环，考虑组合，必须满足逆序数，后面的数字一定要大于前面
		// i 从 start 开始，
		for i := start; i < n; i++ {
			track[idxTrack] = nums[i]
			// 从 i + 1 开始
			dfs(i+1, idxTrack+1)
		}
	}
	dfs(0, 0)
	return result
}
// @lc code=end

