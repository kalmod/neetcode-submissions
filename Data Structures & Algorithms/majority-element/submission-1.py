class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        nums.sort()
        ans = 0
        i = 0
        while i < len(nums):
            cur = nums[i]
            startPos = i
            while i < len(nums) and nums[i] == cur:
                i+=1
            if i - startPos > len(nums)//2:
                ans = nums[i-1] 
            
        return ans