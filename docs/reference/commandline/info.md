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
```

This command displays system wide information regarding the HBM installation.
Information displayed includes the server version, number of policies, groups, clusters and collections.

# Examples

## Display HBM information

Here is a sample output for a daemon running on Ubuntu, using the overlay
storage driver:

```
# hbm info
Policies: 0
Groups: 0
 Users: 0
Clusters: 0
 Hosts: 0
Collections: 0
 Resources: 0
  Actions: 0
  Config: 0
  Capabilities: 0
  Devices: 0
  DNS Servers: 0
  Images: 0
  Ports: 0
  Registries: 0
  Volumes: 0
Server Version: 0.2.0
Storage Driver: sqlite
Logging Driver: standard
Docker AuthZ Plugin Enabled: false
Harbormaster Root Dir: /var/lib/hbm
```
