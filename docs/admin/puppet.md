---
title: "Using Puppet"
description: "Installing and using Puppet"
keywords: [ "puppet", "installation", "usage", "hbm", "documentation" ]
date: "2017-01-27"
url: "/docs/hbm/admin/puppet/"
menu:
  docs:
    parent: "hbm_admin"
    weight: 10
github_edit: "https://github.com/kassisol/hbm/edit/master/docs/admin/puppet.md"
toc: true
---

## Requirements

To use this guide you'll need a working installation of Puppet from
[Puppet Labs](https://puppetlabs.com) .

## Installation

The module is available on the [GitHub](https://github.com/kassisol/puppet-module-hbm).

## Usage

The module provides a puppet class for installing HBM and types
for managing groups, users, collections, resources and policies.

```
classes:
  - 'hbm'
```

### Configs

#### Enable

```
hbm::configs:
  'authorization:
    ensure: 'present'
```

This is equivalent to running:

```bash
# hbm config set authorization true
```

#### Disable

```
hbm::configs:
  'authorization:
    ensure: 'absent'
```

This is equivalent to running:

```bash
# hbm config set authorization false
```

### Groups

#### Add

```
hbm::groups:
  'group1:
    ensure: 'present'
```

This is equivalent to running:

```bash
# hbm group add group1
```

#### Remove

```
hbm::groups:
  'group1:
    ensure: 'absent'
```

This is equivalent to running:

```bash
# hbm group rm group1
```

### Users

#### Add

```
hbm::users:
  'user1:
    ensure: 'present'
```

This is equivalent to running:

```bash
# hbm user add user1
```

#### Remove

```
hbm::users:
  'user1:
    ensure: 'absent'
```

This is equivalent to running:

```bash
# hbm user rm user1
```

#### Member

```
hbm::users:
  'user1:
    ensure: 'present'
    members:
      - group1
      - group2
```

This is equivalent to running:

```bash
# hbm user member --add group1 user1
# hbm user member --add group2 user1
```

### Collections

#### Add

```
hbm::collections:
  'collection1:
    ensure: 'present'
```

This is equivalent to running:

```bash
# hbm collection add collection1
```

#### Remove

```
hbm::collections:
  'collection1:
    ensure: 'absent'
```

This is equivalent to running:

```bash
# hbm collection rm collection1
```

### Resources

The next step is probably to configure HBM resources. For this, we have a
defined type which can be used like so:

#### Add

```
hbm::resources:
  'resource1:
    ensure: 'present'
    type: 'action'
    value: 'container_list'
```

This is equivalent to running:

```bash
# hbm resource add --type action --value container_list resource1
```

#### Remove

```
hbm::resources:
  'resource1:
    ensure: 'absent'
```

This is equivalent to running:

```bash
# hbm resource rm resource1
```

#### Member

```
hbm::resources:
  'resource1:
    ensure: 'present'
    members:
      - collection1
      - collection2
```

This is equivalent to running:

```bash
# hbm resource member --add collection1 resource1
# hbm resource member --add collection2 resource1
```

### Policies

Now you have an image where you can run commands within a container
managed by Docker.

#### Add

```
hbm::policies:
  'policy1':
    ensure: 'present'
    collection: 'collection1'
    group: 'group1'
```

This is equivalent to running the following command:

```bash
# hbm policy add --collection collection1 --group group1 policy1
```

#### Remove

```
hbm::policies:
  'policy1':
    ensure: 'absent'
```

This is equivalent to running the following command:

```bash
# hbm policy rm policy1
```
