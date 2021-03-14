/*
 * @lc app=leetcode.cn id=47 lang=cpp
 *
 * [47] 全排列 II
 *
 * https://leetcode-cn.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (59.59%)
 * Likes:    381
 * Dislikes: 0
 * Total Accepted:    82K
 * Total Submissions: 137.4K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个可包含重复数字的序列，返回所有不重复的全排列。
 * 
 * 示例:
 * 
 * 输入: [1,1,2]
 * 输出:
 * [
 * ⁠ [1,1,2],
 * ⁠ [1,2,1],
 * ⁠ [2,1,1]
 * ]
 * 
 */

// @lc code=start
class Solution {
public:
    vector<vector<int>> permuteUnique(vector<int>& nums) {
        vector<vector<int>> result;
        function<void(int)> dfs;
        dfs = [&](int start) {
            if (start == nums.size() - 1) {
                result.push_back(nums);
            }
            set<int> unique;
            for (int i = start; i < nums.size(); ++i) {
                if (unique.find(nums[i]) != unique.end()) {
                    continue;
                }
                unique.insert(nums[i]);
                swap(nums[start], nums[i]);
                dfs(start + 1);
                swap(nums[start], nums[i]);
            }
        };
        dfs(0);
        return result;
    }
};
// @lc code=end

