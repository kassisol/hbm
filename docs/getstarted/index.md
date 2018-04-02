---
title: "Get started with HBM"
description: "Getting started with HBM"
keywords: [ "getting started", "HBM" ]
date: "2017-01-27"
url: "/docs/getstarted/hbm/"
menu:
  docs:
    parent: "getstarted"
    weight: -85
github_edit: "https://github.com/kassisol/hbm/edit/master/docs/getstarted/index.md"
---

## Install

Refer to the [page](../installation/index.md) specific to your Linux distribution.

## Configure

By default, all Docker commands and restricted parameters are allowed. To change that behavior, an option needs to be set to true, then all commands will be blocked and so need to be white-listed.

```bash
hbm config set authorization true
```

## Add a policy

To add a policy you need to have 2 elements, a group and a collection.

First create a group and an user.

> If Docker Daemon is listening on Unix socket, the only user will be root.

```bash
hbm group add local
hbm user add root
hbm user member --add local root
```

Then create a collection to which resources will be assigned to. Resources could be anything from Docker commands to images, volumes, restricted parameters like `--privileged`; `--net=host` and so on... A list of types and values can be found on that [page](../reference/commandline/resource_add.md).

```bash
hbm collection add collection1
hbm resource add --type action --value info info
hbm resource add --type action --value version version
hbm resource member --add collection1 info
hbm resource member --add collection1 version
```

To finish create the policy.

```bash
hbm policy add --group local --collection collection1 policy1
```
