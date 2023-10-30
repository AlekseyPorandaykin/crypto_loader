package cmd

import (
	"context"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/repositories"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/server/http"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/storage"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os/signal"
	"syscall"
)

var httpServerCmd = &cobra.Command{
	Use:   "http",
	Short: "Run http server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer cancel()
		//DB
		db, err := repositories.CreateDB(repositories.Config{
			Driver:   "postgres",
			Username: "crypto_loader",
			Password: "developer",
			Host:     "localhost",
			Port:     "5433",
			Database: "crypto_loader",
		})
		if err != nil {
			fmt.Println("Error init database: ", err.Error())
			return
		}
		defer func() { _ = db.Close() }()
		//Repository
		priceRepo := repositories.NewPriceRepository(db)
		//Storage
		priceStorage := storage.NewPriceStorage(priceRepo, cacheDir)
		serv := http.NewServer(":8080", priceStorage)
		defer serv.Close()
		go func() {
			if err := serv.Run(); err != nil && !errors.Is(err, context.Canceled) {
				zap.L().Error("failed start serve", zap.Error(err))
			}
		}()
		<-ctx.Done()
	},
}

func init() {
	ServerCmd.AddCommand(httpServerCmd)
}
