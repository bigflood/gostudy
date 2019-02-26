package main

import (
	"fmt"
	"net/http"
	"net/url"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, WebAssembly!")

	js.Global().Set("onClickImg", js.NewCallback(
		func(args []js.Value) {
			x := args[0].Int()
			y := args[1].Int()
			go onClickImg(x, y)
		},
	))

	select {}
}

func onClickImg(x, y int) {
	u, _ := url.Parse(js.Global().Get("window").Get("location").Get("href").String())
	u.Path = "/click"
	u.RawQuery = fmt.Sprintf("x=%v&y=%v", x, y)

	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println("Get failed:", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Get failed:", resp.Status)
		return
	}
}
