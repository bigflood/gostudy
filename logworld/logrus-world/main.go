package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetReportCaller(true)

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	//log.SetLevel(log.WarnLevel)
}
func main() {

	logger := log.WithFields(log.Fields{
		"animal": "walrus",
	})

	logger = logger.WithField("name", "dog")
	logger.Println("println~")

	SomeFunc(logger)
}

func SomeFunc(logger *log.Entry) {
	logger.Info("A walrus appears", " and disappears")
	logger.Info(struct{
		Field1, Field2 string
	}{Field1:"123", Field2:"abc"})

	logger.Panic("panic!!!!")
}
