/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 *
 * https://leetcode-cn.com/problems/partition-equal-subset-sum/description/
 *
 * algorithms
 * Medium (48.49%)
 * Likes:    340
 * Dislikes: 0
 * Total Accepted:    42.8K
 * Total Submissions: 87.9K
 * Testcase Example:  '[1,5,11,5]'
 *
 * 给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
 * 
 * 注意:
 * 
 * 
 * 每个数组中的元素不会超过 100
 * 数组的大小不会超过 200
 * 
 * 
 * 示例 1:
 * 
 * 输入: [1, 5, 11, 5]
 * 
 * 输出: true
 * 
 * 解释: 数组可以分割成 [1, 5, 5] 和 [11].
 * 
 * 
 * 
 * 
 * 示例 2:
 * 
 * 输入: [1, 2, 3, 5]
 * 
 * 输出: false
 * 
 * 解释: 数组不能分割成两个元素和相等的子集.
 * 
 * 
 * 
 * 
 */

// dp[i][j] 表示对于前 i 个元素，和为 j, 存在分割方式
// dp[0][:]=0; dp[0][0]=1
// dp[i][j]=max{dp[i-1][j], dp[i-1][j-nums[i-1]]}
// @lc code=start
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	half := sum / 2
	dp := make([]bool, half+1)
	dp[0] = true
	for i := 1; i <= len(nums); i++ {
		for j := half; j >= 0; j-- {
			if j-nums[i-1] >= 0 {
				dp[j] = dp[j] || dp[j-nums[i-1]]
			}
		}
	}
	return dp[half]
}
// todo dp := make([]int, half+1) 用计数的方式居然出错了
// @lc code=end

