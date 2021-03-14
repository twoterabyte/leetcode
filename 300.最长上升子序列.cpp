/*
 * @lc app=leetcode.cn id=300 lang=cpp
 *
 * [300] 最长上升子序列
 */

// len[i] 表示长度为 i+1 的上升子序列的最后一个元素的最小值
// base: len[0] = arr[0]
// arr[n]:
// - if find lower_bound in len, lower_bound = arr[n]
// - else len push_back(arr[n])
// @lc code=start
#include <algorithm>
using namespace std;
class Solution {
public:
    int lengthOfLIS(vector<int>& nums) {
        if (nums.empty()) {
            return 0;
        }
        vector<int> len;
        // 初始状态
        len.push_back(nums[0]);
        // 记录目前最长
        for (int i = 1; i < nums.size(); ++i) {
            // 找第一个大于等于
            auto iter = lower_bound(len.begin(), len.end(), nums[i]);
            // 如果找到
            if (iter != len.end()) {
                *iter = nums[i];
                continue;
            }
            // 如果没找到，说明长度可以加 1
            len.push_back(nums[i]);
        }
        return len.size();
    }
};
// @lc code=end

