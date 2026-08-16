impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
		let (mut l, mut r) = (0i32, nums.len() as i32 - 1);

		while l <= r {
			let m:i32 = l + (r-l)/2;	
		if nums[m as usize] == target {
				return m as i32;
			} else {
				// find which portion of array is sorted
				if nums[l as usize] <= nums[m as usize] { // left side is sorted
					if target >= nums[l as usize] && target < nums[m as usize] {
						r = m - 1;
					} else {
						l = m + 1;
					}
				} else {
					if target > nums[m as usize] && target <= nums[r as usize] {
						l = m + 1;
					} else  {
						r = m - 1;
					}
				}
			}
		}		

		return -1
    }
}
