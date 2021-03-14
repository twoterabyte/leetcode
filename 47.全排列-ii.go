/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 *
 * https://leetcode-cn.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (59.59%)
 * Likes:    381
 * Dislikes: 0
 * Total Accepted:    82K
 * Total Submissions: 137.4K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个可包含重复数字的序列，返回所有不重复的全排列。
 * 
 * 示例:
 * 
 * 输入: [1,1,2]
 * 输出:
 * [
 * ⁠ [1,1,2],
 * ⁠ [1,2,1],
 * ⁠ [2,1,1]
 * ]
 * 
 */

// @lc code=start
import "sort"
func permuteUnique(nums []int) [][]int {
	var result [][]int
	// 先对 nums 进行排序，方便去重
	sort.Ints(nums)
	var backtrack func(track []int, isUsed []bool)
	backtrack = func(track []int, isUsed []bool) {
		// 到达结果
		if len(track) == len(nums) {
			// slice 需要 copy 一个出来，否则会被后面修改到
			tmp := make([]int, len(nums))
			copy(tmp, track)
			result = append(result, tmp)
		}
		for i, num := range nums {
			// 判断是否已经被选择过了
			if isUsed[i] {
				continue
			}
			// 当前元素和上一个元素一样，存在两种情况
			// 如果上一个元素被使用了，说明进入更深的节点，不再同一层遍历
			// 如果上一个元素没有使用，那就是因为再同一层回溯被撤销了
			// 这里的 !isUsed[i-1] 是核心的部分
			if i > 0 && num == nums[i-1] && !isUsed[i-1] {
				continue
			}
			// 增加回溯路径里面
			track = append(track, num)
			// 标记已经选择过了
			isUsed[i] = true
			backtrack(track, isUsed)
			// 从回溯路径里面删除
			track = track[:len(track)-1]
			// 标记未被选择过
			isUsed[i] = false
		}
	}
	track := make([]int, 0, len(nums))
	isUsed := make([]bool, len(nums))
	backtrack(track, isUsed)
	return result
}
// @lc code=end

