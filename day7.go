package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/rafax/aoc2021/util"
)

func sol7_1() {
	crabs := parseLocation()
	fmt.Println(crabs, optimalCrabPosition(crabs, linearCost))
}

func sol7_2() {
	crabs := parseLocation()
	fmt.Println(crabs, optimalCrabPosition(crabs, stepCost))
}

func optimalCrabPosition(crabs []int, cost func([]int, int) int) int {
	min := math.MaxInt
	start, end := minMax(crabs)
	for i := start; i <= end; i++ {
		cost := cost(crabs, i)
		if cost < min {
			min = cost
		}
	}
	return min
}

func linearCost(crabs []int, pos int) int {
	cost := 0
	for _, v := range crabs {
		cost += int(math.Abs(float64(v - pos)))
	}
	return cost
}

func stepCost(crabs []int, pos int) int {
	cost := 0
	for _, v := range crabs {
		diff := int(math.Abs(float64(v - pos)))
		cost += (diff + 1) * diff / 2
	}
	return cost
}

func parseLocation() []int {
	lines := util.ReadLines(os.Getenv("INPUT"))
	crabs := []int{}
	locations := strings.Split(lines[0], ",")
	for _, d := range locations {
		v, err := strconv.Atoi(d)
		check(err, "parsing fish")
		crabs = append(crabs, v)
	}
	return crabs
}

func minMax(in []int) (int, int) {
	var max int = in[0]
	var min int = in[0]
	for _, value := range in {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
