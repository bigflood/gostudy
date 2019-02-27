package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	cfg := zap.NewDevelopmentConfig()
	//cfg := zap.NewProductionConfig()
	cfg.Encoding = "console"

	logger, _ := cfg.Build()
	//logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	url := "http://adsfadf-asdfer.com"
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
		"req", struct{Field1 string}{},
	)
	sugar.Infof("Failed to fetch URL: %s", url)

	//logger.Named()
	fastLog(logger)
}

func fastLog(logger *zap.Logger) {
	url := "http://adsfadf-asdfer.com"

	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

