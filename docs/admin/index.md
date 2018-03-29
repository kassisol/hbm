---
title: "Configuring and running HBM"
description: "Configuring and running the HBM server on various distributions"
keywords: [ "hbm", "server", "configuration", "running", "process managers" ]
date: "2017-01-27"
menu:
  main:
    parent: hbm_admin
    weight: 0
---

## Configuration

After successfully installing HBM, next is to initialize and start the HBM backend database. This is done with just a few commands. Please refer to the [command line reference](../reference/commandline/init.md).

`hbm` server runs with its default configuration and log every Docker commands.

> The logging module cannot be disabled.

There are 2 options available:

* `authorization`
* `default-allow-action-error`

See the [command line reference](../reference/commandline/config_set.md) to manage configuration.

### Authorization

That feature enable / disable verification of restricted Docker commands.

### Default Allow Action Error

On bug, that feature allows the Docker command to be authorized instead of panicing and blocking execution of commands.

## Policy

A policy determines which resources a user may run on specified hosts.
The basic structure of a policy is constituated of group and collection. Let's break that down into its constituent parts:

```
hbm policy add --group <group_name> --collection <collection_name> <policy_name>
```

### Group spec

A group regroup users.

If Docker Daemon is listening on Unix socket, the only user will be `root`. The only way for authz plugin to be aware of multiple users other than `root`, is to configure Docker Daemon with TLS enabled.

A special group `administrators` is created at initialization to allow members of that group to by pass authorization verification.

> `administrators` group cannot be deleted.

### Collection spec

A collection regroup resources that can be of the following types:

* action
* cap
* config
* device
* dns
* image
* logdriver
* logopt
* port
* registry
* volume

For details about the resources' values, refer to the [command line reference](../reference/commandline/resource_add.md).

It's possible to use the keyword `all` for users and resources type or value.

| User     | Resource Type | Resource Value | Resource Options |
|:---------|:--------------|:---------------|:-----------------|
| all      | all           | all            | all              |
| all      | all           | all            | `<empty>`        |
| all      | `<rType>`     | all            | all              |
| all      | `<rType>`     | all            | `<empty>`        |
| all      | `<rType>`     | `<rValue>`     | `<rOptions>`     |
| `<user>` | `<rType>`     | all            | all              |
| `<user>` | `<rType>`     | all            | `<empty>`        |
| `<user>` | `<rType>`     | `<rValue>`     | `<rOptions>`     |
