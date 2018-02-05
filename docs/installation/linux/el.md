---
description: Instructions for installing Docker on CentOS.
keywords:
  - HBM, HBM documentation, requirements, linux, centos
menu:
  main:
    parent: hbm_linux
    weight: -4
title: Installation on Enterprise Linux
---

# EL
***

HBM runs on EL. An installation on other binary compatible EL7
distributions such as CentOS, RHEL, Scientific Linux might succeed, but Harbormaster does not test
or support HBM on these distributions.

These instructions install HBM using release packages and installation
mechanisms managed by Harbormaster.

## Prerequisites

HBM requires a 64-bit OS and version 1.12.x of the Docker Engine.

To check your current Docker version, open a terminal and use `docker version -f '{{ .Server.Version }}'` to
display your Docker version:

```bash
# docker version -f '{{ .Server.Version }}'
1.12.1
```

Finally, it is recommended that you fully update your system. Keep in mind
that your system should be fully patched to fix any potential kernel bugs.
Any reported kernel bugs may have already been fixed on the latest kernel
packages.

## Install HBM

### Install with yum

1. Log into your machine as a user with `sudo` or `root` privileges.

2. Make sure your existing packages are up-to-date.

    ```bash
    # yum update
    ```

3. Add the `yum` repo.

    ```bash
    # cd /usr/local/src
    # wget 
    ```

4. Install the HBM package.

    ```bash
    # yum localinstall hbm-x.x.x-x86_64-el7.rpm
    ```

5. Verify `hbm` is installed correctly by executing `hbm`, you should see help output similar to the following:

```
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
server      Starts a Docker AuthZ server
user        Manage whitelisted users
version     Show the HBM version information

Flags:
-h, --help   help for hbm

Use "hbm [command] --help" for more information about a command.
```

## Start the hbm server at boot

Configure the HBM server to start automatically when the host starts:

```bash
# systemctl enable hbm.service
```

## Uninstall

You can uninstall the HBM software with `yum`.

1. List the installed HBM package.

    ```bash
    # yum list installed | grep hbm

    hbm.x86_64     x.x.x-1.el7.centos
    ```

2. Remove the package.

    ```bash
    # yum -y remove hbm.x86_64
    ```

	This command does not remove data files on your host.

3. To delete all data, run the following command:

    ```bash
    # rm -rf /var/lib/hbm
    ```
