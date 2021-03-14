/*
 * @lc app=leetcode.cn id=1140 lang=golang
 *
 * [1140] 石子游戏 II
 *
 * https://leetcode-cn.com/problems/stone-game-ii/description/
 *
 * algorithms
 * Medium (62.91%)
 * Likes:    58
 * Dislikes: 0
 * Total Accepted:    3.3K
 * Total Submissions: 5.2K
 * Testcase Example:  '[2,7,9,4,4]'
 *
 * 亚历克斯和李继续他们的石子游戏。许多堆石子 排成一行，每堆都有正整数颗石子 piles[i]。游戏以谁手中的石子最多来决出胜负。
 * 
 * 亚历克斯和李轮流进行，亚历克斯先开始。最初，M = 1。
 * 
 * 在每个玩家的回合中，该玩家可以拿走剩下的 前 X 堆的所有石子，其中 1 <= X <= 2M。然后，令 M = max(M, X)。
 * 
 * 游戏一直持续到所有石子都被拿走。
 * 
 * 假设亚历克斯和李都发挥出最佳水平，返回亚历克斯可以得到的最大数量的石头。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入：piles = [2,7,9,4,4]
 * 输出：10
 * 解释：
 * 如果亚历克斯在开始时拿走一堆石子，李拿走两堆，接着亚历克斯也拿走两堆。在这种情况下，亚历克斯可以拿到 2 + 4 + 4 = 10 颗石子。 
 * 如果亚历克斯在开始时拿走两堆石子，那么李就可以拿走剩下全部三堆石子。在这种情况下，亚历克斯可以拿到 2 + 7 = 9 颗石子。
 * 所以我们返回更大的 10。 
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= piles.length <= 100
 * 1 <= piles[i] <= 10 ^ 4
 * 
 * 
 */
// dp[i][j] 表示剩下 piles[i:] 时，M=j 时，先手可以获得最多石头数
// dp[i][j]=sum(piles[i:]) for i+2*j>=len(piles)
// dp[i][j]=max{sum(piles[i:])-dp[i+k][max(j,k)] for 1<=k<=2j}
// sum(piles[i:]) 表示剩下总的石头数
// dp[i+k][max(j,k)] 假如当前取 k 堆后，另外一个人可以拿到的石头数
// @lc code=start
func stoneGameII(piles []int) int {
	// dp 数组
	dp := make([][]int, len(piles))
	for i := range dp {
		dp[i] = make([]int, len(piles)+1)
	}
	// 先求出所有的 sum(piles[i:])
	sum := make([]int, len(piles))
	sum[len(piles)-1] = piles[len(piles)-1]
	for i := len(piles) - 2; i >= 0; i-- {
		sum[i] = sum[i+1] + piles[i]
	}
	var dpFunc func(i, j int) int
	dpFunc = func(i, j int) int {
		if dp[i][j] > 0 {
			return dp[i][j]
		}
		// 取 2*j 超过了剩下的了，则全取
		if i+2*j >= len(piles) {
			dp[i][j] = sum[i]
			return dp[i][j]
		}
		// 最大值绝对不会超过全取
		result := sum[i]
		for k := 1; i+k < len(piles) && k <= 2*j; k++ {
			result = min(result, dpFunc(i+k, max(j, k)))
		}
		dp[i][j] = sum[i] - result
		return dp[i][j]
	}
	return dpFunc(0, 1)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

