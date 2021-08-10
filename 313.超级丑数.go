/*
 * @lc app=leetcode.cn id=313 lang=golang
 *
 * [313] 超级丑数
 */

// @lc code=start

// 假如存在存在超级丑数 x，则 x*primes[i] 也是超级丑数
// dp(i) 表示第 i 个超级丑数，则有 dp(0) = 1
// dp(n) = min{dp(i)*primes[j] for 0<=i<n && 0<=j<len(primes) && dp(i)*primes[j] > dp(n-1)}
// dp(i) 和 rimes[j] 是递增的，在上面的遍历过程中，如果 dp(i)*primes[j] 已经是 dp 中的超级丑数了
// 那么对于 k<=i，h<=j，有 dp(k)*primes[h] <= dp(i)*primes[j]，则 dp(k)*primes[h] 也一定在 dp 中
// index[j] 表示 primes[j]*dp(index[j]) 不在 dp 中
// dp(0)=1, index[:]=0
// dp(n)=min{primes[i]*dp(index[i]) for 0<=i<len(primes)}
// primes[i] =
// if primes[i]*dp(index[i]) == dp(n) then primes[i] + 1
// else primes[i]

func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n)
	dp[0] = 1
	index := make([]int, len(primes))
	for i := 1; i < n; i++ {
		dp[i] = primes[0] * dp[index[0]]
		for j := 1; j < len(primes); j++ {
			dp[i] = min(dp[i], primes[j]*dp[index[j]])
		}
		for j := 0; j < len(primes); j++ {
			if dp[i] == primes[j]*dp[index[j]] {
				index[j]++
			}
		}
	}
	return dp[n-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// @lc code=end

