package main

import "sort"

//Sort
//Take first element
//Do two sum two on remaining elements
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ret [][]int
	for i := 0; i < len(nums); i++ {
		//Prevent duplicate. The current value is same as the value on left
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		//len(nums)-1 because we need three numbers
		l, r := i+1, len(nums)-1
		target := 0

		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
				continue
			}
			if sum > target {
				r--
			} else {
				l++
			}

		}

	}
	return ret
}
