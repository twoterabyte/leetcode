/*
 * @lc app=leetcode.cn id=541 lang=golang
 *
 * [541] 反转字符串 II
 *
 * https://leetcode-cn.com/problems/reverse-string-ii/description/
 *
 * algorithms
 * Easy (58.04%)
 * Likes:    141
 * Dislikes: 0
 * Total Accepted:    44.7K
 * Total Submissions: 76.5K
 * Testcase Example:  '"abcdefg"\n2'
 *
 * 给定一个字符串 s 和一个整数 k，从字符串开头算起，每 2k 个字符反转前 k 个字符。
 *
 *
 * 如果剩余字符少于 k 个，则将剩余字符全部反转。
 * 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "abcdefg", k = 2
 * 输出："bacdfeg"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "abcd", k = 2
 * 输出："bacd"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^4
 * s 仅由小写英文组成
 * 1 <= k <= 10^4
 *
 *
 */

// @lc code=start
func reverseStr(s string, k int) string {
	k2 := 2 * k
	bytes := []byte(s)
	for i := 0; i < len(bytes); i += k2 {
		l := i + k - 1
		if l >= len(bytes) {
			l = len(bytes) - 1
		}
		r := i
		for r < l {
			bytes[r], bytes[l] = bytes[l], bytes[r]
			r++
			l--
		}
	}
	return string(bytes)
}

// @lc code=end

