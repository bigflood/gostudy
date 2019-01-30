package main

import (
	"log"
	"math/rand"
	_ "net/http/pprof"

	"github.com/bigflood/gostudy/gameoflife/api"
	"github.com/bigflood/gostudy/gameoflife/sim"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	addr   = ":8080"
	width  = 500
	height = 500
)

func main() {

	rand.Seed(1234)

	sim := sim.New(width, height)

	go sim.EncodeImages()

	handlers := api.New(sim)

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Static("/static", "static")
	e.File("/", "static/index.html")
	e.GET("/image.png", handlers.WriteImage)

	log.Println("listen:", addr)
	e.Logger.Fatal(e.Start(addr))
}
