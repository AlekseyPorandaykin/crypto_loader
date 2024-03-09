package cmd

import (
	"context"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/config"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/database"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Execute migrate",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer cancel()
		conf := config.Create()
		//DB
		db, err := database.CreateConnection(conf.ConfDB)
		if err != nil {
			fmt.Println("Error init database: ", err.Error())
			return
		}
		defer func() { _ = db.Close() }()
		dirMigration := filepath.Join(os.Getenv("APP_DIR"), "migrations", conf.ConfDB.Driver)
		dirs, err := os.ReadDir(dirMigration)
		if err != nil {
			zap.L().Error("read dir", zap.Error(err))
			return
		}
		for _, dir := range dirs {
			if dir.IsDir() {
				continue
			}
			executeSqlFile(ctx, filepath.Join(dirMigration, dir.Name()), db)
		}

	},
}

func executeSqlFile(ctx context.Context, path string, db *sqlx.DB) {
	f, errF := os.Open(path)
	if errF != nil {
		zap.L().Error("open file", zap.Error(errF))
		return
	}
	defer f.Close()
	data, errR := io.ReadAll(f)
	if errR != nil {
		zap.L().Error("read data file", zap.Error(errR))
		return
	}
	_, errExec := db.ExecContext(ctx, string(data))
	if errExec != nil {
		zap.L().Error("error execute sql", zap.Error(errExec))
		return
	}
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
