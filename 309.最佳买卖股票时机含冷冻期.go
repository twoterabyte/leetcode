/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/
 *
 * algorithms
 * Medium (57.10%)
 * Likes:    706
 * Dislikes: 0
 * Total Accepted:    78K
 * Total Submissions: 135.2K
 * Testcase Example:  '[1,2,3,0,2]'
 *
 * 给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​
 * 
 * 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
 * 
 * 
 * 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 * 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
 * 
 * 
 * 示例:
 * 
 * 输入: [1,2,3,0,2]
 * 输出: 3 
 * 解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
 * 
 */

// dpBuy(i) 表示在 prices[:i+1] 下，持有股票的的最大利润
// dpSell(i) 表示在 prices[:i+1] 下，没持有股票的最大利润
// dpBuy(0) = -prices[0]; dpSell(0) = 0
// dpSell(-1) = 0
// dpBuy(i) = max{dpBuy(i-1), dpSell(i-2)-prices[i]}
// dpSell(i) = max{dpSell(i-1), dpBuy(i-1)+prices[i]}
// 求 dpSell(len(prices)-1)
// @lc code=start
func maxProfit(prices []int) int {
	dpBuy := 0 - prices[0]
	dpSell := 0
	dpSellPre := 0
	for i := 1; i < len(prices); i++ {
		buy := max(dpBuy, dpSellPre-prices[i])
		sell := max(dpSell, dpBuy+prices[i])
		dpBuy, dpSell, dpSellPre = buy, sell, dpSell
	}
	return dpSell
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

