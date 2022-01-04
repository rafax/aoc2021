package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/rafax/aoc2021/util"
)

type entry struct {
	patterns []string
	output   []string
}

func sol8_1() {
	entries := parseEntries()
	fmt.Println(entries)
	lengths := map[int]bool{2: true, 3: true, 4: true, 7: true}
	cnt := 0
	for _, e := range entries {
		for _, o := range e.output {
			if _, ok := lengths[len(o)]; ok {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}

func sol8_2() {
	entries := parseEntries()
	cnt := 0
	for _, e := range entries {
		letters := guessLetters(e.patterns)
		sorted := map[string]string{}
		for l, v := range letters {
			sorted[sortLetters(l)] = v
		}
		out := ""
		for _, o := range e.output {
			s := sortLetters(o)
			n := sorted[s]
			out += n
		}
		v, err := strconv.Atoi(out)
		check(err, "parsing output")
		cnt += v
	}
	fmt.Println(cnt)
}

func guessLetters(in []string) map[string]string {
	letters := map[string]string{}

	one := findWithLength(in, 2)[0]
	letters[one] = "1"

	seven := findWithLength(in, 3)[0]
	letters[seven] = "7"

	four := findWithLength(in, 4)[0]
	letters[four] = "4"

	eight := findWithLength(in, 7)[0]
	letters[eight] = "8"

	lsix := findWithLength(in, 6)
	zero := ""
	for _, v := range lsix {
		diff := minus(eight, v)
		if !containsAny(seven, diff) {
			if containsAny(four, diff) {
				zero = v
				letters[zero] = "0"
				lsix = remove(lsix, zero)
			}
		}
	}

	lfive := findWithLength(in, 5)
	three := ""
	for _, v := range lfive {
		diff := minus(eight, v)
		if !containsAny(one, diff) {
			three = v
			letters[three] = "3"
			lfive = remove(lfive, three)
		}
	}

	nine := ""
	for _, v := range lsix {
		diff := minus(eight, v)
		if !containsAny(four, diff) {
			nine = v
			letters[nine] = "9"
			lsix = remove(lsix, nine)
		}
	}

	if len(lsix) != 1 {
		log.Fatalln(lsix)
	}
	six := lsix[0]
	letters[six] = "6"

	two := ""
	diff := minus(eight, six)

	for _, v := range lfive {
		if containsAny(v, diff) {
			two = v
			letters[two] = "2"
			lfive = remove(lfive, two)
		}
	}
	five := lfive[0]
	letters[five] = "5"

	return letters
}

func remove(in []string, v string) []string {
	res := []string{}
	for _, e := range in {
		if e != v {
			res = append(res, e)
		}
	}
	return res
}

func minus(left, right string) string {
	diff := []string{}
	for i := range left {
		v := string(left[i])
		if !strings.Contains(right, v) {
			diff = append(diff, v)
		}
	}
	return strings.Join(diff, "")
}

func containsAny(left, right string) bool {
	for i := range right {
		v := string(right[i])
		if strings.Contains(left, v) {
			return true
		}
	}
	return false
}

func findWithLength(in []string, l int) []string {
	res := []string{}
	for _, v := range in {
		if len(v) == l {
			res = append(res, v)
		}
	}
	return res
}

func sortLetters(o string) string {
	s := strings.Split(o, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func parseEntries() []entry {
	return util.ReadLine(os.Getenv("INPUT"), func(in string) (entry, error) {
		parts := strings.Split(in, "|")
		p := strings.Split(strings.TrimSpace(parts[0]), " ")
		o := strings.Split(strings.TrimSpace(parts[1]), " ")
		return entry{patterns: p, output: o}, nil
	})
}
