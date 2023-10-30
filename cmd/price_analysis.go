package cmd

import (
	"context"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/analasis"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/repositories"
	"github.com/spf13/cobra"
	"os/signal"
	"syscall"
	"time"
)

var priceAnalysis = &cobra.Command{
	Use: "price-analysis",
	Run: func(cmd *cobra.Command, args []string) {
		const DefaultRecalculateDuration = time.Hour
		ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer cancel()
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
		priceRepo := repositories.NewPriceRepository(db)
		priceChangesRepo := repositories.NewPriceChanges(db)
		ap := analasis.NewPrice(priceRepo, priceChangesRepo)
		go func() {
			defer cancel()
			if err := ap.Run(ctx, DefaultRecalculateDuration); err != nil {
				fmt.Printf("error execute app: %s \n", err.Error())
			}
		}()

		<-ctx.Done()
	},
}

func init() {
	AnalysisCmd.AddCommand(priceAnalysis)
}
