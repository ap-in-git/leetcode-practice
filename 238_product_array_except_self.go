package main

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	for i := range nums {
		res[i] = 1
	}

	prefix := 1
	for i := 0; i < len(nums); i++ {
		res[i] = prefix
		prefix = prefix * nums[i]
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = res[i] * postfix
		postfix = postfix * nums[i]
	}

	return res
}
