package system

import (
	"go.uber.org/zap"
	"os"
	"runtime/debug"
	"time"
)

func Go(run func()) {
	defer HandlePanic()
	go run()
}

func HandlePanic() {
	if err := recover(); err != nil {
		zap.L().Error(
			"handle panic",
			zap.Any("recover", err),
			zap.ByteString("stack", debug.Stack()),
		)
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
}
