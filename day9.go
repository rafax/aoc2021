package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/rafax/aoc2021/util"
)

type lavaPoint struct {
	x, y, height int
}

func sol9_1() {
	entries := parseLavaPoints()
	low := findLowPoints(entries)
	cnt := 0
	for _, l := range low {
		cnt += (l.height + 1)
	}
	fmt.Println(cnt)
}

func sol9_2() {
	entries := parseLavaPoints()
	low := findLowPoints(entries)
	basinMap := findBasins(entries, low)
	bsize := make([]int, len(low))
	for _, basinPoint := range basinMap {
		bsize[basinPoint]++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(bsize)))
	fmt.Println(bsize[0] * bsize[1] * bsize[2])
}

func findBasins(entries [][]int, low []lavaPoint) map[lavaPoint]int {
	res := map[lavaPoint]int{}
	for i, l := range low {
		res = markBasin(res, entries, l, i)
	}
	return res
}

func markBasin(res map[lavaPoint]int, entries [][]int, l lavaPoint, basin int) map[lavaPoint]int {
	res[l] = basin

	for _, p := range adjacent(entries, l.x, l.y) {
		if _, ok := res[p]; !ok {
			if p.height != 9 {
				res[p] = basin
				res = markBasin(res, entries, p, basin)
			}
		}
	}
	return res
}

func findLowPoints(in [][]int) []lavaPoint {
	res := []lavaPoint{}
	for x := 0; x < len(in); x++ {
		for y := 0; y < len(in[0]); y++ {
			if isLow(in, x, y) {
				res = append(res, lavaPoint{x: x, y: y, height: in[x][y]})
			}
		}
	}
	return res
}

func isLow(in [][]int, x, y int) bool {
	for _, v := range adjacent(in, x, y) {
		if in[x][y] >= v.height {
			return false
		}
	}
	return true
}

func adjacent(in [][]int, x, y int) []lavaPoint {
	res := []lavaPoint{}
	if x > 0 {
		res = append(res, lavaPoint{x: x - 1, y: y, height: in[x-1][y]})
	}
	if x < len(in)-1 {
		res = append(res, lavaPoint{x: x + 1, y: y, height: in[x+1][y]})
	}
	if y > 0 {
		res = append(res, lavaPoint{x: x, y: y - 1, height: in[x][y-1]})
	}
	if y < len(in[0])-1 {
		res = append(res, lavaPoint{x: x, y: y + 1, height: in[x][y+1]})
	}
	return res
}

func parseLavaPoints() [][]int {
	return util.ReadLine(os.Getenv("INPUT"), func(in string) ([]int, error) {
		res := []int{}
		for _, h := range in {
			v, err := strconv.Atoi(string(h))
			if err != nil {
				return nil, err
			}
			res = append(res, v)
		}
		return res, nil
	})
}
