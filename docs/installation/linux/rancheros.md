---
title: "Installation on RancherOS"
linktitle: "RancherOS"
description: "Instructions for installing HBM on RancherOS"
keywords: [ "HBM", "documentation", "requirements", "linux", "rancheros" ]
date: "2018-03-31"
url: "/docs/hbm/install/linux/rancheros/"
menu:
  docs:
    parent: "hbm_linux"
    weight: -4
github_edit: "https://github.com/kassisol/hbm/edit/master/docs/install/linux/rancheros.md"
toc: true
---

## Install HBM

```bash
# mkdir -p /opt/hbm/data
```

```yaml
#cloud-config
rancher:
  docker:
    extra_args: [ '--authorization-plugin=hbm' ]

  services:
    hbm:
      image: kassisol/hbm:0.9.4
      labels:
        io.rancher.os.scope: "system"
        io.rancher.os.after: "network"
        io.rancher.os.before: "docker"
      restart: always
      volumes:
        - /etc/docker:/etc/docker
        - /var/run/docker:/var/run/docker
        - /opt/hbm/data:/var/lib/hbm
```

## Start the hbm server

After editing the user cloud config file, the server can be rebooted

```bash
# reboot
```

or start the services manually:

```bash
ros service stop docker
ros service rm docker
ros service up hbm
ros service up docker
```

## Execute hbm commands

To run hbm commands

```bash
# system-docker exec -t hbm hbm info
```

## Uninstall

Remove config from cloud-config.yml

```bash
ros service stop docker
ros service rm docker
ros service stop hbm
ros service rm hbm
ros service up docker
```
