package main

import (
	"github.com/bigflood/gostudy/gameoflife/sim"
	"log"
	"math/rand"
	"os"
	"runtime/trace"
)

const (
	width  = 50
	height = 50
)

func main() {
	traceFile, err := os.Create("sim.trace")
	if err != nil {
		log.Fatal(err)
	}
	defer traceFile.Close()

	trace.Start(traceFile)
	defer trace.Stop()

	rand.Seed(1234)

	s := sim.New(width, height)

	for index := 0; index < 5; index++ {
		s.WaitForFrame()
		s.EncodeImage()
		log.Println("frame:", index)
	}
}
