package main

import (
	"fmt"
	"os"

	"github.com/rafax/aoc2021/util"
)

var sols = map[string]func(){
	"1_1": sol1_1,
	"1_2": sol1_2,
}

func main() {
	sols[os.Args[1]]()
}

func sol1_1() {
	in := util.ReadInts(os.Getenv("INPUT"))
	cnt := 0
	prev := in[0]
	for _, v := range in[1:] {
		if v > prev {
			cnt++
		}
		prev = v
	}
	fmt.Println(cnt)
}

func sol1_2() {
	in := util.ReadInts(os.Getenv("INPUT"))
	cnt := 0
	prev := in[0] + in[1] + in[2]
	for i, v := range in[3:] {
		curr := prev - in[i] + v
		if curr > prev {
			cnt++
		}
		prev = curr
	}
	fmt.Println(cnt)
}
