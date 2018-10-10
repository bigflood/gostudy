package main

import (
	"github.com/bigflood/gostudy/gointerface/todo"
	"log"
)

type Todo interface {
	Add(desc string) error
	List() ([]todo.Task, error)
	Done(index int) error
}

type Logging struct {
	next Todo
}

func (p *Logging) Add(desc string) error {
	log.Println("Add: ", desc)
	return p.next.Add(desc)
}

func (p *Logging) List() ([]todo.Task, error) {
	log.Println("List: ")
	return p.next.List()
}

func (p *Logging) Done(index int) error {
	log.Println("Done: ", index)
	return p.next.Done(index)
}

func main() {
	//obj := todo.NewInMem()
	obj := todo.NewInFile("todo.json")
	example(&Logging{next: obj})
}

func example(p Todo) {
	p.Add("task1")
	p.Add("task2")

	log.Println(p.List())

	p.Done(0)

	log.Println(p.List())
}
