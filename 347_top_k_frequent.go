package main

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)

	for _, num := range nums {
		counts[num]++
	}

	countNums := make(map[int][]int, 0)
	for val, count := range counts {
		_, ok := countNums[count]
		if ok {
			countNums[count] = append(countNums[count], val)
		} else {
			countNums[count] = []int{val}
		}
	}

	var returns []int
	for i := len(nums); i > 0; i-- {
		at := countNums[i]
		if at != nil {
			for j := len(at); j > 0; j-- {
				returns = append(returns, at[j-1])
				if len(returns) == k {
					break
				}
			}
		}
		if len(returns) == k {
			break
		}

	}
	return returns
}
