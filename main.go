package main

import (
	"fmt"
	"log"
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
	"3_1": sol3_1,
	"3_2": sol3_2,
}

func main() {
	if _, ok := sols[os.Args[1]]; !ok {
		log.Fatalf("Solution for %v not found\n", os.Args[1])
	}
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

func sol3_1() {
	in := util.ReadLine(os.Getenv("INPUT"), func(v string) (string, error) {
		return v, nil
	})
	cnt := make([]struct{ zeros, ones int }, len(in[0]), len(in[0]))
	for _, v := range in {
		for i := 0; i < len(v); i++ {
			if v[i] == '1' {
				cnt[i].ones++
			} else {
				cnt[i].zeros++
			}
		}
	}
	lcbs, mcbs := "", ""
	for _, v := range cnt {
		if v.zeros > v.ones {
			lcbs += "1"
			mcbs += "0"
		} else {
			lcbs += "0"
			mcbs += "1"
		}
	}
	fmt.Println(mcbs, lcbs)
	mcb, err := strconv.ParseInt(mcbs, 2, 64)
	check(err, "parse mcb")
	lcb, err := strconv.ParseInt(lcbs, 2, 64)
	check(err, "parse lcb")
	fmt.Println(mcb, lcb, mcb*lcb)

}

func sol3_2() {
	in := util.ReadLine(os.Getenv("INPUT"), func(v string) (string, error) {
		return v, nil
	})
	ox, co2 := in[:], in[:]
	for bit := 0; len(ox) > 1; bit++ {
		ox = filterOx(ox, bit)
	}
	for bit := 0; len(co2) > 1; bit++ {
		co2 = filterCo2(co2, bit)
	}
	o, err := strconv.ParseInt(ox[0], 2, 64)
	check(err, "parse ox")
	c, err := strconv.ParseInt(co2[0], 2, 64)
	check(err, "parse co2")
	fmt.Println(o, c, o*c)
}

func filterOx(in []string, bit int) []string {
	ones, zeros := 0, 0
	for _, v := range in {
		if v[bit] == '1' {
			ones++
		} else {
			zeros++
		}
	}
	if ones >= zeros {
		return haveBit(in, bit, '1')
	} else {
		return haveBit(in, bit, '0')
	}
}

func filterCo2(in []string, bit int) []string {
	ones, zeros := 0, 0
	for _, v := range in {
		if v[bit] == '1' {
			ones++
		} else {
			zeros++
		}
	}
	if ones < zeros {
		return haveBit(in, bit, '1')
	} else {
		return haveBit(in, bit, '0')
	}
}

func haveBit(in []string, bit int, val byte) []string {
	res := []string{}
	for _, v := range in {
		if v[bit] == val {
			res = append(res, v)
		}
	}
	return res
}

func check(err error, ctx string) {
	if err != nil {
		log.Fatalln(ctx, err)
	}
}
