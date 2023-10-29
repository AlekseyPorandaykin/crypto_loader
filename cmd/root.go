package cmd

import (
	"context"
	"errors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"path"
)

var homeDir string = "/Users/alexey.porandaikin/Projects/go/projects/crypto_loader"

var cacheDir string = path.Join(homeDir, "storage/cache")

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Servers application",
}
var LoaderCmd = &cobra.Command{
	Use:   "loader",
	Short: "Scripts for load data from external sources",
}
var AnalysisCmd = &cobra.Command{
	Use:   "analysis",
	Short: "Scripts for analysis information",
}

var rootCmd = &cobra.Command{Use: "crypto-loader"}

func init() {
	rootCmd.AddCommand(ServerCmd, LoaderCmd, AnalysisCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil && !errors.Is(err, context.Canceled) {
		zap.L().Error("execute root cmd", zap.Error(err))
	}
}
