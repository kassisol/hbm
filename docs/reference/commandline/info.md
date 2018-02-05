---
description: The info command description and usage
keywords:
- display, hbm, information
menu:
  main:
    parent: smn_cli
title: info
---

# hbm info
***

```markdown
Display information about HBM

Usage:
  hbm info [flags]

Flags:
  -h, --help   help for info
```

This command displays system wide information regarding the HBM installation.
Information displayed includes the server version, number of policies, groups and collections.

# Examples

## Display HBM information

```bash
# hbm info
Features Enabled:
 Authorization: false
 Default Allow Action On Error: false
Policies: 0
Groups: 1
 Users: 0
Collections: 0
 Resources: 0
  Actions: 0
  Configs: 0
  Capabilities: 0
  Devices: 0
  DNS Servers: 0
  Images: 0
  Ports: 0
  Registries: 0
  Volumes: 0
Server Version: 0.8.0
Storage Driver: sqlite
Logging Driver: standard
Harbormaster Root Dir: /var/lib/hbm
Docker AuthZ Plugin Enabled: true
```
