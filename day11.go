package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/rafax/aoc2021/util"
)

type energyMap struct {
	octopii      []int
	width        int
	flashed      map[int]struct{}
	totalFlashes int
}

func sol11_1() {
	m := parseOctopusEnergy()
	fmt.Println(m)
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		m.step()
		fmt.Println(m)
	}
	fmt.Println(m.totalFlashes)
}

func sol11_2() {
	m := parseOctopusEnergy()
	for i := 0; ; i++ {
		m.step()
		if len(m.flashed) == len(m.octopii) {
			fmt.Println(i + 1)
			break
		}
	}
}

func parseOctopusEnergy() *energyMap {
	m := energyMap{octopii: []int{}}
	util.ReadLine(os.Getenv("INPUT"), func(in string) ([]int, error) {
		res := []int{}
		for _, h := range in {
			if m.width == 0 {
				m.width = len(in)
			}
			v, err := strconv.Atoi(string(h))
			if err != nil {
				return nil, err
			}
			res = append(res, v)
		}
		m.octopii = append(m.octopii, res...)
		return res, nil
	})
	return &m
}

func (m *energyMap) String() string {
	var b bytes.Buffer
	for row := 0; row < len(m.octopii)/m.width; row++ {
		for i := 0; i < m.width; i++ {
			fmt.Fprint(&b, m.octopii[row*m.width+i])
		}
		fmt.Fprintln(&b)
	}
	return b.String()
}

func (m *energyMap) step() {
	m.flashed = map[int]struct{}{}
	for i, v := range m.octopii {
		m.octopii[i] = (v + 1)
	}
	for {
		changed := false
		for i, v := range m.octopii {
			if v > 9 {
				m.octopii[i] = 0
				m.flashed[i] = struct{}{}
				m.increaseAdjacent(i)
				changed = true
				m.totalFlashes++
			}
		}
		if !changed {
			break
		}
	}
	for i := range m.flashed {
		m.octopii[i] = 0
	}
}

func (m *energyMap) increaseAdjacent(i int) {
	row, col := i/m.width, i%m.width
	neighbours := []int{}
	if row > 0 {
		if col > 0 {
			neighbours = append(neighbours, i-m.width-1)
		}
		neighbours = append(neighbours, i-m.width)
		if col < m.width-1 {
			neighbours = append(neighbours, i-m.width+1)
		}
	}
	if col > 0 {
		neighbours = append(neighbours, i-1)
	}
	if col < m.width-1 {
		neighbours = append(neighbours, i+1)
	}
	if row < (len(m.octopii)/m.width)-1 {
		if col > 0 {
			neighbours = append(neighbours, i+m.width-1)
		}
		neighbours = append(neighbours, i+m.width)
		if col < m.width-1 {
			neighbours = append(neighbours, i+m.width+1)
		}
	}

	for _, v := range neighbours {
		if v >= 0 && v < len(m.octopii) {
			if _, ok := m.flashed[v]; !ok {
				m.octopii[v]++
			}
		}
	}
}
