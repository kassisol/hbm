# HBM (Harbourmaster)
Harbourmaster is a basic extendable Docker Engine [access authorization plugin] (https://docs.docker.com/engine/extend/plugins_authorization/) that runs on directly on the host.

By default Harbourmaster plugin prevents from executing commands with certains parameters.
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

# Getting Started
## Install Harbourmaster
The authorization plugin run as a host service.

### Manual
*  Download HBM (Harbourmaster) binary (link)
*  Install HBM (Harbourmaster) as a service
```bash
# wget xxx -O /usr/local/bin/hbm
```

### RPM Package
*  Download HBM (Harbourmaster) rpm (link)
*  Install HBM (Harbourmaster) package
```bash
# yum localinstall hbm-0.1.0-x86_64-el7.rpm
```

### Verifying the installation
After installing `Harbourmaster`, verify the installation worked by opening a new terminal session as `root` and checking that `hbm` is available. By executing `hbm`, you should see help output similar to the following:

```bash
# hbm
HBM is a command line to restrict docker use

Usage:
  hbm [command]

Available Commands:
  action      Manage whitelisted actions
  cap         Manage whitelisted caps
  config      Manage whitelisted configs
  device      Manage whitelisted devices
  dns         Manage whitelisted DNS server
  image       Manage whitelisted images
  info        Display information about Harbourmaster
  init        Initialize config
  port        Manage whitelisted ports
  registry    Manage whitelisted registries
  server      Starts a Docker AuthZ server
  version     Print the version number of Harbourmaster
  volume      Manage whitelisted volumes

Flags:
  -h, --help   help for hbm

Use "hbm [command] --help" for more information about a command.
```

If you get an error that `hbm` could not be found, then your PATH environment variable was not setup properly. Please go back and ensure that your PATH variable contains the directory where Harbourmaster was installed.

Otherwise, Harbourmaster is installed and ready to go!

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
ExecStart=/usr/bin/docker daemon -H fd:// --authorization-plugin=hbm

# systemctl daemon-reload
# systemctl restart docker.service
```

## Starting the Server
### Starting the Harbourmaster Server
With Harbourmaster installed, the next step is to start a Harbourmaster server.

```bash
# hbm init
# hbm server
2016/05/18 16:14:43 Listening on socket file
```
