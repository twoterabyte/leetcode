/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 *
 * https://leetcode-cn.com/problems/burst-balloons/description/
 *
 * algorithms
 * Hard (66.73%)
 * Likes:    478
 * Dislikes: 0
 * Total Accepted:    28.3K
 * Total Submissions: 42.4K
 * Testcase Example:  '[3,1,5,8]'
 *
 * 有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
 * 
 * 现在要求你戳破所有的气球。如果你戳破气球 i ，就可以获得 nums[left] * nums[i] * nums[right] 个硬币。 这里的
 * left 和 right 代表和 i 相邻的两个气球的序号。注意当你戳破了气球 i 后，气球 left 和气球 right 就变成了相邻的气球。
 * 
 * 求所能获得硬币的最大数量。
 * 
 * 说明:
 * 
 * 
 * 你可以假设 nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。
 * 0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
 * 
 * 
 * 示例:
 * 
 * 输入: [3,1,5,8]
 * 输出: 167 
 * 解释: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
 * coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167
 * 
 * 
 */

// dp(i,j) 表示戳破 nums[i+1:j] 气球可以获得最大硬币数量
// dp(k,k)=0
// dp(i,j)=max{dp(i,k)+dp(k,j)+nums[i]*nums[k]*nums[j] for i<k<j}
// 最后戳破 k，需要先戳破 (i,k) 和 (k,j)

// @lc code=start
func maxCoins(nums []int) int {
	// 前后插入 1，方便运算
	newNums := make([]int, 0, len(nums)+2)
	newNums = append(newNums, 1)
	newNums = append(newNums, nums...)
	newNums = append(newNums, 1)
	// dp 数组初始化
	dp := make([][]int, len(newNums))
	for i := range dp {
		dp[i] = make([]int, len(newNums))
	}
	for i := len(newNums) - 2; i >= 0; i-- {
		for j := i + 1; j < len(newNums); j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j],
					dp[i][k]+dp[k][j]+newNums[i]*newNums[j]*newNums[k])
			}
		}
	}
	return dp[0][len(newNums)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

