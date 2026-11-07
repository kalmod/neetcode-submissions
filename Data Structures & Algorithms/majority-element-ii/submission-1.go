import (
	"slices"
)
func majorityElement(nums []int) []int {
	ans := []int{}
	i := 0
	limit := len(nums)/3
	slices.Sort(nums)
	for i < len(nums) {
		num := nums[i]
		count := 0
		for i < len(nums) && nums[i] == num {
			count+=1
			i+=1
		}
		if count > limit {
			ans = append(ans, num)
		}
	}
	return ans
}
