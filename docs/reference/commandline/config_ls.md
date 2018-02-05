---
description: The config ls command description and usage
keywords:
- config, ls
menu:
  main:
    parent: smn_cli
title: config ls
---

# hbm config ls
***

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
