/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/description/
 *
 * algorithms
 * Hard (44.85%)
 * Likes:    493
 * Dislikes: 0
 * Total Accepted:    54.7K
 * Total Submissions: 120.8K
 * Testcase Example:  '[3,3,5,0,0,3,1,4]'
 *
 * 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
 * 
 * 设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
 * 
 * 注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 * 
 * 示例 1:
 * 
 * 输入: [3,3,5,0,0,3,1,4]
 * 输出: 6
 * 解释: 在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
 * 随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
 * 
 * 示例 2:
 * 
 * 输入: [1,2,3,4,5]
 * 输出: 4
 * 解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4
 * 。   
 * 注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。   
 * 因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
 * 
 * 
 * 示例 3:
 * 
 * 输入: [7,6,4,3,1] 
 * 输出: 0 
 * 解释: 在这个情况下, 没有交易完成, 所以最大利润为 0。
 * 
 */
// dpBuy1(i), dpSell1(i), dpBuy2(i), dpSell2(i) 分别表示在 prices[:i+1] 下，第一次买入后最大利润，第一次卖出后，第二次买入后，第二次卖出后
// dpBuy1(0) = -prices[0]; dpSell1(0) = 0;
// dpBuy2(0) = -prices[0]; dpSell2(0) = 0
// dpBuy2(0) = -prices[0] 初始值为这个了，这样一开始登录是没有第一次交易的利润，否则需要进行判断第一次交易成功后才给dpBuy2和dpSell2赋值初始值
// dpBuy1(i) = max{dpBuy1(i-1), -prices[i]}
// dpSell1(i) = max{dpSell1(i-1), dpBuy1(i-1)+prices[i]}
// dpBuy2(i) = max{dpBuy2(i-1), dpSell1(i-1)-prices[i]}
// dpSell2(i) = max{dpSell2(i-1), dpBuy2(i-1)+prices[i]}
// @lc code=start
func maxProfit(prices []int) int {
	dpBuy1 := - prices[0]
	dpSell1 := 0
	dpBuy2 := - prices[0]
	dpSell2 := 0
	for i := 1; i < len(prices); i++ {
		buy1 := max(dpBuy1, -prices[i])
		sell1 := max(dpSell1, dpBuy1+prices[i])
		buy2 := max(dpBuy2, dpSell1-prices[i])
		sell2 := max(dpSell2, dpBuy2+prices[i])
		dpBuy1, dpSell1, dpBuy2, dpSell2 = buy1, sell1, buy2, sell2
	}
	return max(dpSell1, dpSell2)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

