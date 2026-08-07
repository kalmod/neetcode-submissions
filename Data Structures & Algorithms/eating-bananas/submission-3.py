class Solution:
	def minEatingSpeed(self, piles: List[int], h: int) -> int:
		low = 1
		high = max(piles)
		ans = high
		while low <= high:
			k = low + (high-low)//2
			hOFk = 0
			for p in piles:
				hOFk += math.ceil(p/k)	
			if hOFk > h: # too many hours to eat, need to be bigger
				low = k + 1
			else: # can eat within limit, move it down to find min
				high = k-1	
				ans = min(ans, k)
		return ans