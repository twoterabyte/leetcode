/*
 * @lc app=leetcode.cn id=1137 lang=cpp
 *
 * [1137] 第 N 个泰波那契数
 *
 * https://leetcode-cn.com/problems/n-th-tribonacci-number/description/
 *
 * algorithms
 * Easy (59.34%)
 * Likes:    107
 * Dislikes: 0
 * Total Accepted:    54K
 * Total Submissions: 91K
 * Testcase Example:  '4'
 *
 * 泰波那契序列 Tn 定义如下： 
 * 
 * T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2
 * 
 * 给你整数 n，请返回第 n 个泰波那契数 Tn 的值。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：n = 4
 * 输出：4
 * 解释：
 * T_3 = 0 + 1 + 1 = 2
 * T_4 = 1 + 1 + 2 = 4
 * 
 * 
 * 示例 2：
 * 
 * 输入：n = 25
 * 输出：1389537
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 <= n <= 37
 * 答案保证是一个 32 位整数，即 answer <= 2^31 - 1。
 * 
 * 
 */
// Tn = Tn_1 + Tn_2 + Tn_3

// [Tn  ]   [1,1,1] [Tn_1]
// [Tn_1] = [1,0,0]*[Tn_2]
// [Tn-2]   [0,1,0] [Tn_3]

// [Tn  ]   ([1,1,1])         [T2]
// [Tn_1] = ([1,0,0])^(n-2) * [T1]
// [Tn-2]   ([0,1,0])         [T0]
// @lc code=start
vector<vector<unsigned int>> mat = {
    {1, 1, 1},
    {1, 0, 0},
    {0, 1, 0}
};

vector<vector<unsigned int>> t2 = {
    {1},
    {1},
    {0}
};

// 矩阵乘法
template<typename T>
vector<vector<T>> mul(const vector<vector<T>> &a, const vector<vector<T>> &b) {
    auto m = a.size();
    auto n = b[0].size();
    auto l = a[0].size();
    vector<vector<T>> res(m, vector<T>(n));
    for (auto i = 0; i < m; i++) {
        for (auto k = 0; k < l; k++) {
            for (auto j = 0; j < n; j++) {
                res[i][j] += a[i][k] * b[k][j];
            }
        }
    }
    return res;
}

// 矩阵快速幂
template<typename T>
vector<vector<T>> pow(const vector<vector<T>> &mat, int n) {
    vector<vector<T>> res(mat.size(), vector<T>(mat.size()));
    for (auto i = 0; i < mat.size(); i++) {
        res[i][i] = 1;
    }
    auto tmp = mat;
    while (n > 0) {
        if ((n & 1) == 1) {
            res = mul(res, tmp);
        }
        tmp = mul(tmp, tmp);
        n >>= 1;
    }
    return res;
}

class Solution {
public:
    int tribonacci(int n) {
        if (n == 0) {
            return 0;
        }
        if (n == 1) {
            return 1;
        }
        if (n == 2) {
            return 1;
        }
        auto matN_2 = pow(mat, n - 2);
        return mul(matN_2, t2)[0][0];
    }
};
// @lc code=end

