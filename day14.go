package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/rafax/aoc2021/util"
)

type polymers struct {
	state string
	rules map[string]string
}

func sol14_1() {
	p := parsePolymers()
	for i := 0; i < 10; i++ {
		p.transform()
	}
	lce, mce := findFrequentElements(p.state)
	fmt.Println(mce - lce)
}

func sol14_2() {
	p := parsePolymers()
	for i := 0; i < 40; i++ {
		// if i%10 == 0 && i > 0 {
		fmt.Println(i, len(p.state))
		// }
		p.transform()
	}
	lce, mce := findFrequentElements(p.state)
	fmt.Println(mce - lce)
}

func findFrequentElements(s string) (int, int) {
	cnt := map[rune]int{}
	for _, v := range s {
		cnt[v]++
	}
	lce, mce := math.MaxInt, -1
	for _, v := range cnt {
		if v < lce {
			lce = v
		}
		if v > mce {
			mce = v
		}
	}
	return lce, mce
}

func (p *polymers) transform() {
	new := strings.Builder{}
	for i := 0; i < len(p.state)-1; i++ {
		new.WriteByte(p.state[i])
		in := p.state[i : i+2]
		if v, ok := p.rules[in]; ok {
			new.WriteString(v)
		}
	}
	new.WriteByte(p.state[len(p.state)-1])
	p.state = new.String()
}

func parsePolymers() polymers {
	p := polymers{rules: map[string]string{}}
	inputMode := true
	util.ReadLine(os.Getenv("INPUT"), func(in string) ([]int, error) {
		if len(in) == 0 {
			inputMode = false
			return nil, nil
		}
		if inputMode {
			p.state = in
		} else {
			parts := strings.Split(in, " -> ")
			p.rules[parts[0]] = parts[1]
		}
		return nil, nil
	})
	return p
}
