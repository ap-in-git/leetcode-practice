package main

func twoSum(numbers []int, target int) []int {

	l, r := 0, len(numbers)-1
	for l <= r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			break
		}
		if sum > target {
			r--
		} else {
			l++
		}
	}
	return []int{l, r}
}
