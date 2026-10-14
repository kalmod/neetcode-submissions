func twoSum(nums []int, target int) []int {
    // brute force is done with two loops
	// a + b = target. 
	// As I Go through the list in one pass I need a way to reference what's already been done
	//	A hash works for that
	// What do I store? Well, I have a, what is a equal to? a = target - b.
	// Well the hash is empty, and I want my b to exist
	// So i can store hash[target-a] = index_a (need to return the index of a)

	m := make(map[int]int)
	for idx, n := range nums {
		if jdx, exists := m[n]; exists {
			return []int{jdx, idx}
		}
		m[target-n] = idx
	}
	return []int{}
}
