# SMB Manager

A command-line tool to manage SMB (Server Message Block) shares mounting and unmounting.

## Features

- Mount SMB shares defined in a configuration file
- Unmount SMB shares defined in a configuration file
- Support for default mount options with per-share overrides
- Logging with debug mode
- Built with Go using Cobra and Slog

## Installation

```bash
go install
```

## Configuration

The tool expects a configuration file at `/etc/samba/mounts.yaml` with the following structure:

```yaml
mounts:
  default-options: # Default mount options
    # <filesystem type>
    type: cifs
    # <mount options>
    options: nofail,iocharset=utf8,x-gvfs-show,x-systemd.automount,credentials=/etc/samba/credentials/share,uid=1000,gid=1000
  points:
  - src: //server/share
    dst: /mnt/share
    type: cifs
  - src: //server/documents
    dst: /home/user/Documents
    type: cifs
```

## Usage

### Mount SMB shares

```bash
smb-manager mount
```

### Unmount SMB shares

```bash
smb-manager umount
```

### Enable debug logging

```bash
smb-manager --debug mount
```

## Implementation Details

The tool uses the standard `mount` and `umount` commands under the hood, so it needs to be run with appropriate permissions (typically as root or with sudo).

## Directory Structure

```
smb-manager/
├── cmd/
│   ├── root.go
│   ├── mount.go
│   └── umount.go
├── config/
│   └── config.go
├── internal/
│   └── mounter/
│       └── mounter.go
├── main.go
├── go.mod
└── README.md
```