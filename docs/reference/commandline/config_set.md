---
title: "hbm config set"
description: "The config set command description and usage"
keywords: [ "config", "set" ]
date: "2017-01-27"
menu:
  docs:
    parent: "hbm_cli_config"
github_edit: "https://github.com/kassisol/hbm/edit/master/docs/reference/commandline/config_set.md"
---

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
