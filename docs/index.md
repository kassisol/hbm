---
description: HBM
keywords:
  - hbm
menu:
  main:
    identifier: hbm_use
    weight: -85
title: HBM
---

# About HBM

**Secure, restrict Docker's capabilities**

[**HBM**](http://harbormaster.io) is a basic extendable
Docker Engine [access authorization plugin]
(https://docs.docker.com/engine/extend/plugins_authorization/)
that runs on directly on the Docker host.

## Why HBM?

By default Harbormaster plugin prevents from executing commands and certains parameters.

*Commands*

*Pull images*

*Start containers*

* `--privileged`
* `--ipc=host`
* `--net=host`
* `--pid=host`
* `--userns=host`
* `--uts=host`
* any Linux capabilities with parameter `--cap-add=[]`
* any devices added with parameter `--device=[]`
* any dns servers added with parameter `--dns=`
* any ports added with parameter `--port=`
* any volumes mounted with parameter `-v`
* any logging with parameters `--log-driver` and `--log-opt`

## About this guide

### Installation guides

The [installation section](installation/index.md) will show you how to install HBM
on a variety of platforms.


### HBM admin guide

To learn about HBM in more detail and to answer questions about usage and
implementation, check out the [Docker Admin Guide](admin/index.md).

## Release notes

A summary of the changes in each release in the current series can now be found
on the separate [Release Notes page](http://harbormaster.io/docs/release-notes)

## Licensing

HBM is licensed under the GNU GPL 3. See
[LICENSE](https://github.com/kassisol/hbm/blob/master/LICENSE) for the full
license text.
