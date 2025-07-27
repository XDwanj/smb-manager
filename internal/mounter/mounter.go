package mounter

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"smb-manager/config"
)

// Mount mounts an SMB share based on the provided configuration
func Mount(mount config.Mount) error {
	slog.Debug("Mounting SMB share", "src", mount.Src, "dst", mount.Dst, "type", mount.Type, "options", mount.Option)

	// Check if the destination is already mounted
	mountpointCmd := exec.Command("mountpoint", "-q", mount.Dst)
	if err := mountpointCmd.Run(); err == nil {
		slog.Info("SMB share already mounted", "dst", mount.Dst)
		return nil
	}

	// Ensure destination directory exists
	err := os.MkdirAll(mount.Dst, 0755)
	if err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	// Build mount command
	cmd := exec.Command("mount", "-t", mount.Type, "-o", mount.Option, mount.Src, mount.Dst)

	// Execute mount command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to mount SMB share: %w, output: %s", err, output)
	}

	slog.Info("Successfully mounted SMB share", "src", mount.Src, "dst", mount.Dst)
	return nil
}

// Umount unmounts an SMB share at the specified destination
func Umount(dst string) error {
	slog.Debug("Unmounting SMB share", "dst", dst)

	// Check if the destination is mounted
	mountpointCmd := exec.Command("mountpoint", "-q", dst)
	if err := mountpointCmd.Run(); err != nil {
		slog.Info("SMB share is not mounted, skipping unmount", "dst", dst)
		return nil
	}

	// Build umount command
	cmd := exec.Command("umount", dst)

	// Execute umount command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to unmount SMB share: %w, output: %s", err, output)
	}

	slog.Info("Successfully unmounted SMB share", "dst", dst)
	return nil
}

// MountAll mounts all SMB shares defined in the configuration
func MountAll(cfg *config.Config) error {
	slog.Info("Mounting all SMB shares...")

	for _, mount := range cfg.Mounts.Points {
		err := Mount(mount)
		if err != nil {
			slog.Error("Failed to mount SMB share", "src", mount.Src, "dst", mount.Dst, "error", err)
			// Continue with other mounts even if one fails
		}
	}

	return nil
}

// UmountAll unmounts all SMB shares defined in the configuration
func UmountAll(cfg *config.Config) error {
	slog.Info("Unmounting all SMB shares...")

	// Unmount in reverse order
	for i := len(cfg.Mounts.Points) - 1; i >= 0; i-- {
		mount := cfg.Mounts.Points[i]
		err := Umount(mount.Dst)
		if err != nil {
			slog.Error("Failed to unmount SMB share", "dst", mount.Dst, "error", err)
			// Continue with other unmounts even if one fails
		}
	}

	return nil
}