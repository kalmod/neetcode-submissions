class Solution:
    def majorityElement(self, nums: List[int]) -> List[int]:
        # first thing that comes to mind is a freq table. 
        # as I iterate through the list, add to my freq table. 
        # once I reach x > floor(n/3), then append ans
        ans = []
        freq = {}
        limit = len(nums)//3
        for n in nums:
            if n not in freq:
                freq[n] = 0
            freq[n]+=1
        for num, f in freq.items():
            if f > limit:
                ans.append(num)
        return ans 