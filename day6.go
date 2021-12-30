package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rafax/aoc2021/util"
)

func sol6_1() {
	fish := parseFish()
	sum := simulate(fish, 80)
	fmt.Println(sum)
}

func sol6_2() {
	fish := parseFish()
	sum := simulate(fish, 256)
	fmt.Println(sum)
}

func simulate(fish []int, rounds int) int {
	fishCount := map[int]int{}
	for _, v := range fish {
		fishCount[v] += 1
	}
	for i := 0; i < rounds; i++ {
		newFishCount := map[int]int{}
		if fishCount[0] > 0 {
			newFishCount[8] += fishCount[0]
			newFishCount[6] += fishCount[0]
		}
		for i := 1; i < 9; i++ {
			if fishCount[i] != 0 {
				newFishCount[i-1] += fishCount[i]
			}
		}
		fishCount = newFishCount
	}
	sum := 0
	for _, v := range fishCount {
		sum += v
	}
	return sum
}

func parseFish() []int {
	lines := util.ReadLines(os.Getenv("INPUT"))
	fish := []int{}
	days := strings.Split(lines[0], ",")
	for _, d := range days {
		v, err := strconv.Atoi(d)
		check(err, "parsing fish")
		fish = append(fish, v)
	}
	return fish
}
