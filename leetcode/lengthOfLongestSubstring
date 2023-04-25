func lengthOfLongestSubstring(s string) int {
	left := 0
	right := 0
	curLongest := 0
	Hashmap := make(map[string]int)
	for right < len(s) {
		if len(s)-left <= curLongest {
			break
		}
		char := string(s[right])
		Hashmap[char] = Hashmap[char] + 1

		for Hashmap[char] != 1 {
			charl := string(s[left])
			Hashmap[charl] = Hashmap[charl] - 1
			left++
		}

		if curLongest <= right-left+1 {
			curLongest = right - left + 1
		}

		right++
	}

	return curLongest
}
