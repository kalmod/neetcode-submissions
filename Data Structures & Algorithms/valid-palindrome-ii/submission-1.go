// what is a palindrome?
// the same string forwards and backwards r a c e c a r
// other characterstrics of a palindrome? even number of characters with at most one odd number.  
// this is two pointer, so that's less important
// ... function that checks for palindrome, when we encounter a possible error, we try either incrementing left or right
func validPalindrome(s string) bool {

	chars := []rune(s)

	var isPalindrome func(chars []rune, l, r, k int) bool
	isPalindrome = func(chars []rune, l, r, k int) bool {
		// met or passed
		if l >= r {
			return true
		} else if chars[l] == chars[r] {
			return isPalindrome(chars, l+1, r-1, k)
		} else if k == 0 {
			// k is to track if we already took at most one char. If we encounter another miss match, we propogate back up
			return isPalindrome(chars, l+1, r, k+1) || isPalindrome(chars, l, r-1, k+1)
		}

		return false

	}

	return isPalindrome(chars, 0, len(chars)-1, 0)
}
