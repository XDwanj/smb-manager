package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var (
	debug bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "smb-manager",
	Short: "A tool to manage SMB mounts",
	Long:  `smb-manager is a CLI tool that allows you to easily mount and unmount SMB shares.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// 权限检查逻辑
		return checkPermissions()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func checkPermissions() error {
	// 检查是否以root用户运行
	if os.Geteuid() != 0 {
		return fmt.Errorf("this tool requires root privileges, please run with sudo")
	}
	return nil
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug logging")
}

func initConfig() {
	// Initialize logger based on debug flag
	var level slog.Level
	if debug {
		level = slog.LevelDebug
	} else {
		level = slog.LevelInfo
	}

	// Set up logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))

	slog.SetDefault(logger)
}