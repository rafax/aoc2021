package main

import (
	"log"
	"os"
)

var sols = map[string]func(){
	"1_1": sol1_1,
	"1_2": sol1_2,
	"2_1": sol2_1,
	"2_2": sol2_2,
	"3_1": sol3_1,
	"3_2": sol3_2,
	"4_1": sol4_1,
	"4_2": sol4_2,
	"5_1": sol5_1,
	"5_2": sol5_2,
}

func main() {
	if _, ok := sols[os.Args[1]]; !ok {
		log.Fatalf("Solution for %v not found\n", os.Args[1])
	}
	sols[os.Args[1]]()
}

func check(err error, ctx string) {
	if err != nil {
		log.Fatalln(ctx, err)
	}
}
