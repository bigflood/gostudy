package main

import (
	"os"

	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = ""

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	logger.Debug().
		Str("Scale", "833 cents").
		Float64("Interval", 833.09).
		Interface("i", "abc").
		Msg("Fibonacci is everywhere")
}