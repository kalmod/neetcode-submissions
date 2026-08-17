type TimeMap struct {
    hash map[string][]Pair
}
type Pair struct {
    value string
    timestamp int
}

func Constructor() TimeMap {
    return TimeMap{make(map[string][]Pair)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
    this.hash[key] = append(this.hash[key], Pair{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
    keyArray, exists := this.hash[key]
    if !exists {
        return ""
    }

    res := ""
    l := 0
    r := len(keyArray) - 1
    for l <= r {
        m := l + (r - l) / 2
        if keyArray[m].timestamp <= timestamp {
            res = keyArray[m].value
            l = m + 1
        } else {
            r = m - 1
        }
    }
    return res

}
