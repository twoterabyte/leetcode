/*
 * @lc app=leetcode.cn id=576 lang=cpp
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

// dp[i][j][k] 表示移动 k 次后到达点 [i,j] 的路径数，
// 为什么不和 go 语言一样的定义的呢，这中初始化状态一个，但是终止状态多个
// 初始化状态
// dp[startRow][startColumn][0]=1, dp[i][j][0]=0
// 状态转移方程
// dp[i][j][k] = dp[i-1][j][k-1]+dp[i+1][j][k-1]+dp[i][j-1][k-1]+dp[i][j+1][k-1]
// 上面的状态转移方程也可以写为
// dp[i-1][j][k+1] += dp[i][j][k]
// dp[i+1][j][k+1] += dp[i][j][k]
// dp[i][j-1][k+1] += dp[i][j][k]
// dp[i][j+1][k+1] += dp[i][j][k]
// 求 sum{dp[i][j][k] for i<0 or i>=m or j<0 or m>=n}
// k 可以进行状态压缩
// @lc code=start
int directions[4][2] = {
    {-1, 0},
    {1, 0},
    {0, -1},
    {0, 1}
};

constexpr int MOD = 1e9+7;

class Solution {
public:
    int findPaths(int m, int n, int maxMove, int startRow, int startColumn) {
        vector<vector<int>> dp(m, vector<int>(n));
        // 初始状态
        dp[startRow][startColumn] = 1;
        // 出界路径数
        int cnt = 0;
        // dp[i][j][k] 向 4 个方向算出 dp[i’][dp[j'][k+1]]
        for (int k = 0; k < maxMove; k++) {
            vector<vector<int>> dpNew(m, vector<int>(n));
            for (int i = 0; i < m; i++) {
                for (int j = 0; j < n; j++) {
                    // 如果当前路径数为 0，说明不能给四个方向贡献路径数，直接跳过
                    auto cntCur = dp[i][j];
                    if (cntCur == 0) {
                        continue;
                    }
                    // 处理 4 个方向
                    for (auto direction : directions) {
                        auto i1 = i + direction[0];
                        auto j1 = j + direction[1];
                        // 出界了，直接算到 cnt
                        if (i1 == -1 || i1 == m || j1 == -1 || j1 == n) {
                            cnt = (cnt + cntCur) % MOD;
                            continue;
                        }
                        // 没有出界，把当前的步数加到当前方向上
                        dpNew[i1][j1] = (dpNew[i1][j1] + cntCur) % MOD;
                    }
                }
            }
            dp = dpNew;
        }
        return cnt;
    }
};
// @lc code=end

