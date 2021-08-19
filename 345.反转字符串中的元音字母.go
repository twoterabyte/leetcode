/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 *
 * https://leetcode-cn.com/problems/reverse-vowels-of-a-string/description/
 *
 * algorithms
 * Easy (51.79%)
 * Likes:    210
 * Dislikes: 0
 * Total Accepted:    102.9K
 * Total Submissions: 191.4K
 * Testcase Example:  '"hello"'
 *
 * 给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
 *
 * 元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "hello"
 * 输出："holle"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "leetcode"
 * 输出："leotcede"
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 3 * 10^5
 * s 由 可打印的 ASCII 字符组成
 *
 *
 */

// @lc code=start
var vowels = map[byte]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'A': true,
	'E': true,
	'I': true,
	'O': true,
	'U': true,
}

func reverseVowels(s string) string {
	bytes := []byte(s)
	i := next(bytes, 0)
	j := prev(bytes, len(bytes)-1)
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i = next(bytes, i+1)
		j = prev(bytes, j-1)
	}
	return string(bytes)
}

func next(bytes []byte, i int) int {
	for ; i < len(bytes); i++ {
		if vowels[bytes[i]] {
			return i
		}
	}
	return i
}

func prev(bytes []byte, i int) int {
	for ; i >= 0; i-- {
		if vowels[bytes[i]] {
			return i
		}
	}
	return i
}

// @lc code=end

