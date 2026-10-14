func isAnagram(s string, t string) bool {
    // I think if i put both into slice
    // if a character is odd, it shows in only one string
    if len(s) != len(t) {
        return false
    }

    charS := []rune(s)
    sort.Slice(charS, func(i, j int) bool {
        return charS[i] < charS[j]
    })
    charT := []rune(t)
    sort.Slice(charT, func(i, j int) bool {
        return charT[i] < charT[j]
    })

    for i := range len(charS) {
        if charS[i] != charT[i] {
            return false
        }
    }

    return true
}

