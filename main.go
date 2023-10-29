package main

import (
	"github.com/AlekseyPorandaykin/crypto_loader/cmd"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/logger"
	"go.uber.org/zap"
)

var version string

func main() {
	logger.InitDefaultLogger()
	defer func() { _ = zap.L().Sync() }()
	zap.L().Debug("Start app", zap.String("version", version))
	cmd.Execute()
}
