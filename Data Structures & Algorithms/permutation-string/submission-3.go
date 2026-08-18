func checkInclusion(s1 string, s2 string) bool {
    // freq table to check if char in s2 exists in s1
    // when do we reset and adjust our window?
    // 1. When we encounter a character that does not exists in s1
    // 2. When we encounter an extra character that does not exists in s1
    // When do we stop? When our window has the same length as s1
    // if we reach the end of our loop, we return False
    freq := make(map[rune]int)
    for _, r := range s1 {
            freq[r]++
    }

    l := 0
    for idx, r := range s2 {
        if _, exists := freq[r]; !exists {
            for l < idx {
                if _, exists := freq[rune(s2[l])]; exists {
                    freq[rune(s2[l])]++
                    l++
                }
            }
            l++
        } else {
                for freq[r] == 0 {
                    if _, exists := freq[rune(s2[l])]; exists {
                        freq[rune(s2[l])]++
                        l++
                    }
                }
            
            freq[r]--
            if idx - l + 1 == len(s1) {
                return true
            }
        }
    }

    
    return false
}
