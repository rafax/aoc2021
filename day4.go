package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rafax/aoc2021/util"
)

type game struct {
	curr   int
	marked []int
	boards []board
}

type value struct {
	number int
	marked bool
}

type board []value

func sol4_1() {
	game := parse()
	fmt.Println(game)
	won := false
	score := 0
	for ; !won; won, score = game.wonBoardScore() {
		game.mark()
		fmt.Println(game)
	}
	fmt.Println(score)
}

func sol4_2() {
	game := parse()
	fmt.Println(game)
	unwonBoardIndices := map[int]bool{}
	for i := 0; i < len(game.boards); i++ {
		unwonBoardIndices[i] = true
	}
	for {
		game.mark()
		for i, b := range game.boards {
			if b.won() {
				if len(unwonBoardIndices) == 1 {
					if unwonBoardIndices[i] == true {
						// this is the board that won the last
						fmt.Println(b.sumUnmarked() * game.marked[game.curr])
						return
					}
				}
				delete(unwonBoardIndices, i)
			}

		}
	}
}

func parse() game {
	g := game{curr: -1}
	in := util.ReadLines(os.Getenv("INPUT"))
	for _, v := range strings.Split(in[0], ",") {
		n, err := strconv.Atoi(v)
		check(err, "parsing marked")
		g.marked = append(g.marked, n)
	}
	for i := 2; i < len(in); i += 6 {
		board := []value{}
		for j := 0; j < 5; j++ {
			z := 0
			for x := 0; x < len(in[i+j]); x += 3 {
				v := strings.Trim(in[i+j][x:x+2], " ")
				n, err := strconv.Atoi(v)
				check(err, "parsing board")
				board = append(board, value{number: n, marked: false})
				z++
			}
		}
		g.boards = append(g.boards, board)
	}
	return g
}

func (b board) String() string {
	res := bytes.Buffer{}
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			fmt.Fprint(&res, b[x+y*5])
			if b[x+y*5].marked {
				fmt.Fprint(&res, "*")
			}
		}
		fmt.Fprintln(&res)
	}
	return res.String()
}

func (b board) mark(c int) {
	for i := 0; i < len(b); i++ {
		v := &b[i]
		if v.number == c {
			v.marked = true
		}
	}
}

func (b board) won() bool {
	// rows
	for i := 0; i < 5; i++ {
		all := true
		for j := 0; j < 5; j++ {
			if !b[i*5+j].marked {
				all = false
				break
			}
		}
		if all {
			return true
		}
	}
	// cols
	for i := 0; i < 5; i++ {
		all := true
		for j := 0; j < 5; j++ {
			if !b[i+j*5].marked {
				all = false
				break
			}
		}
		if all {
			return true
		}
	}
	return false
}

func (b board) sumUnmarked() int {
	sum := 0
	for _, v := range b {
		if !v.marked {
			sum += v.number
		}
	}
	return sum
}

func (g game) String() string {
	res := bytes.Buffer{}
	fmt.Fprintln(&res, g.marked)
	for _, v := range g.boards {
		fmt.Fprintln(&res, v)
	}
	return res.String()
}

func (g *game) mark() {
	g.curr++
	c := g.marked[g.curr]
	for _, b := range g.boards {
		b.mark(c)
	}
}

func (g game) wonBoardScore() (bool, int) {
	for _, b := range g.boards {
		if b.won() {
			return true, b.sumUnmarked() * g.marked[g.curr]
		}
	}
	return false, 0
}
