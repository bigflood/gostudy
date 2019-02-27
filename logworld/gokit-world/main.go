package main

import (
	"os"

	"github.com/go-kit/kit/log"
)

func main() {

	logger := log.NewJSONLogger(os.Stdout)

	logger = log.With(
		logger,
		"component", "handler",
		"mytime", log.DefaultTimestamp,
	)

	logger.Log(
		"url", "http://asdf-dfgdfg-sdf.com/asdfsdf/123",
		"msg", "hello go-kit",
	)
}
