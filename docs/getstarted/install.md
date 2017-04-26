---
title: "Get started with HBM"
description: "Getting started with HBM"
tags: [ "getting started", "HBM" ]
date: "2017-01-27"
url: "/docs/getstarted/install/"
menu:
  main:
    identifier: hbm_getstarted_install
    parent: getstarted
    weight: -85
---

# Getting Started
## Install Harbormaster
The authorization plugin run as a host service.

### Manual
*  Download HBM (Harbormaster) binary ([link](https://github.com/kassisol/hbm/releases))
*  Copy binary to /usr/sbin/
```bash
# wget https://github.com/kassisol/hbm/releases/download/<version>/hbm -O /usr/sbin/hbm
```

### RPM Package
*  Download HBM (Harbormaster) rpm ([link](https://github.com/kassisol/hbm/releases))
*  Install HBM (Harbormaster) package
```bash
# yum localinstall hbm-<version>.el7.centos.x86_64.rpm
```

### Verifying the installation
After installing Harbormaster, verify the installation worked by opening a new terminal session as `root` and checking that `hbm` is available. By executing `hbm`, you should see help output similar to the following:

```bash
# hbm

HBM is an application to authorize and manage authorized docker commands

Usage:
hbm [command]

Available Commands:
  collection  Manage whitelisted collections
  group       Manage whitelisted groups
  info        Display information about HBM
  init        Initialize config
  policy      Manage policies
  resource    Manage whitelisted resources
  server      Launch the HBM server
  user        Manage whitelisted users
  version     Show the HBM version information

Flags:
  -h, --help   help for hbm

Use "hbm [command] --help" for more information about a command.

```

If you get an error that `hbm` could not be found, then your PATH environment variable was not setup properly. Please go back and ensure that your PATH variable contains the directory where Harbormaster was installed.

Otherwise, Harbormaster is installed and ready to go!


### Configuring Docker Engine

 * Update Docker daemon to run with authorization enabled.
     For example, if Docker is installed as a systemd service:
```bash
# mkdir /etc/systemd/system/docker.service.d
```

add authz-plugin parameter to ExecStart parameter
```bash
# vim /etc/systemd/system/docker.service.d/daemon.conf
[Service]
ExecStart=
ExecStart=/usr/bin/dockerd --authorization-plugin=hbm

# systemctl daemon-reload
# systemctl restart docker.service
```

## Starting the Server
### Starting the Harbormaster Server
With Harbormaster installed, the next step is to start a Harbormaster server.

```bash
# hbm init
# hbm server
INFO[0000] Server has completed initialization
INFO[0000] HBM server                           logdriver=standard storagedriver=sqlite version=0.2.0
INFO[0000] Listening on socket file
```
