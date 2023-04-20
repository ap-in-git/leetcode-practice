package main

func twoSum(nums []int, target int) []int {
	nmap := make(map[int]int, len(nums))
	for index, el := range nums {
		remainingSum := target - el
		val, ok := nmap[remainingSum]
		if ok {
			return []int{val, index}
		}
		nmap[el] = index
	}
	return []int{}
}
