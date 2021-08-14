/*
 * @lc app=leetcode.cn id=233 lang=golang
 *
 * [233] 数字 1 的个数
 *
 * https://leetcode-cn.com/problems/number-of-digit-one/description/
 *
 * algorithms
 * Hard (41.06%)
 * Likes:    314
 * Dislikes: 0
 * Total Accepted:    31.1K
 * Total Submissions: 65.8K
 * Testcase Example:  '13'
 *
 * 给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 13
 * 输出：6
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 0
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 *
 *
 */

// 所有的非负整数，位数 10^i 上，1 出现的次数记为 cnt(i)
// cnt(i) =
// if n%(10^(i+1)) < 10^i then n/(10^(i+1))*10^i
// if 10^i <= n%(10^(i+1)) < 2*10^i then n/(10^(i+1))*10^i + n%(10^(i+1))-10^i+1
// if n%(10^(i+1)) >= 2*10^i then n/(10^(i+1))*10^i + 10^i
// 则是求，sum{cnt(i)}

// 当位数是 10^i 时，最后的 i 位每 n/(10^(i+1)) 个数就会循环一遍，且包含 10^i 个 1
// 还剩下不在循环中的部分是 n%(10^(i+1)) 个数，这部分数
// 如果 n%(10^(i+1)) < 10^i，则一次也不会出现 1
// 如果 10^i <= n%(10^(i+1)) < 2*10^i，则会出现 n%(10^(i+1))-10^i+1 次 1
// 如果 n%(10^(i+1)) >= 2*10^i，则会出现 10^i 次

// 以 1234567 的 10^2 位为例子，每 1234 次就会剩下的位数就会出现[100,199]，即 10^2 次1
// 超过 1234000 的 567 部分，因为大于 10^2，则会出现 10^2 次，即 [100,199]
// 如果是 1234012，则因为 012 小于 10^2，则一次也不会出现
// 如果是 1234123，则因为 123 大于等于 100 小于 200，则会出现 123-10^2+1 次 1，即 [100,123]

// @lc code=start
func countDigitOne(n int) int {
	var result int
	// pow 就是 10^i
	for pow := 1; pow <= n; {
		// 10^(i+1)
		pow10 := pow * 10
		// n/(10^(i+1))
		quotient := n / pow10
		// n%(10^(i+1))
		remainder := n % pow10
		if remainder < pow {
			result += quotient * pow
		} else if remainder < 2*pow {
			result += quotient*pow + remainder - pow + 1
		} else {
			result += quotient*pow + pow
		}
		pow = pow10
	}
	return result
}

// @lc code=end

