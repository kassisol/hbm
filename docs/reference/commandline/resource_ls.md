---
title: "hbm resource ls"
description: "The resource ls command description and usage"
keywords: [ "resource", "ls" ]
date: "2017-01-27"
menu:
  docs:
    parent: "hbm_cli_resource"
github_edit: "https://github.com/kassisol/hbm/edit/master/docs/reference/commandline/resource_ls.md"
---

```markdown
List whitelisted resources

Usage:
  hbm resource ls [flags]

Aliases:
  ls, list

Flags:
  -f, --filter value   Filter output based on conditions provided (default [])
```

## Filtering

The filtering flag (`-f` or `--filter`) format is a `key=value` pair. If there is more
than one filter, then pass multiple flags (e.g. `--filter "foo=bar" --filter "bif=baz"`)

The currently supported filters are:

* name (resource's name)
* type
* value
* elem (collection)

## Examples

```bash
# hbm resource ls
NAME                TYPE                VALUE               OPTION              COLLECTIONS
resource1           action              info                                    collection1
resource2           action              container_list                          collection2
```

### Filtering

```bash
# hbm resource ls -f "elem=collection1"
NAME                TYPE                VALUE               OPTION              COLLECTIONS
resource1           action              info                                    collection1
```

## Related information

* [resource_add](resource_add.md)
* [resource_find](resource_find.md)
* [resource_member](resource_member.md)
* [resource_rm](resource_rm.md)
