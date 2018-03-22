package main

import (
	"fmt"
	"os"

	"github.com/bigflood/gostudy/gocalc"
)

func main() {
	for _, src := range os.Args[1:] {
		fmt.Println("src=", src)

		v, err := gocalc.CalcExpr(src)
		if err != nil {
			panic(err)
		}

		fmt.Println("result=", v)
	}
}
