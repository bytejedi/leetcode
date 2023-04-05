package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	var ret int
	l := len(s)
	m := make(map[uint8]struct{}, l)

	for i := range s {
		for j := i; j < l; j++ {
			_, ok := m[s[j]]
			if ok {
				break
			} else {
				m[s[j]] = struct{}{}
				if j-i+1 > ret {
					ret = j - i + 1
				}
			}
		}
		for k := range m {
			delete(m, k)
		}
	}

	return ret
}
