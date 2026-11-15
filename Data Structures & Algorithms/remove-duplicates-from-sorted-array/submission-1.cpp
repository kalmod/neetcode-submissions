class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int wp = 0;
        for (int i = 1; i < nums.size(); i++) {
            if (nums[i] != nums[wp]) {
                wp++;
                nums[wp] = nums[i];
            }
        }
        return wp+1;
    }
};