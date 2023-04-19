package main

func longestConsecutive(nums []int) int {
	se := make(map[int]bool)
	for _, el := range nums {
		se[el] = true
	}

	lcs := 0
	for _, el := range nums {
		currentLcs := 1
		current := el
		if !se[current-1] {
			for se[current+1] {
				currentLcs++
				current++
			}
		}
		if currentLcs >= lcs {
			lcs = currentLcs
		}
	}
	return lcs
}
