/*
 * @lc app=leetcode.cn id=552 lang=golang
 *
 * [552] 学生出勤记录 II
 *
 * https://leetcode-cn.com/problems/student-attendance-record-ii/description/
 *
 * algorithms
 * Hard (43.74%)
 * Likes:    220
 * Dislikes: 0
 * Total Accepted:    20.4K
 * Total Submissions: 35.7K
 * Testcase Example:  '2'
 *
 * 可以用字符串表示一个学生的出勤记录，其中的每个字符用来标记当天的出勤情况（缺勤、迟到、到场）。记录中只含下面三种字符：
 *
 * 'A'：Absent，缺勤
 * 'L'：Late，迟到
 * 'P'：Present，到场
 *
 *
 * 如果学生能够 同时 满足下面两个条件，则可以获得出勤奖励：
 *
 *
 * 按 总出勤 计，学生缺勤（'A'）严格 少于两天。
 * 学生 不会 存在 连续 3 天或 连续 3 天以上的迟到（'L'）记录。
 *
 *
 * 给你一个整数 n ，表示出勤记录的长度（次数）。请你返回记录长度为 n 时，可能获得出勤奖励的记录情况 数量 。答案可能很大，所以返回对 10^9 +
 * 7 取余 的结果。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 2
 * 输出：8
 * 解释：
 * 有 8 种长度为 2 的记录将被视为可奖励：
 * "PP" , "AP", "PA", "LP", "PL", "AL", "LA", "LL"
 * 只有"AA"不会被视为可奖励，因为缺勤次数为 2 次（需要少于 2 次）。
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：3
 *
 *
 * 示例 3：
 *
 *
 * 输入：n = 10101
 * 输出：183236316
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

// @lc code=start
// dp(i,a,l) 表示出勤记录次数 i 的情况下，缺勤 a 次，最后两天连续迟到 l 次时，可以获得出勤奖励的个数
// a >= 2 dp(i,a,l)=0，所以 a 只看 0，1。l 也只能 0<=l<=3
// dp(0,0,0)=1
// if 'A' then
//   dp(i,1,0) += dp(i-1,0,0)+dp(i-1,0,1)+dp(i-1,0,2)
// if 'L' then
//   dp(i,a,1) += dp(i-1,a,0)
//   dp(i,a,2) += dp(i-1,a,1)
// if 'P' then
//   dp(i,a,0) += dp(i-1,a,0)+dp(i-1,a,1)+dp(i-1,a,2)
// 求 dp(n,a,l) for 0<=a<=1 and 0<=l<=2

var dp1 [2][3]int
var dp2 [2][3]int

func dpClear(dp *[2][3]int) {
	dp[0][0] = 0
	dp[0][1] = 0
	dp[0][2] = 0
	dp[1][0] = 0
	dp[1][1] = 0
	dp[1][2] = 0
}

const mod = 1e9 + 7

func checkRecord(n int) int {
	dpClear(&dp1)
	dpClear(&dp2)
	dpPre := dp1
	dp := dp2
	dpPre[0][0] = 1
	for i := 1; i <= n; i++ {
		// A
		for l := 0; l <= 2; l++ {
			dp[1][0] = (dp[1][0] + dpPre[0][l]) % mod
		}
		// L
		for a := 0; a <= 1; a++ {
			for l := 1; l <= 2; l++ {
				dp[a][l] = (dp[a][l] + dpPre[a][l-1]) % mod
			}
		}
		// P
		for a := 0; a <= 1; a++ {
			for l := 0; l <= 2; l++ {
				dp[a][0] = (dp[a][0] + dpPre[a][l]) % mod
			}
		}
		dp, dpPre = dpPre, dp
		dpClear(&dp)
	}
	var sum int
	for a := 0; a <= 1; a++ {
		for l := 0; l <= 2; l++ {
			sum = (sum + dpPre[a][l]) % mod
		}
	}
	return sum
}

// @lc code=end

