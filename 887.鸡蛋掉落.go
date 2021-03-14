/*
 * @lc app=leetcode.cn id=887 lang=golang
 *
 * [887] 鸡蛋掉落
 *
 * https://leetcode-cn.com/problems/super-egg-drop/description/
 *
 * algorithms
 * Hard (28.69%)
 * Likes:    447
 * Dislikes: 0
 * Total Accepted:    27.9K
 * Total Submissions: 97K
 * Testcase Example:  '1\n2'
 *
 * 你将获得 K 个鸡蛋，并可以使用一栋从 1 到 N  共有 N 层楼的建筑。
 * 
 * 每个蛋的功能都是一样的，如果一个蛋碎了，你就不能再把它掉下去。
 * 
 * 你知道存在楼层 F ，满足 0 <= F <= N 任何从高于 F 的楼层落下的鸡蛋都会碎，从 F 楼层或比它低的楼层落下的鸡蛋都不会破。
 * 
 * 每次移动，你可以取一个鸡蛋（如果你有完整的鸡蛋）并把它从任一楼层 X 扔下（满足 1 <= X <= N）。
 * 
 * 你的目标是确切地知道 F 的值是多少。
 * 
 * 无论 F 的初始值如何，你确定 F 的值的最小移动次数是多少？
 * 
 * 
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：K = 1, N = 2
 * 输出：2
 * 解释：
 * 鸡蛋从 1 楼掉落。如果它碎了，我们肯定知道 F = 0 。
 * 否则，鸡蛋从 2 楼掉落。如果它碎了，我们肯定知道 F = 1 。
 * 如果它没碎，那么我们肯定知道 F = 2 。
 * 因此，在最坏的情况下我们需要移动 2 次以确定 F 是多少。
 * 
 * 
 * 示例 2：
 * 
 * 输入：K = 2, N = 6
 * 输出：3
 * 
 * 
 * 示例 3：
 * 
 * 输入：K = 3, N = 14
 * 输出：4
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= K <= 100
 * 1 <= N <= 10000
 * 
 * 
 */

// dp[i][j] 表示 i 个鸡蛋，j 层楼，确定 F 的最小移动次数
// dp[1][j]=j，dp[i][0]=0
// dp[i][j]=min{max{dp[i-1][k-1], dp[i][j-k]}+1 for 1<=k<=j}
// 从 j 层中选择 k 层，如果在 k 层破了 则看 dp[i-1][k-1]，如果没破则看 dp[i][j-k]
// @lc code=start

import "math"

func superEggDrop(K int, N int) int {
	dp := make([][]int, K+1)
	for i := range dp {
		dp[i] = make([]int, N+1)
	}
	var eggDrop func(k, n int) int
	eggDrop = func(k, n int) int {
		if n == 0 {
			return 0
		}
		if k == 1 {
			return n
		}
		if dp[k][n] != 0 {
			return dp[k][n]
		}
		// 最大不会超过 n 所以这里默认取 n
		// res := n
		// for i := 1; i <= n; i++ {
		// 	res = min(res, max(eggDrop(k-1, i-1), eggDrop(k, n-i))+1)
		// }
		// dp[k][n] = res
		// 由于固定 k，dp[k][n] 是单调递增的，max(eggDrop(k-1, i-1), eggDrop(k, n-i) 则是先递减再递增的函数
		// 所以这里可以使用二分查找法
		res := n
		l, r := 1, n
		for l <= r {
			mid := l + (r - l) / 2
			broken := eggDrop(k-1, mid-1)
			not_broken := eggDrop(k, n-mid)
			if broken > not_broken {
				r = mid - 1
				res = min(res, broken+1)
			} else {
				l = mid + 1
				res = min(res, not_broken+1)
			}
		}
		dp[k][n] = res
		return res
	}
	return eggDrop(K, N)
}

// 超时了，改成自顶向下，试试
func superEggDrop1(K int, N int) int {
	dp := make([]int, N+1)
	dpPre := make([]int, N+1)
	// 初始化 dp[1][j]=j，dp[i][0]=0
	for j := range dp {
		dp[j] = j
	}
	for i := 2; i <= K; i++ {
		// dp, dpPre := dpPre, dp 和下面的语句区别非常大
		dp, dpPre = dpPre, dp
		for j := 1; j <= N; j++ {
			dp[j] = math.MaxInt32
			for k := 1; k <= j; k++ {
				dp[j] = min(max(dpPre[k-1], dp[j-k])+1, dp[j])
			}
		}
	}
	return dp[N]
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

