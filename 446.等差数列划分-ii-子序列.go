/*
 * @lc app=leetcode.cn id=446 lang=golang
 *
 * [446] 等差数列划分 II - 子序列
 *
 * https://leetcode-cn.com/problems/arithmetic-slices-ii-subsequence/description/
 *
 * algorithms
 * Hard (38.12%)
 * Likes:    196
 * Dislikes: 0
 * Total Accepted:    15.6K
 * Total Submissions: 29.2K
 * Testcase Example:  '[2,4,6,8,10]'
 *
 * 给你一个整数数组 nums ，返回 nums 中所有 等差子序列 的数目。
 *
 * 如果一个序列中 至少有三个元素 ，并且任意两个相邻元素之差相同，则称该序列为等差序列。
 *
 *
 * 例如，[1, 3, 5, 7, 9]、[7, 7, 7, 7] 和 [3, -1, -5, -9] 都是等差序列。
 * 再例如，[1, 1, 2, 5, 7] 不是等差序列。
 *
 *
 * 数组中的子序列是从数组中删除一些元素（也可能不删除）得到的一个序列。
 *
 *
 * 例如，[2,5,10] 是 [1,2,1,2,4,1,5,10] 的一个子序列。
 *
 *
 * 题目数据保证答案是一个 32-bit 整数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,4,6,8,10]
 * 输出：7
 * 解释：所有的等差子序列为：
 * [2,4,6]
 * [4,6,8]
 * [6,8,10]
 * [2,4,6,8]
 * [4,6,8,10]
 * [2,4,6,8,10]
 * [2,6,10]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [7,7,7,7,7]
 * 输出：16
 * 解释：数组中的任意子序列都是等差子序列。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1  <= nums.length <= 1000
 * -2^31 <= nums[i] <= 2^31 - 1
 *
 *
 */

// @lc code=start

// 弱等差序列表示一个序列中至少有两个元素 ，并且任意两个相邻元素之差相同
// dp(i)(j) 表示以 nums[i] 为结尾，公差为 j 的弱等差序列个数
// 则是求 sum{dp(i)(j)-1 for 2<=i<len(nums) && dp(i)(j) > 1}
// dp(0)(j) = 0
// dp(i)(nums[i]-nums[j]) += {dp(j)(nums[i]-nums[j])+1 for 0<j<i}
// 则求的等差序列个数则是在，遍历 i 的过程中，求和 dp(j)(nums[i]-nums[j])
// 表示 nums[i] 可以和以 nums[j] 为结尾以 nums[i]-nums[j] 为公差的弱等差序列拼接成等差序列
// 时间复杂度：O(n^2)
// 空间复杂度：O(nm)
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	dp := make([]map[int]int, len(nums))
	sum := 0
	dp[0] = make(map[int]int)
	for i := 1; i < len(nums); i++ {
		dp[i] = make(map[int]int)
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			res := dp[j][diff]
			dp[i][diff] += res + 1
			// 遍历的过程顺便把 sum 也给求了
			sum += res
		}
	}
	return sum
}

// @lc code=end

