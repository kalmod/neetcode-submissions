class Solution:
    def sortArray(self, nums: List[int]) -> List[int]:
        # How can I determine, what is small as I iterate through the list.
        # minheap?
        heapq.heapify(nums)
        ans = []
        while nums:
            ans.append(heapq.heappop(nums))

        return ans