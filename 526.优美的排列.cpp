/*
 * @lc app=leetcode.cn id=526 lang=cpp
 *
 * [526] 优美的排列
 *
 * https://leetcode-cn.com/problems/beautiful-arrangement/description/
 *
 * algorithms
 * Medium (65.64%)
 * Likes:    147
 * Dislikes: 0
 * Total Accepted:    12.7K
 * Total Submissions: 19.2K
 * Testcase Example:  '2'
 *
 * 假设有从 1 到 N 的 N 个整数，如果从这 N 个数字中成功构造出一个数组，使得数组的第 i 位 (1 <= i <= N)
 * 满足如下两个条件中的一个，我们就称这个数组为一个优美的排列。条件：
 * 
 * 
 * 第 i 位的数字能被 i 整除
 * i 能被第 i 位上的数字整除
 * 
 * 
 * 现在给定一个整数 N，请问可以构造多少个优美的排列？
 * 
 * 示例1:
 * 
 * 
 * 输入: 2
 * 输出: 2
 * 解释: 
 * 
 * 第 1 个优美的排列是 [1, 2]:
 * ⁠ 第 1 个位置（i=1）上的数字是1，1能被 i（i=1）整除
 * ⁠ 第 2 个位置（i=2）上的数字是2，2能被 i（i=2）整除
 * 
 * 第 2 个优美的排列是 [2, 1]:
 * ⁠ 第 1 个位置（i=1）上的数字是2，2能被 i（i=1）整除
 * ⁠ 第 2 个位置（i=2）上的数字是1，i（i=2）能被 1 整除
 * 
 * 
 * 说明:
 * 
 * 
 * N 是一个正整数，并且不会超过15。
 * 
 * 
 */

// @lc code=start

// status 第 i 位上为 1，表示第 1 个数被选择了
// 则有 (1<<(n-1))&status != 0，表示 n 被选择了
// ~(1<<(n-1))&status，表示在 status 的情况下不选择 n
// dp(i,status) 表示只选择 i 个的情况下，选择状态为 status 下，满足优美排列的个数
// dp(0, 0)=1
// dp(i,status) = sum{dp(i-1,~(1<<(j-1))&status) for 1<=j<=15 and (1<<(j-1))&status != 0 and (i%j==0 or j%i==0)}
// dp(i-1,~(1<<(j-1))&status) 表示不选择 j 时的个数
// 1<=j<=15 and (1<<(j-1))&status != 0 表示遍历 status 中为 1 的部分，即被选择了的才需要
// (i%j==0 or j%i==0) 表示不选择 j，也表明需要加上 j 才能满足，所以 j 需要满足有效判断

// 上面可以看到实际 dp(i, status) 只依赖于 dp(i-1, status')，所以可以做状态压缩
// status' 是 status 中清除掉每个 1 后的数字，所以 status' < status
// 那么可以对 status 从小到大遍历即可。
// dp(status) 表示在选择状态为 status 下，满足优美排列的个数
// 初始状态 dp(0)=1，求 dp(2^n-1)
// count1(status) 表示 status 中有 1 的个数
// dp(status) = sum{dp(~(1<<(i-1))&status) for 1<=i<=15 and (1<<(i-1))&status != 0 and (count1(status)%i==0 or i%count1(status)==0)}
class Solution {
public:
    int countArrangement(int n) {
        // dp 数组
        vector<int> dp(1<<n);
        // 初始状态
        dp[0] = 1;
        // 从小到大遍历
        for (int status = 1; status < (1<<n); status++) {
            // status 中 1 的个数
            auto count1 = __builtin_popcount(status);
            for (int i = 1; i <= 15; i++) {
                // 遍历每位 1
                int pre_status = (1<<(i-1))&status;
                if (pre_status == 0) continue;
                // 当前 1 无效
                if (count1%i != 0 and i%count1 != 0) continue;
                // 状态转移
                dp[status] += dp[(~(1<<(i-1)))&status];
            }
        }
        return dp[(1<<n)-1];
    }
};
// @lc code=end

