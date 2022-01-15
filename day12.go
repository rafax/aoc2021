package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rafax/aoc2021/util"
)

func sol12_1() {
	m := parseMap()
	acc := [][]string{}
	findPaths("start", "end", m, []string{"start"}, map[string]struct{}{}, &acc)
	fmt.Println(len(acc))
}

func sol12_2() {
	m := parseMap()
	acc := [][]string{}
	findPaths2("start", "end", m, []string{"start"}, map[string]int{}, &acc)
	fmt.Println(len(acc))
}

func findPaths(s, e string, m map[string][]string, path []string, visited map[string]struct{}, acc *[][]string) {
	if _, ok := visited[s]; ok {
		return
	}
	if s == e {
		*acc = append(*acc, path)
		return
	}
	if strings.ToLower(s) == s {
		visited[s] = struct{}{}
	}
	for _, n := range m[s] {
		newVisited := map[string]struct{}{}
		for k, v := range visited {
			newVisited[k] = v
		}
		path = append(path, n)
		findPaths(n, e, m, path, newVisited, acc)
		path = path[:len(path)-1]
	}
}

func findPaths2(s, e string, m map[string][]string, path []string, visited map[string]int, acc *[][]string) {
	last := path[len(path)-1]
	if last == e {
		*acc = append(*acc, path)
		return
	}
	if _, ok := visited[last]; ok {
		for _, vv := range visited {
			if vv == 2 {
				return
			}
		}
	}
	if strings.ToLower(last) == last {
		visited[last]++
	}
	for _, n := range m[last] {
		if n == "start" {
			continue
		}
		newVisited := map[string]int{}
		for k, v := range visited {
			newVisited[k] = v
		}
		newPath := append(path, n)
		findPaths2(s, e, m, newPath, newVisited, acc)
	}
}

func parseMap() map[string][]string {
	m := map[string][]string{}
	util.ReadLine(os.Getenv("INPUT"), func(in string) ([]int, error) {
		parts := strings.Split(in, "-")
		m[parts[0]] = append(m[parts[0]], parts[1])
		m[parts[1]] = append(m[parts[1]], parts[0])
		return nil, nil
	})
	return m
}
