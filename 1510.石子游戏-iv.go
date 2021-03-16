/*
 * @lc app=leetcode.cn id=1510 lang=golang
 *
 * [1510] 石子游戏 IV
 *
 * https://leetcode-cn.com/problems/stone-game-iv/description/
 *
 * algorithms
 * Hard (48.14%)
 * Likes:    9
 * Dislikes: 0
 * Total Accepted:    2.3K
 * Total Submissions: 4.5K
 * Testcase Example:  '1'
 *
 * Alice 和 Bob 两个人轮流玩一个游戏，Alice 先手。
 * 
 * 一开始，有 n 个石子堆在一起。每个人轮流操作，正在操作的玩家可以从石子堆里拿走 任意 非零 平方数 个石子。
 * 
 * 如果石子堆里没有石子了，则无法操作的玩家输掉游戏。
 * 
 * 给你正整数 n ，且已知两个人都采取最优策略。如果 Alice 会赢得比赛，那么返回 True ，否则返回 False 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 1
 * 输出：true
 * 解释：Alice 拿走 1 个石子并赢得胜利，因为 Bob 无法进行任何操作。
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 2
 * 输出：false
 * 解释：Alice 只能拿走 1 个石子，然后 Bob 拿走最后一个石子并赢得胜利（2 -> 1 -> 0）。
 * 
 * 示例 3：
 * 
 * 
 * 输入：n = 4
 * 输出：true
 * 解释：n 已经是一个平方数，Alice 可以一次全拿掉 4 个石子并赢得胜利（4 -> 0）。
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：n = 7
 * 输出：false
 * 解释：当 Bob 采取最优策略时，Alice 无法赢得比赛。
 * 如果 Alice 一开始拿走 4 个石子， Bob 会拿走 1 个石子，然后 Alice 只能拿走 1 个石子，Bob 拿走最后一个石子并赢得胜利（7
 * -> 3 -> 2 -> 1 -> 0）。
 * 如果 Alice 一开始拿走 1 个石子， Bob 会拿走 4 个石子，然后 Alice 只能拿走 1 个石子，Bob 拿走最后一个石子并赢得胜利（7
 * -> 6 -> 2 -> 1 -> 0）。
 * 
 * 示例 5：
 * 
 * 
 * 输入：n = 17
 * 输出：false
 * 解释：如果 Bob 采取最优策略，Alice 无法赢得胜利。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 10^5
 * 
 * 
 */

// dp(i) 表示有 i 个石子时，先手是否取胜
// dp(0) = false
// dp(i) = true if exist dp(i-j^2) = false for j^2<=i
// @lc code=start

func winnerSquareGame(n int) bool {
	dp := make([]bool, n+1)
	dp[0] = false
	for i := 0; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			if !dp[i-j*j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

// var dpTable = make([]bool, 100001)
// func winnerSquareGame(n int) bool {
// 	exist := make([]bool, n+1)
// 	exist[0] = true
// 	var dp func(n int) bool
// 	dp = func(i int) bool {
// 		if exist[i] {
// 			return dpTable[i]
// 		}
// 		for j := 1; j*j <= i; j++ {
// 			if !dp(i-j*j) {
// 				exist[i] = true
// 				dpTable[i] = true
// 				return true
// 			}
// 		}
// 		exist[i] = false
// 		dpTable[i] = false
// 		return false
// 	}
// 	return dp(n)
// }

// func winnerSquareGame(n int) bool {
// 	dpTable := make(map[int]bool)
// 	dpTable[0] = false
// 	var dp func(n int) bool
// 	dp = func(i int) bool {
// 		if result, ok := dpTable[i]; ok {
// 			return result
// 		}
// 		for j := 1; j*j <= i; j++ {
// 			if !dp(i-j*j) {
// 				dpTable[i] = true
// 				return true
// 			}
// 		}
// 		dpTable[i] = false
// 		return false
// 	}
// 	return dp(n)
// }
// @lc code=end

