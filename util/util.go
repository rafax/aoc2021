package util

import (
	"bufio"
	"log"
	"os"
)

func ReadLine[T any](in string, parser func(string) (T, error)) []T {
	file, err := os.Open(in)
	check(err, "open input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := []T{}
	for scanner.Scan() {
		v, err := parser(scanner.Text())
		check(err, "parse ")
		res = append(res, v)
	}

	if err := scanner.Err(); err != nil {
		check(err, "scanning file")
	}
	return res
}

func check(err error, ctx string) {
	if err != nil {
		log.Fatalln(ctx, err)
	}
}
