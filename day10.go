package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/rafax/aoc2021/util"
)

var (
	pairs      = map[rune]rune{'[': ']', '{': '}', '(': ')', '<': '>'}
	cost       = map[rune]int{']': 57, '}': 1197, ')': 3, '>': 25137}
	appendCost = map[rune]int{'[': 2, '{': 3, '(': 1, '<': 4}
)

func sol10_1() {
	lines := util.ReadLines(os.Getenv("INPUT"))
	errs := []rune{}
	for _, l := range lines {
		stack := []rune{}
		for _, c := range l {
			if _, ok := pairs[c]; ok {
				stack = append(stack, c)
				continue
			}
			if c == ']' || c == '}' || c == ')' || c == '>' {
				if pairs[stack[len(stack)-1]] == c {
					stack = stack[:len(stack)-1]
				} else {
					errs = append(errs, c)
					break
				}
			}
		}
	}
	sum := 0
	for _, v := range errs {
		sum += cost[v]
	}
	fmt.Println(sum)
}

func sol10_2() {
	lines := util.ReadLines(os.Getenv("INPUT"))
	scores := []int{}
	for _, l := range lines {
		stack := []rune{}
		isErr := false
		for _, c := range l {
			if _, ok := pairs[c]; ok {
				stack = append(stack, c)
				continue
			}
			if c == ']' || c == '}' || c == ')' || c == '>' {
				if pairs[stack[len(stack)-1]] == c {
					stack = stack[:len(stack)-1]
				} else {
					isErr = true
					break
				}
			}
		}
		if !isErr && len(stack) > 0 {
			tot := 0
			for i := len(stack) - 1; i >= 0; i-- {
				tot *= 5
				tot += appendCost[stack[i]]
			}
			scores = append(scores, tot)
		}
	}
	sort.Sort(sort.IntSlice(scores))

	fmt.Println(scores[(len(scores) / 2)])
}
