/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 *
 * https://leetcode-cn.com/problems/house-robber-ii/description/
 *
 * algorithms
 * Medium (39.35%)
 * Likes:    363
 * Dislikes: 0
 * Total Accepted:    54.7K
 * Total Submissions: 138.9K
 * Testcase Example:  '[2,3,2]'
 *
 * 
 * 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
 * 
 * 给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
 * 
 * 示例 1:
 * 
 * 输入: [2,3,2]
 * 输出: 3
 * 解释: 你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
 * 
 * 
 * 示例 2:
 * 
 * 输入: [1,2,3,1]
 * 输出: 4
 * 解释: 你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
 * 偷窃到的最高金额 = 1 + 3 = 4 。
 * 
 */

// 利用 198 题结果，考虑存在三种情况
// 1. 取第一个，转化为 198 题的 nums[:len(nums)-1]
// 2. 取最后一个，转化为 198 题的 nums[1:]
// 3. 两个都取，转化成 198 题的 nums[1:len(nums)-1]
// 3 包含在 1 和 2 中，所以只要考虑 1 和 2 两种情况
// @lc code=start
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(rob198(nums[:len(nums)-1]), rob198(nums[1:]))
}

func rob198(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp_i2 := 0 // dp[i-2]
	dp_i1 := 0 // dp[i-1]
	dp_i := nums[0] // dp[i]
	for i := 1; i < len(nums); i++ {
		dp_i1, dp_i2 = dp_i, dp_i1
		dp_i = max(dp_i1, dp_i2+nums[i])
	}
	return dp_i
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

