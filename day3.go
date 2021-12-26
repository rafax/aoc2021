package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rafax/aoc2021/util"
)

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
