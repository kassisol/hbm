---
description: The config set command description and usage
keywords:
- config, set
menu:
  main:
    parent: smn_cli
title: config set
---

# hbm config set
***

```markdown
Set HBM config option

Usage:
  hbm config set [key] [value] [flags]

Flags:
  -h, --help   help for set
```

## Examples

```bash
# hbm config set authorization true
# hbm config ls
NAME                        VALUE
authorization               true
default-allow-action-error  false
```

## Related information

* [config_get](config_get.md)
* [config_ls](config_ls.md)
