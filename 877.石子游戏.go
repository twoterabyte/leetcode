/*
 * @lc app=leetcode.cn id=877 lang=golang
 *
 * [877] 石子游戏
 *
 * https://leetcode-cn.com/problems/stone-game/description/
 *
 * algorithms
 * Medium (69.68%)
 * Likes:    163
 * Dislikes: 0
 * Total Accepted:    18.1K
 * Total Submissions: 25.9K
 * Testcase Example:  '[5,3,4,5]'
 *
 * 亚历克斯和李用几堆石子在做游戏。偶数堆石子排成一行，每堆都有正整数颗石子 piles[i] 。
 * 
 * 游戏以谁手中的石子最多来决出胜负。石子的总数是奇数，所以没有平局。
 * 
 * 亚历克斯和李轮流进行，亚历克斯先开始。 每回合，玩家从行的开始或结束处取走整堆石头。
 * 这种情况一直持续到没有更多的石子堆为止，此时手中石子最多的玩家获胜。
 * 
 * 假设亚历克斯和李都发挥出最佳水平，当亚历克斯赢得比赛时返回 true ，当李赢得比赛时返回 false 。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入：[5,3,4,5]
 * 输出：true
 * 解释：
 * 亚历克斯先开始，只能拿前 5 颗或后 5 颗石子 。
 * 假设他取了前 5 颗，这一行就变成了 [3,4,5] 。
 * 如果李拿走前 3 颗，那么剩下的是 [4,5]，亚历克斯拿走后 5 颗赢得 10 分。
 * 如果李拿走后 5 颗，那么剩下的是 [3,4]，亚历克斯拿走后 4 颗赢得 9 分。
 * 这表明，取前 5 颗石子对亚历克斯来说是一个胜利的举动，所以我们返回 true 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= piles.length <= 500
 * piles.length 是偶数。
 * 1 <= piles[i] <= 500
 * sum(piles) 是奇数。
 * 
 * 
 */

// 这是一个先手必胜的问题，先手可以保证自己拿到的都是奇数堆或者偶数堆
// dp[i][j].first 表示当从 i 到 j 堆进行游戏，先手获得的最大石头数
// dp[i][j].last 表示当从 i 到 j 堆进行游戏，后手获得的最大石头数
// dp[i][i].first=piles[i]; dp[i][j].last=0
// dp[i][j].first=max{piles[i]+dp[i+1][j].last, piles[j]+dp[i][j-1].last}
// dp[i][j].last=
//   dp[i+1][j].first if dp[i][j].first select piles[i]
//   dp[i][j-1].first if dp[i][j].first select piles[j]
// @lc code=start
type stoneNum struct {
	fisrt int
	last  int
}

func stoneGame(piles []int) bool {
	// dp 初始化，这里还可做状态压缩，太简单了。所以不做了
	dp := make([][]stoneNum, len(piles))
	for i := range dp {
		dp[i] = make([]stoneNum, len(piles))
		dp[i][i].fisrt = piles[i]
	}
	for i := len(piles) - 2; i >= 0; i-- {
		for j := i + 1; j < len(piles); j++ {
			selectLeft := piles[i] + dp[i+1][j].last
			selectRight := piles[j] + dp[i][j-1].last
			if selectLeft > selectRight {
				// 先手选择左边
				dp[i][j].fisrt = selectLeft
				dp[i][j].last = dp[i+1][j].fisrt
			} else {
				// 先手选择右边
				dp[i][j].fisrt = selectRight
				dp[i][j].last = dp[i][j-1].fisrt
			}
		}
	}
	return dp[0][len(piles)-1].fisrt > dp[0][len(piles)-1].last
}
// @lc code=end

