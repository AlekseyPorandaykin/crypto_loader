package main

import (
	"github.com/AlekseyPorandaykin/crypto_loader/cmd"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/logger"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/metrics"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/shutdown"
	"go.uber.org/zap"
)

var version string

func main() {
	logger.InitDefaultLogger()
	defer func() { _ = zap.L().Sync() }()
	zap.L().Debug("Start app", zap.String("version", version))
	go func() {
		defer shutdown.HandlePanic()
		if err := metrics.Handler("localhost", "9081"); err != nil {
			zap.L().Fatal("error start metric", zap.Error(err))
		}
	}()
	cmd.Execute()
}
