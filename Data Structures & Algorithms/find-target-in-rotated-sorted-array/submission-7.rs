impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
		let mut l:usize = 0;
		let mut r:usize = nums.len()-1;

		while l <= r {
			let m:usize = l + (r-l)/2;	
		if nums[m] == target {
				return m as i32;
			} else {
				// find which portion of array is sorted
				if nums[l] <= nums[m] { // left side is sorted
					if target >= nums[l] && target < nums[m] {
						r = m - 1;
					} else {
						l = m + 1;
					}
				} else {
					if target > nums[m] && target <= nums[r] {
						l = m + 1;
					} else  {
						if m == 0 { break; }
						r = m - 1;
					}
				}
			}
		}		

		return -1
    }
}
