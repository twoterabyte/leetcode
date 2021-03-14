/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 *
 * https://leetcode-cn.com/problems/subsets/description/
 *
 * algorithms
 * Medium (77.77%)
 * Likes:    720
 * Dislikes: 0
 * Total Accepted:    123.9K
 * Total Submissions: 159.2K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
 * 
 * 说明：解集不能包含重复的子集。
 * 
 * 示例:
 * 
 * 输入: nums = [1,2,3]
 * 输出:
 * [
 * ⁠ [3],
 * [1],
 * [2],
 * [1,2,3],
 * [1,3],
 * [2,3],
 * [1,2],
 * []
 * ]
 * 
 */

// @lc code=start
func subsets(nums []int) [][]int {
	var result [][]int
	track := make([]int, len(nums))
	var dfs func(start, idxTrack int)
	dfs = func(start, idxTrack int) {
		tmp := make([]int, idxTrack)
		copy(tmp, track[:idxTrack])
		result = append(result, tmp)
		// 到达递归结尾
		if idxTrack == len(nums) {
			return
		}
		// 剪枝，后面的数字不能出现前面的数字前面，保证不重复
		for i := start; i < len(nums); i++ {
			track[idxTrack] = nums[i]
			dfs(i+1, idxTrack+1)
		}
	}
	dfs(0, 0)
	return result
}
// @lc code=end

