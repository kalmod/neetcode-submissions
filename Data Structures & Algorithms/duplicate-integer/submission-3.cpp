class Solution {
public:
    bool hasDuplicate(vector<int>& nums) {
        std::unordered_set<int> numSet;
        for (int n : nums) {
            if (numSet.contains(n)) {
                return true;
            }
            numSet.insert(n);
        }
        return false;
    }
};