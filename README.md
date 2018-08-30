# HBM (Harbormaster)

[![Build Status](https://travis-ci.org/kassisol/hbm.svg?branch=master)](https://travis-ci.org/kassisol/hbm)
[![Go Report Card](https://goreportcard.com/badge/github.com/kassisol/hbm)](https://goreportcard.com/report/github.com/kassisol/hbm)
[![MicroBadger](https://images.microbadger.com/badges/version/kassisol/hbm:0.11.0.svg)](https://microbadger.com/images/kassisol/hbm:0.11.0 "Get your own version badge on microbadger.com")

Harbormaster is a basic extendable Docker Engine [access authorization plugin](https://docs.docker.com/engine/extend/plugins_authorization/) that runs on directly on the host.

By default Harbormaster plugin prevents from executing commands with certains parameters.
 1. Pull images
 2. Start containers
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
  * any logging with parameters "--log-driver" and "--log-opt"

## Versions

Supported Docker versions with HBM.

| HBM Version | Docker Version | Docker API |
|-------------|----------------|------------|
| 0.2.x       | 1.12.x         | 1.24       |
| 0.3.x       | 17.05.x        | 1.29       |
| 0.5.x       | 17.06.x        | 1.30       |
| 0.5.x       | 17.09.x        | 1.32       |
| >= 0.6.0    | >= 1.12.x      | >= 1.24    |

## Getting Started & Documentation

All documentation is available on the [Harbormaster website](http://harbormaster.io/docs/hbm/).

## User Feedback

### Issues

If you have any problems with or questions about this application, please contact us through a [GitHub](https://github.com/kassisol/hbm/issues) issue.
