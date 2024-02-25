package main

import (
	"github.com/AlekseyPorandaykin/crypto_loader/cmd"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/database"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/logger"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/monitoring"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/system"
	"go.uber.org/zap"
)

var version string

func main() {
	logger.InitDefaultLogger()
	defer func() { _ = zap.L().Sync() }()
	zap.L().Debug("Start app", zap.String("version", version))

	system.Go(func() {
		if err := monitoring.Handler("localhost", "9081"); err != nil {
			zap.L().Fatal("error start metric", zap.Error(err))
		}
	})
	database.Init("crypto_loader")
	cmd.Execute()
}
