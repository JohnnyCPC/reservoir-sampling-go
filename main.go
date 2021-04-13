package main

import (
	"flag"
	"fmt"

	"github.com/JohnnyCPC/reservoir-sampling-go/sks"
)

func main() {

	nFlag := flag.Int("n", 39, "n items")
	k := flag.Int("k", 5, "choosing k samples")

	flag.Parse()

	a := make([]int, *nFlag)
	for i := range a {
		a[i] = i + 1
	}
	ln := len(a)
	fmt.Println(*nFlag, *k)

	fmt.Println(sks.SelectKItems(a, ln, *k))
}

func printResult(s []int, n int) {
	for i, v := range s {
		if i < n {
			fmt.Print(v, " ")
		}
	}
	fmt.Println("")
}
