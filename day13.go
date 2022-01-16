package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rafax/aoc2021/util"
)

type dotPaper struct {
	dots          []point
	width, height int
	folds         []fold
}

type fold struct {
	alongX bool
	pos    int
}

func sol13_1() {
	p := parsePaper()
	fmt.Println(p)
	fmt.Println()
	p.fold()
	fmt.Println(p)
	fmt.Println("# of dots:", len(p.dots))
}

func sol13_2() {
	p := parsePaper()
	for len(p.folds) > 0 {
		p.fold()
	}
	fmt.Println(p)
}

func (p dotPaper) String() string {
	b := bytes.Buffer{}
	fmt.Fprintln(&b, p.dots)
	m := map[point]struct{}{}
	for _, d := range p.dots {
		m[d] = struct{}{}
	}
	for i := 0; i < p.height*p.width; i++ {
		if i > 0 && i%p.width == 0 {
			fmt.Fprintln(&b)
		}
		if _, ok := m[point{x: i % p.width, y: i / p.width}]; ok {
			fmt.Fprint(&b, "#")
		} else {
			fmt.Fprint(&b, ".")
		}
	}
	return b.String()
}

func (p *dotPaper) fold() {
	f := p.folds[0]
	p.folds = p.folds[1:]
	// copy
	for _, d := range p.dots {
		if f.alongX {
			if d.x > f.pos {
				p.dots = append(p.dots, point{x: f.pos - (d.x - f.pos), y: d.y})
			}
		} else {
			if d.y > f.pos {
				p.dots = append(p.dots, point{x: d.x, y: f.pos - (d.y - f.pos)})
			}
		}
	}
	// dedupe
	if f.alongX {
		p.width = f.pos
	} else {
		p.height = f.pos
	}
	new := map[point]struct{}{}
	for _, d := range p.dots {
		if _, ok := new[d]; ok {
			continue
		}
		if d.x > p.width {
			continue
		}
		if !f.alongX && d.y >= f.pos {
			continue
		}
		new[d] = struct{}{}
	}
	newDots := []point{}
	for k := range new {
		newDots = append(newDots, k)
	}
	p.dots = newDots
}

func parsePaper() dotPaper {
	m := dotPaper{}
	dotMode := true
	util.ReadLine(os.Getenv("INPUT"), func(in string) ([]int, error) {
		if len(in) == 0 {
			dotMode = false
			return nil, nil
		}
		if dotMode {
			parts := strings.Split(in, ",")
			x, err := strconv.Atoi(parts[0])
			check(err, "parsing x")
			if x+1 > m.width {
				m.width = x + 1
			}
			y, err := strconv.Atoi(parts[1])
			if y+1 > m.height {
				m.height = y + 1
			}
			check(err, "parsing y")
			m.dots = append(m.dots, point{x: x, y: y})
		} else {
			in = strings.TrimPrefix(in, "fold along ")
			parts := strings.Split(in, "=")
			alongX := parts[0] == "x"
			pos, err := strconv.Atoi(parts[1])
			check(err, "parsing pos")
			m.folds = append(m.folds, fold{alongX: alongX, pos: pos})
		}
		return nil, nil
	})
	return m
}
