/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode-cn.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (76.08%)
 * Likes:    1283
 * Dislikes: 0
 * Total Accepted:    174K
 * Total Submissions: 228.5K
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入：n = 3
 * 输出：[
 * ⁠      "((()))",
 * ⁠      "(()())",
 * ⁠      "(())()",
 * ⁠      "()(())",
 * ⁠      "()()()"
 * ⁠    ]
 * 
 * 
 */

// @lc code=start
func generateParenthesis(n int) []string {
	var result []string
	track := make([]byte, 2*n)
	var dfs func(i, right, left int)
	dfs = func(i, right, left int) {
		if i == 2*n {
			tmp := make([]byte, 2*n)
			copy(tmp, track)
			result = append(result, string(tmp))
		}
		if right < n {
			track[i] = '('
			dfs(i+1, right+1, left)
		}
		if left < right {
			track[i] = ')'
			dfs(i+1, right, left+1)
		}
	}
	dfs(0, 0, 0)
	return result
}
// @lc code=end

