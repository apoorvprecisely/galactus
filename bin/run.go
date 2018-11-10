package main

import (
	"github.com/apoorvprecisely/galactus"
	"github.com/apoorvprecisely/galactus/reader"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	gal, err := galactus.New();
	if err != nil {
		sugar.Errorf("Failed to start galactus", err)
	}
	reader.Start(gal.Config)
	// start kafka consumers for events
}
