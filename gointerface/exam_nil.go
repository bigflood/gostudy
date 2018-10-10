package main

import "fmt"

type MyErr struct {
	errCode int
}

func (p *MyErr) Error() string {
	if p == nil {
		return "nil"
	}

	return fmt.Sprintf("code=%v", p.errCode)
}

func doTask() error {
	return &MyErr{errCode:400}
}

type MyErrCode int

func (p MyErrCode) Error() string {
	return fmt.Sprintf("code=%v", int(p))
}

func doTask2() error {
	return MyErrCode(400)
}

func main() {
	var a *MyErr

	fmt.Printf("a.Errror(): %v \n", a.Error())

	fmt.Printf("a is nil: %v \n", a == nil)

	// interface
	// [type, *data]

	// [nil, nil] == nil

	var e error
	e = a

	// [*MyErr, nil] != nil

	fmt.Printf("e is nil: %v \n", e == nil)
}
