package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	n := len(a)
	k := 2

	selectKItems(a, n, k)
}

func selectKItems(streams []int, n int, k int) {

	var i int
	reservoir := make([]int, k)
	rand.Seed(time.Now().UnixNano())

	for i = 0; i < k; i++ {
		reservoir[i] = streams[i]
	}
	for ; i < n; i++ {
		j := rand.Intn(i + 1)
		//fmt.Println(i+1, j)
		if j < k {
			//fmt.Println("j<k:", j, k, streams[i])
			reservoir[j] = streams[i]
		}
	}
	fmt.Println(reservoir)
	//printResult(reservoir, k)
}

func printResult(s []int, n int) {
	for i, v := range s {
		if i < n {
			fmt.Print(v, " ")
		}
	}
	fmt.Println("")
}
