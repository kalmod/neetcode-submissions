class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        start = len(nums1) - len(nums2) 

        for i in range(len(nums2)):
            ptr = len(nums1) - len(nums2) + i
            nums1[ptr] = nums2[i]
            while ptr >= 1 and nums1[ptr] < nums1[ptr-1]:
                nums1[ptr], nums1[ptr-1] = nums1[ptr-1], nums1[ptr]
                ptr-=1