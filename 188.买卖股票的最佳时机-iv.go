/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/description/
 *
 * algorithms
 * Hard (30.29%)
 * Likes:    462
 * Dislikes: 0
 * Total Accepted:    62.8K
 * Total Submissions: 171.4K
 * Testcase Example:  '2\n[2,4,1]'
 *
 * 给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。
 * 
 * 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
 * 
 * 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：k = 2, prices = [2,4,1]
 * 输出：2
 * 解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
 * 
 * 示例 2：
 * 
 * 
 * 输入：k = 2, prices = [3,2,6,5,0,3]
 * 输出：7
 * 解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
 * ⁠    随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 =
 * 3 。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 
 * 0 
 * 0 
 * 
 * 
 */

// dp_buy(i,j) 表示在 prices[:i+1] 时，最多完成 j 笔交易，且持有股票的情况下，的最大利润
// dp_sell(i,j) 表示在 prices[:i+1] 时，最多完成 j 笔交易，且没持有股票的情况下，的最大利润
// 为了方便理解，我们在买入时候，就认为发生了一次交易，卖出的时候不算
// dp_buy(0,j) = -prices[0], dp_sell(0,j) = 0
// dp_buy(i,0) = 0, dp_sell(i,0) = 0
// dp_buy(0,0) 是一个不可能存在的状态，所以这个值不用关心，动态规划里面也没用到这个状态
// dp_buy(i,j) = max{dp_buy(i-1,j), dp_sell(i-1,j-1)-prices[i]}
//     其中 dp_buy(i-1,j) 表示，prices[:i] 买后已经有 j 交易，那么当前则不能在买了
//     dp_sell(i-1,j-1)-prices[i] 表示 prices[:i] 卖后只有 j-1 交易，那么当前可以继续买
// dp_sell(i,j) = max{dp_sell(i-1,j), dp_buy(i-1,j)+prices[i]}
// 求dp_sell(len(prices)-1,k)
// @lc code=start
func maxProfit(k int, prices []int) int {
	if k == 0 || len(prices) == 0 {
		return 0
	}
	dpBuy := make([]int, k + 1)
	dpBuyPre := make([]int, k + 1)
	dpSell := make([]int, k + 1)
	dpSellPre := make([]int, k + 1)
	// 初始化状态
	for i := range dpBuyPre {
		dpBuyPre[i] = 0 - prices[0]
	}
	dpBuyPre[0] = 0
	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			dpBuy[j] = max(dpBuyPre[j], dpSellPre[j-1]-prices[i])
			dpSell[j] = max(dpSellPre[j], dpBuyPre[j]+prices[i])
		}
		dpBuyPre, dpSellPre = dpBuy, dpSell
	}
	return dpSellPre[k]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

