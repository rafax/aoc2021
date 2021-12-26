package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadInts(in string) []int {
	file, err := os.Open(in)
	check(err, "open input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := []int{}
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		check(err, "parse int")
		res = append(res, v)
	}

	if err := scanner.Err(); err != nil {
		check(err, "scanning file")
	}
	return res
}

func check(err error, ctx string) {
	if err != nil {
		log.Fatal(ctx, err)
	}
}
