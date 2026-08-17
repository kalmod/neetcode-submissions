class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        # frequency table?
        freq = {}
        for i in range(len(s1)):
            if s1[i] not in freq:
                freq[s1[i]] = 0
            freq[s1[i]] += 1
        
        l, r = 0, 0
        while r < len(s2):
            if s2[r] not in freq:
                while l < r:
                    if s2[l] in freq: freq[s2[l]] += 1
                    l += 1
                l += 1
            else:
                while freq[s2[r]] == 0:
                    if s2[l] in freq:
                        freq[s2[l]] += 1
                    l += 1
                freq[s2[r]] -= 1
                if r - l + 1 == len(s1):
                    return True
            r += 1
        
        return False