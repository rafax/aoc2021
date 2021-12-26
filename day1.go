package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rafax/aoc2021/util"
)

func sol1_1() {
	in := util.ReadLine(os.Getenv("INPUT"), func(v string) (int, error) {
		return strconv.Atoi(v)
	})
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
	in := util.ReadLine(os.Getenv("INPUT"), func(v string) (int, error) {
		return strconv.Atoi(v)
	})
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
