/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
* [322] 零钱兑换
 *
 * https://leetcode-cn.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (40.45%)
 * Likes:    709
 * Dislikes: 0
 * Total Accepted:    106.7K
 * Total Submissions: 263.8K
 * Testcase Example:  '[1,2,5]\n11'
 *
 * 给定不同面额的硬币 coins 和一个总金额
 * amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: coins = [1, 2, 5], amount = 11
 * 输出: 3 
 * 解释: 11 = 5 + 5 + 1
 * 
 * 示例 2:
 * 
 * 输入: coins = [2], amount = 3
 * 输出: -1
 * 
 * 
 * 
 * 说明:
 * 你可以认为每种硬币的数量是无限的。
 */

// @lc code=start
// 函数 dp(n) 表示总金额n需要的最少硬币数
// dp(0) = 0, dp(n)=-1 for n < 0
// dp(n) = -1 if dp(n - i) = -1 else min{1 + dp(n - i) for i in all coins}
import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)
	for i := 1; i <= amount; i++ {
		dp[i] = int(math.MaxInt32)
		for _, coin := range coins {
			index := i - coin
			if (index) < 0 {
				continue
			}
			if dp[i] > dp[index] + 1 {
				dp[i] = dp[index] + 1
			}
		}
	}
	if dp[amount] != math.MaxInt32 {
		return dp[amount]
	}
	return -1
}
// @lc code=end

