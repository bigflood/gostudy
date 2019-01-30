package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, WebAssembly!")

	js.Global().Set("onClickImg", js.NewCallback(
		func(args []js.Value) {
			x := args[0].Int()
			y := args[1].Int()
			onClickImg(x, y)
		},
	))

	select {}
}

func onClickImg(x, y int) {
	fmt.Println("onClickImg:", x, y)

}
