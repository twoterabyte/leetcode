/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/description/
 *
 * algorithms
 * Easy (62.71%)
 * Likes:    837
 * Dislikes: 0
 * Total Accepted:    213.7K
 * Total Submissions: 339.7K
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
 * 
 * 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
 * 
 * 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: [7,1,5,3,6,4]
 * 输出: 7
 * 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
 * 随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
 * 
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
 * 示例 3:
 * 
 * 输入: [7,6,4,3,1]
 * 输出: 0
 * 解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= prices.length <= 3 * 10 ^ 4
 * 0 <= prices[i] <= 10 ^ 4
 * 
 * 
 */

// @lc code=start

// 实际就是求所有递增的子串，并用最后一个减去第一个
// 假如有递增子串 array[i:j]，则有
// array[j-1]-array[i]
// = array[j-1]-array[j-2]
// + array[j-2]-array[j-3]
// + ......
// + array[i+2]-array[i+1]
// + array[i+1]-array[i]
func maxProfit(prices []int) int {
	resust := 0
	front := prices[0]
	for _, price := range prices[1:] {
		if price > front {
			resust += price - front
		}
		front = price
	}
	return resust
}

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// 虽然过了，但是刚好卡着过了
// func maxProfit(prices []int) int {
// 	// 表示 prices[i:] 进行买卖的最大收益
// 	maxProfixs := make([]int, len(prices)+1)
// 	for i := range maxProfixs {
// 		maxProfixs[i] = -1
// 	}
// 	// 没有交易则收益为 0
// 	maxProfixs[len(prices)] = 0
// 	// 只有一次交易收益只能为 0
// 	maxProfixs[len(prices)-1] = 0
// 	// 返回 prices[i:] 进行买卖的最大收益
// 	var subMaxProfit func(i int) int
// 	subMaxProfit = func(i int) int {
// 		// 先查表
// 		profix := maxProfixs[i]
// 		if profix != -1 {
// 			return profix
// 		}
// 		// 今天不买
// 		profix = subMaxProfit(i+1)
// 		// 今天买，需要在后面找一天卖掉
// 		for j := i + 1; j < len(prices); j++ {
// 			if prices[j] > prices[i] {
// 				profix = max(profix, prices[j]-prices[i]+subMaxProfit(j+1))
// 			}
// 		}
// 		maxProfixs[i] = profix
// 		return profix
// 	}
// 	return subMaxProfit(0)
// }
// 超时了
// func maxProfit(prices []int) int {
// 	var dfs func(i, profix, priceBuy int)
// 	dfs = func(i, profix, priceBuy int) {
// 		if i == len(prices) {
// 			if profix > max {
// 				max = profix
// 			}
// 			return
// 		}
// 		// 当前不买不卖
// 		dfs(i+1, profix, priceBuy);
// 		// 如果当前没买，可以选择买
// 		if priceBuy == -1 {
// 			dfs(i+1, profix-prices[i], prices[i])
// 		}
// 		// 如果当前价格大于买入价，可以选择卖
// 		if priceBuy != -1 && priceBuy < prices[i] {
// 			dfs(i+1, profix+prices[i], -1)
// 		}
// 	}
// 	dfs(0, 0, -1)
// 	return max
// }
// @lc code=end

