class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        # same characters, but different order...
        h = {}
        ans = []
        for s in strs:
            sorted_s = ''.join(sorted(s))
            if sorted_s in h:
                ans[h[sorted_s]].append(s)
            else:
                ans.append([s])
                h[sorted_s] = len(ans)-1
            
        return ans
