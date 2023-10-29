package cmd

import (
	"context"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/repositories"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/server/grpc"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/storage"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os/signal"
	"syscall"
)

var grpcServerCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Run grpc server",
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
		serv := grpc.NewServer(priceStorage, ":50052")
		defer serv.Close()
		go func() {
			if err := serv.Start(); err != nil && !errors.Is(err, context.Canceled) {
				zap.L().Error("failed start serve", zap.Error(err))
			}
		}()
		<-ctx.Done()
	},
}

func init() {
	ServerCmd.AddCommand(grpcServerCmd)
}
