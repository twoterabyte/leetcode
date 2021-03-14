/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 *
 * https://leetcode-cn.com/problems/regular-expression-matching/description/
 *
 * algorithms
 * Hard (30.17%)
 * Likes:    1461
 * Dislikes: 0
 * Total Accepted:    108.3K
 * Total Submissions: 358.3K
 * Testcase Example:  '"aa"\n"a"'
 *
 * 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
 * 
 * '.' 匹配任意单个字符
 * '*' 匹配零个或多个前面的那一个元素
 * 
 * 
 * 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
 * 
 * 说明:
 * 
 * 
 * s 可能为空，且只包含从 a-z 的小写字母。
 * p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
 * 
 * 
 * 示例 1:
 * 
 * 输入:
 * s = "aa"
 * p = "a"
 * 输出: false
 * 解释: "a" 无法匹配 "aa" 整个字符串。
 * 
 * 
 * 示例 2:
 * 
 * 输入:
 * s = "aa"
 * p = "a*"
 * 输出: true
 * 解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
 * 
 * 
 * 示例 3:
 * 
 * 输入:
 * s = "ab"
 * p = ".*"
 * 输出: true
 * 解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
 * 
 * 
 * 示例 4:
 * 
 * 输入:
 * s = "aab"
 * p = "c*a*b"
 * 输出: true
 * 解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
 * 
 * 
 * 示例 5:
 * 
 * 输入:
 * s = "mississippi"
 * p = "mis*is*p*."
 * 输出: false
 * 
 */

// dp(i, j) 表示 s[i:] 和 p[j:] 是否匹配
// @lc code=start
func isMatch(s string, p string) bool {
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(p))
	}
	isCharMatch := func (i, j int) bool {
		// 因为允许 i == len(s) 还进行匹配，所以要对 i == len(s) 做特殊处理
		if i >= len(s) {
			return false
		}
		if p[j] == '.' || s[i] == p[j] {
			return true
		}
		return false
	}
	var dpFunc func(i, j int) bool
	dpFunc = func(i, j int) bool {
		if i == len(s) && j == len(p) {
			return true
		}
		if j == len(p) {
			return false
		}
		// 如果 i == len(s) 还可能 .* 的匹配在 p，但是是满足匹配的，所以不能返回
		// 先查dp
		if dp[i][j] != 0 {
			return dp[i][j] == 1
		}
		ret := false
		if j + 1 < len(p) && p[j+1] == '*' {
			ret = dpFunc(i, j+2) || isCharMatch(i, j) && dpFunc(i+1, j)
		} else {
			ret = isCharMatch(i, j) && dpFunc(i+1, j+1)
		}
		if ret {
			dp[i][j] = 1
		} else {
			dp[i][j] = -1
		}
		return ret
	}
	return dpFunc(0, 0)
}
// @lc code=end

