/*
 * @lc app=leetcode.cn id=576 lang=golang
 *
 * [576] 出界的路径数
 *
 * https://leetcode-cn.com/problems/out-of-boundary-paths/description/
 *
 * algorithms
 * Medium (39.58%)
 * Likes:    138
 * Dislikes: 0
 * Total Accepted:    9.8K
 * Total Submissions: 24.3K
 * Testcase Example:  '2\n2\n2\n0\n0'
 *
 * 给你一个大小为 m x n 的网格和一个球。球的起始坐标为 [startRow, startColumn]
 * 。你可以将球移到在四个方向上相邻的单元格内（可以穿过网格边界到达网格之外）。你 最多 可以移动 maxMove 次球。
 *
 * 给你五个整数 m、n、maxMove、startRow 以及 startColumn ，找出并返回可以将球移出边界的路径数量。因为答案可能非常大，返回对
 * 10^9 + 7 取余 后的结果。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：m = 2, n = 2, maxMove = 2, startRow = 0, startColumn = 0
 * 输出：6
 *
 *
 * 示例 2：
 *
 *
 * 输入：m = 1, n = 3, maxMove = 3, startRow = 0, startColumn = 1
 * 输出：12
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= m, n <= 50
 * 0 <= maxMove <= 50
 * 0 <= startRow < m
 * 0 <= startColumn < n
 *
 *
 */

// dfs(i,j,k) 表示起始坐标是 [i,j]，至少移动 k 次可以移出边界的路径数量。
// 初始状态：
// dfs(-1,j,k) = 1，dfs(i,-1,k) = 1
// dfs(i,j,k) = 0 if
//   i-k>0 && i+k<m && j-k>0 && j+k<n
// 转移方程：
// dfs(i,j,k) = dfs(i,j-1,k-1)+dfs(i,j+1,k-1)+dfs(i-1,j,k-1)+dfs(i+1,j,k-1)
// 最终状态：
// dfs(startRow, startColumn, maxMove)
// @lc code=start
// mod 结果需要取模
const mod int64 = 1e9 + 7

// store 用一个全局内存，通过 initStore，减少内存分配
var store [50][50][51]int

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	// 初始化
	initStore(m, n, maxMove)
	var dfs func(i, j, k int) int
	dfs = func(i, j, k int) int {
		// 剪枝，一路直走都不能出界，直接返回 0
		if i-k > 0 && i+k < m && j-k > 0 && j+k < n {
			return 0
		}
		// 出界了
		if i == -1 || i == m || j == -1 || j == n {
			return 1
		}
		// 没有剩余次数了
		if k == 0 {
			return 0
		}
		// 查表，有结果则直接返回
		if res := store[i][j][k]; res > 0 {
			return res
		}
		// 用 int64 做中间变量，只需要取一次模
		// 其实我发现 int 应该就是 int64，只需要最后结果时取模就可以了
		res := int64(dfs(i, j-1, k-1)) +
			int64(dfs(i, j+1, k-1)) +
			int64(dfs(i-1, j, k-1)) +
			int64(dfs(i+1, j, k-1))
		result := int(res % mod)
		store[i][j][k] = result
		return result
	}
	return dfs(startRow, startColumn, maxMove)
}

func initStore(m, n, maxMove int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k <= maxMove; k++ {
				store[i][j][k] = -1
			}
		}
	}
}

// @lc code=end

