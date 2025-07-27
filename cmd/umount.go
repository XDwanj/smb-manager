package cmd

import (
	"log/slog"
	"smb-manager/config"
	"smb-manager/internal/mounter"

	"github.com/spf13/cobra"
)

// umountCmd represents the umount command
var umountCmd = &cobra.Command{
	Use:   "umount",
	Short: "Unmount SMB shares",
	Long:  `Unmount SMB shares based on the configuration in the mounts.yaml file.`,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("Unmounting SMB shares...")
		
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
		
		// Unmount all shares
		err = mounter.UmountAll(cfg)
		if err != nil {
			slog.Error("Failed to unmount some shares", "error", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(umountCmd)
}