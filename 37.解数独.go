/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 *
 * https://leetcode-cn.com/problems/sudoku-solver/description/
 *
 * algorithms
 * Hard (62.53%)
 * Likes:    510
 * Dislikes: 0
 * Total Accepted:    38.8K
 * Total Submissions: 62.1K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * 编写一个程序，通过已填充的空格来解决数独问题。
 * 
 * 一个数独的解法需遵循如下规则：
 * 
 * 
 * 数字 1-9 在每一行只能出现一次。
 * 数字 1-9 在每一列只能出现一次。
 * 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
 * 
 * 
 * 空白格用 '.' 表示。
 * 
 * 
 * 
 * 一个数独。
 * 
 * 
 * 
 * 答案被标成红色。
 * 
 * Note:
 * 
 * 
 * 给定的数独序列只包含数字 1-9 和字符 '.' 。
 * 你可以假设给定的数独只有唯一解。
 * 给定数独永远是 9x9 形式的。
 * 
 * 
 */

// @lc code=start
func solveSudoku(board [][]byte)  {
	// 判断 board[i][j] 是否可以放入数字 c
	isValid := func(i, j int, c byte) bool {
		for k := 0; k < 9; k++ {
			if board[i][k] == c {
				// 同行是否存在
				return false
			}
			if board[k][j] == c {
				// 同列是否存在
				return false
			}
			// 9 宫格里面是否已经存在了
			if board[i/3*3+k/3][j/3*3+k%3] == c {
				return false
			}
		}
		return true
	}
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if j == 9 {
			// 列超过9，移动到下一行
			return dfs(i+1, 0)
		}
		if i == 9 {
			// 到了第9行，说明已经全正确了
			return true
		}
		if board[i][j] != '.' {
			// 已经有数字，直接下一个
			return dfs(i, j+1)
		}
		var c byte
		for c = '1'; c <= '9'; c++ {
			if !isValid(i, j, c) {
				// 无效跳过这个数字
				continue
			}
			board[i][j] = c
			if (dfs(i, j+1)) {
				// 递归有效，直接返回
				return true
			}
			
		}
		// 没有一个有效的数字，回退当前写入的数字
		board[i][j] = '.'
		return false
	}
	dfs(0, 0)
}
// @lc code=end

