class Solution:
    def sortArray(self, nums: List[int]) -> List[int]:
        # How can I determine, what is small as I iterate through the list.
        # minheap?
        def mergeSort(arr):
            if len(arr) == 1:
                return arr
            else:
                m = len(arr)//2
                l = arr[0:m]
                r = arr[m:len(arr)]
                res = mergeJoin(mergeSort(l), mergeSort(r))
            return res
        def mergeJoin(l_arr, r_arr):
            l, r, m, n = 0, 0, len(l_arr), len(r_arr)
            res = []
            while l < m and r < n:
                if l_arr[l] <= r_arr[r]:
                    res.append(l_arr[l])
                    l+=1
                else:
                    res.append(r_arr[r])
                    r+=1

            while l < m:
                res.append(l_arr[l])
                l+=1
            while r < n:
                res.append(r_arr[r])
                r+=1
            return res

        ans = mergeSort(nums)
                

        return ans