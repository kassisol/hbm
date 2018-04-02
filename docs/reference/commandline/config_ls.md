---
title: "hbm config ls"
description: "The config ls command description and usage"
keywords: [ "config", "ls" ]
date: "2017-01-27"
menu:
  docs:
    parent: "hbm_cli_config"
github_edit: "https://github.com/kassisol/hbm/edit/master/docs/reference/commandline/config_ls.md"
---

```markdown
List HBM configs

Usage:
  hbm config ls [flags]

Aliases:
ls, list

Flags:
  -f, --filter strings   Filter output based on conditions provided
  -h, --help             help for ls
```

## Examples

```bash
# hbm config ls
NAME                        VALUE
authorization               false
default-allow-action-error  false
```

## Related information

* [config_get](config_get.md)
* [config_set](config_set.md)
