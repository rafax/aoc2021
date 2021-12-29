package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/rafax/aoc2021/util"
)

type point struct {
	x, y int
}

type line struct {
	start, end point
}

func sol5_1() {
	lines := parseLines()
	covered := map[point]int{}
	for _, l := range lines {
		if l.horizontal() || l.vertical() {
			l.forEachPoint(func(p point) {
				covered[p] += 1
			})
		}
	}
	fmt.Println(countMoreThan(covered, 2))
}

func sol5_2() {
	lines := parseLines()
	covered := map[point]int{}
	for _, l := range lines {
		l.forEachPoint(func(p point) {
			covered[p] += 1
		})
	}
	fmt.Println(countMoreThan(covered, 2))
}

func parseLines() []line {
	return util.ReadLine(os.Getenv("INPUT"), func(in string) (line, error) {
		parts := strings.Split(in, "->")
		start, err := parsePoint(strings.TrimSpace(parts[0]))
		if err != nil {
			return line{}, err
		}
		end, err := parsePoint(strings.TrimSpace(parts[1]))
		if err != nil {
			return line{}, err
		}
		return line{start: start, end: end}, nil
	})
}

func parsePoint(in string) (point, error) {
	parts := strings.Split(in, ",")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return point{}, err
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return point{}, err
	}
	return point{x: x, y: y}, nil
}

func countMoreThan(c map[point]int, limit int) int {
	cnt := 0
	for _, n := range c {
		if n >= limit {
			cnt++
		}
	}
	return cnt
}

func (l line) vertical() bool {
	return l.start.x == l.end.x
}

func (l line) diagonal() bool {
	return math.Abs(float64(l.end.x-l.start.x)) == math.Abs(float64(l.end.y-l.start.y))
}

func (l line) horizontal() bool {
	return l.start.y == l.end.y
}

func (l line) forEachPoint(p func(point)) {
	if l.horizontal() {
		var from, to int
		if l.start.x < l.end.x {
			from, to = l.start.x, l.end.x
		} else {
			from, to = l.end.x, l.start.x
		}
		for i := from; i <= to; i++ {
			p(point{x: i, y: l.start.y})
		}
		return
	}
	if l.vertical() {
		var from, to int
		if l.start.y < l.end.y {
			from, to = l.start.y, l.end.y
		} else {
			from, to = l.end.y, l.start.y
		}
		for i := from; i <= to; i++ {
			p(point{x: l.start.x, y: i})
		}
		return
	}
	if l.diagonal() {
		dx := diff(l.start.x, l.end.x)
		dy := diff(l.start.y, l.end.y)
		x, y := l.start.x, l.start.y
		for x != l.end.x && y != l.end.y {
			p(point{x: x, y: y})
			x += dx
			y += dy
		}
		p(l.end)
		return
	}
}

func diff(x, y int) int {
	if x < y {
		return 1
	}
	if x > y {
		return -1
	}
	return 0
}
