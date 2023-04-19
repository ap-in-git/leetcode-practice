package main

import (
	"fmt"
	"testing"
)

type question128 struct {
	para128
	ans128
}

type para128 struct {
	one []int
}

type ans128 struct {
	one int
}

func Test_Problem128(t *testing.T) {

	qs := []question128{

		{
			para128{[]int{}},
			ans128{0},
		},

		{
			para128{[]int{0}},
			ans128{1},
		},

		{
			para128{[]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}},
			ans128{7},
		},

		{
			para128{[]int{2147483646, -2147483647, 0, 2, 2147483644, -2147483645, 2147483645}},
			ans128{3},
		},

		{
			para128{[]int{100, 4, 200, 1, 3, 2}},
			ans128{4},
		},
		{
			para128{[]int{}},
			ans128{4},
		},
	}

	for _, q := range qs {
		_, p := q.ans128, q.para128
		fmt.Printf("【input】:%v    [expected]:%v   【output】:%v\n", p, q.ans128, longestConsecutive(p.one))
	}
	fmt.Printf("\n\n\n")
}
