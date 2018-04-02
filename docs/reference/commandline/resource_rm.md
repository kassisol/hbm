---
title: "hbm resource rm"
description: "The resource rm command description and usage"
keywords: [ "resource", "rm" ]
date: "2017-01-27"
menu:
  docs:
    parent: "hbm_cli_resource"
github_edit: "https://github.com/kassisol/hbm/edit/master/docs/reference/commandline/resource_rm.md"
---

```markdown
Remove resource from the whitelist

Usage:
  hbm resource rm [flags]

Aliases:
  rm, remove
```

## Examples

```bash
# hbm resource ls
NAME                TYPE                VALUE               OPTION              COLLECTIONS
resource1           action              info                                    collection1
resource2           action              container_list                          collection2
# hbm resource rm resource1
# hbm resource ls
NAME                TYPE                VALUE               OPTION              COLLECTIONS
resource2           action              container_list                          collection2
```

## Related information

* [resource_add](resource_add.md)
* [resource_find](resource_find.md)
* [resource_ls](resource_ls.md)
* [resource_member](resource_member.md)
