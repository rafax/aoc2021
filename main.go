package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rafax/aoc2021/util"
)

var sols = map[string]func(){
	"1_1": sol1_1,
	"1_2": sol1_2,
	"2_1": sol2_1,
	"2_2": sol2_2,
}

func main() {
	sols[os.Args[1]]()
}

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

type Direction string

const (
	Forward Direction = "forward"
	Down              = "down"
	Up                = "up"
)

type Vector struct {
	Direction Direction
	Length    int
}

func sol2_1() {
	in := parseVector()
	f, d := 0, 0
	for _, v := range in {
		switch v.Direction {
		case Forward:
			f += v.Length
		case Down:
			d += v.Length
		case Up:
			d -= v.Length
		}
	}
	fmt.Println(f, d)
	fmt.Println(f * d)
}

func sol2_2() {
	in := parseVector()
	f, d, aim := 0, 0, 0
	for _, v := range in {
		switch v.Direction {
		case Forward:
			f += v.Length
			d += aim * v.Length
		case Down:
			aim += v.Length
		case Up:
			aim -= v.Length
		}
	}
	fmt.Println(f, d)
	fmt.Println(f * d)
}

func parseVector() []*Vector {
	return util.ReadLine(os.Getenv("INPUT"), func(v string) (*Vector, error) {
		parts := strings.SplitN(v, " ", 2)
		d, l := parts[0], parts[1]
		len, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		return &Vector{Direction: Direction(d), Length: len}, nil
	})
}
