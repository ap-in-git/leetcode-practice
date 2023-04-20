package main

import (
	"fmt"
	"testing"
)

type question42 struct {
	para42
	ans42
}

type para42 struct {
	one []int
}

type ans42 struct {
	one int
}

func Test_Problem42(t *testing.T) {

	qs := []question42{

		{
			para42{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
			ans42{6},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 42------------------------\n")

	for _, q := range qs {
		_, p := q.ans42, q.para42
		fmt.Printf("【input】:%v       【output】:%v\n", p, trap(p.one))
	}
	fmt.Printf("\n\n\n")
}
