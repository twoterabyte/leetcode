/*
 * @lc app=leetcode.cn id=551 lang=golang
 *
 * [551] 学生出勤记录 I
 *
 * https://leetcode-cn.com/problems/student-attendance-record-i/description/
 *
 * algorithms
 * Easy (52.98%)
 * Likes:    110
 * Dislikes: 0
 * Total Accepted:    55.2K
 * Total Submissions: 97.5K
 * Testcase Example:  '"PPALLP"'
 *
 * 给你一个字符串 s 表示一个学生的出勤记录，其中的每个字符用来标记当天的出勤情况（缺勤、迟到、到场）。记录中只含下面三种字符：
 *
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
 * 学生 不会 存在 连续 3 天或 3 天以上的迟到（'L'）记录。
 *
 *
 * 如果学生可以获得出勤奖励，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "PPALLP"
 * 输出：true
 * 解释：学生缺勤次数少于 2 次，且不存在 3 天或以上的连续迟到记录。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "PPALLL"
 * 输出：false
 * 解释：学生最后三天连续迟到，所以不满足出勤奖励的条件。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s[i] 为 'A'、'L' 或 'P'
 *
 *
 */

// @lc code=start
func checkRecord(s string) bool {
	// 缺勤天数 连续迟到天数
	var cntA, consecutiveL int
	for _, c := range s {
		if c == 'A' {
			cntA++
			// 缺勤大于两天就不满足了
			if cntA >= 2 {
				return false
			}
			// 重置连续迟到天数
			consecutiveL = 0
		} else if c == 'L' {
			consecutiveL++
			// 连续迟到 3 就不满足了
			if consecutiveL >= 3 {
				return false
			}
		} else {
			// 重置连续迟到天数
			consecutiveL = 0
		}
	}
	return true
}

// @lc code=end

