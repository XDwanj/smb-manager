package cmd

import (
	"log/slog"
	"smb-manager/config"
	"smb-manager/internal/mounter"

	"github.com/spf13/cobra"
)

// mountCmd represents the mount command
var mountCmd = &cobra.Command{
	Use:   "mount",
	Short: "Mount SMB shares",
	Long:  `Mount SMB shares based on the configuration in the mounts.yaml file.`,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("Mounting SMB shares...")
		
		// Load configuration
		configPath, err := config.GetConfigPath()
		if err != nil {
			slog.Error("Failed to get config path", "error", err)
			return
		}
		
		cfg, err := config.LoadConfig(configPath)
		if err != nil {
			slog.Error("Failed to load config", "error", err)
			return
		}
		
		slog.Info("Loaded config", "points", len(cfg.Mounts.Points))
		
		// Mount all shares
		err = mounter.MountAll(cfg)
		if err != nil {
			slog.Error("Failed to mount some shares", "error", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(mountCmd)
}