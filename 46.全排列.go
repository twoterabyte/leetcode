/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 *
 * https://leetcode-cn.com/problems/permutations/description/
 *
 * algorithms
 * Medium (76.66%)
 * Likes:    852
 * Dislikes: 0
 * Total Accepted:    176.1K
 * Total Submissions: 229.6K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个 没有重复 数字的序列，返回其所有可能的全排列。
 * 
 * 示例:
 * 
 * 输入: [1,2,3]
 * 输出:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 * 
 */

// @lc code=start
func permute(nums []int) [][]int {
	var result [][]int
	var backtrack func(track []int, isUsed []bool)
	backtrack = func(track []int, isUsed []bool) {
		if len(track) == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, track)
			result = append(result, tmp)
		}
		for i, num := range nums {
			if isUsed[i] {
				continue
			}
			track = append(track, num)
			isUsed[i] = true
			backtrack(track, isUsed)
			track = track[:len(track)-1]
			isUsed[i] = false
		}
	}
	track := make([]int, 0, len(nums))
	isUsed := make([]bool, len(nums))
	backtrack(track, isUsed)
	return result
}
// @lc code=end

