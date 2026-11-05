class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        h = {}
        for n in nums:
            h[n] = True
        
        longestAns = 0
        for n in nums:
            longestRunning = 0
            counter = n 
            if counter - 1 not in h:
                while counter in h:
                    longestRunning +=1
                    counter += 1
            longestAns = max(longestRunning, longestAns)
        
        return longestAns