package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	a := make([]int, 0)
	for i := 0; i < 1; i++ {
		a = append(a, i)
	}
	t := time.Since(t1)
	fmt.Println(t)
}
