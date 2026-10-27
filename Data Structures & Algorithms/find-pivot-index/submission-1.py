class Solution:
    def pivotIndex(self, nums: List[int]) -> int:
        leftSum = [0]
        rightSum = [0]
        for n in nums:
            leftSum.append(leftSum[-1] + n)
        for i in range(len(nums)-1, -1, -1):
            rightSum = [nums[i]+ rightSum[0]] + rightSum
        
        print(leftSum)
        print(rightSum)
        for i in range(len(nums)):
            if leftSum[i] == rightSum[i+1]:
                return i

        return -1