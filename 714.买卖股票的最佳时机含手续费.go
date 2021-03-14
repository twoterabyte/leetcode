/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/
 *
 * algorithms
 * Medium (68.15%)
 * Likes:    428
 * Dislikes: 0
 * Total Accepted:    68.8K
 * Total Submissions: 98.3K
 * Testcase Example:  '[1,3,2,8,4,9]\n2'
 *
 * 给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。
 * 
 * 你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
 * 
 * 返回获得利润的最大值。
 * 
 * 注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
 * 
 * 示例 1:
 * 
 * 输入: prices = [1, 3, 2, 8, 4, 9], fee = 2
 * 输出: 8
 * 解释: 能够达到的最大利润:  
 * 在此处买入 prices[0] = 1
 * 在此处卖出 prices[3] = 8
 * 在此处买入 prices[4] = 4
 * 在此处卖出 prices[5] = 9
 * 总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
 * 
 * 注意:
 * 
 * 
 * 0 < prices.length <= 50000.
 * 0 < prices[i] < 50000.
 * 0 <= fee < 50000.
 * 
 * 
 */

// dp_buy(i) 表示第 i 天持有股票时的最大利润
// dp_sell(i) 表示第 i 天没有持有股票时的最大利润
// 为了方便计算，在买入股票时就把手续费也进行扣除
// dp_buy(0) = -prices[0]-fee
// dp_sell(0) = 0
// dp_buy(i) = max{dp_buy(i-1), dp_sell(i-1)-prices[i]-fee}
// dp_sell(i) = max{dp_sell(i-1), dp_buy(i-1)+prices[i]}
// 求dp_sell(len(prices)-1)
// @lc code=start
func maxProfit(prices []int, fee int) int {
	// 初始化状态
	dpBuy := 0 - prices[0] - fee
	dpSell := 0
	for _, price := range prices[1:] {
		buy := max(dpBuy, dpSell-price-fee)
		sell := max(dpSell, dpBuy+price)
		dpBuy, dpSell = buy, sell
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

