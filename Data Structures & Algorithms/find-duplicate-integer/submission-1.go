func findDuplicate(nums []int) int {
    slow, fast := nums[0], nums[0]
    for true {
        slow = nums[slow]
        fast = nums[nums[fast]]
        if slow == fast {
            break
        }

    }
    p := nums[0]
    for p != slow {
        p = nums[p]
        slow = nums[slow]
    }
    return p
}
