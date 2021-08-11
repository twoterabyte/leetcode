/*
 * @lc app=leetcode.cn id=413 lang=golang
 *
 * [413] 等差数列划分
 *
 * https://leetcode-cn.com/problems/arithmetic-slices/description/
 *
 * algorithms
 * Medium (66.53%)
 * Likes:    340
 * Dislikes: 0
 * Total Accepted:    63.5K
 * Total Submissions: 92.7K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 如果一个数列 至少有三个元素 ，并且任意两个相邻元素之差相同，则称该数列为等差数列。
 *
 *
 * 例如，[1,3,5,7,9]、[7,7,7,7] 和 [3,-1,-5,-9] 都是等差数列。
 *
 *
 *
 *
 * 给你一个整数数组 nums ，返回数组 nums 中所有为等差数组的 子数组 个数。
 *
 * 子数组 是数组中的一个连续序列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3,4]
 * 输出：3
 * 解释：nums 中有三个子等差数组：[1, 2, 3]、[2, 3, 4] 和 [1,2,3,4] 自身。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -1000
 *
 *
 *
 *
 */

// @lc code=start
// dp(i) 表示以 nums[i] 为结尾的等差数组个数
// dp(0)=0 dp(1)=0
// dp(i)=
// if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] then dp[i-1]+1
// else 0
// 求sum{dp(i) for 2<=i<len(nums)}
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	dp := 0
	sum := 0
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp = dp + 1
			sum += dp
			continue
		}
		dp = 0
	}
	return sum
}

// @lc code=end

