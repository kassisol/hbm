---
description: Configuring and running the HBM server on various distributions
keywords:
  - hbm, server, configuration, running,  process managers
menu:
  main:
    parent: hbm_admin
    weight: 0
title: Configuring and running HBM
---

# Configuring and running HBM
***

After successfully installing HBM, the `hbm` server runs with its default
configuration.

In a production environment, system administrators typically configure the
`hbm` server to start and stop according to an organization's requirements. In most
cases, the system administrator configures a process manager such as `systemd`
to manage the `hbm` daemon's start and stop.

### Running the hbm server directly

The Harbormaster server can be run directly using the `hbm server` command. By default it listens on
the Unix socket `unix:///var/run/hbm.sock`

```bash
# hbm server
INFO[0000] Server has completed initialization
INFO[0000] HBM server        logdriver=standard storagedriver=sqlite version=0.2.0
INFO[0000] Listening on socket file

```

### Manually creating the systemd unit file

When installing the binary without a package, you may want
to integrate HBM with systemd. For this, simply install the service unit file
from [the github repository](https://github.com/kassisol/hbm/tree/master/contrib/init/systemd)
to `/etc/systemd/system`.

## CentOS / Red Hat Enterprise Linux

As of `7.x`, CentOS and RHEL use `systemd` as the process manager.

After successfully installing HBM for [CentOS](../installation/linux/centos.md) / [Red Hat Enterprise Linux](../installation/linux/rhel.md), you can check the running status in this way:

```bash
# systemctl status hbm
```

### Running HBM

You can start/stop/restart the `hbm` server using

```bash
# systemctl start hbm

# systemctl stop hbm

# systemctl restart hbm
```

If you want HBM to start at boot, you should also:

```bash
# systemctl enable hbm
```

### Logs

systemd has its own logging system called the journal. The logs for the `hbm` server can
be viewed using `journalctl -u hbm`

```bash
# journalctl -u hbm
```
