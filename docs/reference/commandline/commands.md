---
title: "The HBM commands"
description: "HBM's CLI command description and usage"
keywords: [ "HBM", "documentation", "CLI", "command line" ]
date: "2017-01-27"
menu:
  docs:
    parent: "hbm_cli"
    weight: -200
github_edit: "https://github.com/kassisol/hbm/edit/master/docs/reference/commandline/commands.md"
toc: true
---

This section contains reference information on using HBM's command line
client. Each command has a reference page along with samples.

### Management commands

| Command | Description                                                                |
|:--------|:---------------------------------------------------------------------------|
| [info](info.md) | Display information about HBM                                      |
| [init](init.md) | Initialize config                                                  |
| [server](server.md) | Launch the HBM server                                          |
| [version](version.md) | Show the HBM version information                             |

### Config commands

| Command | Description                                                                |
|:--------|:---------------------------------------------------------------------------|
| [config get](config_get.md) | Get config option value                                |
| [config ls](config_ls.md) | List HBM configs                                         |
| [config set](config_set.md) | Set HBM config option                                  |

### User commands

| Command | Description                                                                |
|:--------|:---------------------------------------------------------------------------|
| [user add](user_add.md) | Add user to the whitelist                                  |
| [user find](user_find.md) | Verify if user exists in the whitelist                   |
| [user ls](user_ls.md) | List whitelisted users                                       |
| [user member](user_member.md) | Manage user membership to group                      |
| [user rm](user_rm.md) | Remove user from the whitelist                               |

### Group commands

| Command | Description                                                                |
|:--------|:---------------------------------------------------------------------------|
| [group add](group_add.md) | Add group to the whitelist                               |
| [group find](group_find.md) | Verify if group exists in the whitelist                |
| [group ls](group_ls.md) | List whitelisted groups                                    |
| [group rm](group_rm.md) | Remove group from the whitelist                            |

### Resource commands

| Command | Description                                                                |
|:--------|:---------------------------------------------------------------------------|
| [resource add](resource_add.md) | Add resource to the whitelist                      |
| [resource find](resource_find.md) | Verify if resource exists in the whitelist       |
| [resource ls](resource_ls.md) | List whitelisted resources                           |
| [resource member](resource_member.md) | Manage resource membership to collection     |
| [resource rm](resource_rm.md) | Remove resource from the whitelist                   |

### Collection commands

| Command | Description                                                                |
|:--------|:---------------------------------------------------------------------------|
| [collection add](collection_add.md) | Add collection to the whitelist                |
| [collection find](collection_find.md) | Verify if collection exists in the whitelist |
| [collection ls](collection_ls.md) | List whitelisted collections                     |
| [collection rm](collection_rm.md) | Remove collection from the whitelist             |

### Policy commands

| Command | Description                                                                |
|:--------|:---------------------------------------------------------------------------|
| [policy add](policy_add.md) | Add policy                                             |
| [policy find](policy_find.md) | Verify if policy exists                              |
| [policy ls](policy_ls.md) | List policies                                            |
| [policy rm](policy_rm.md) | Remove policy                                            |
