---
description: The collection ls command description and usage
keywords:
- collection, ls
menu:
  main:
    parent: smn_cli
title: collection ls
---

# hbm collection ls
***

```markdown
List whitelisted collections

Usage:
  hbm collection ls [flags]

Aliases:
  ls, list

Flags:
  -f, --filter value   Filter output based on conditions provided (default [])
```

## Filtering

The filtering flag (`-f` or `--filter`) format is a `key=value` pair. If there is more
than one filter, then pass multiple flags (e.g. `--filter "foo=bar" --filter "bif=baz"`)

The currently supported filters are:

* name (collection's name)
* elem (resource)

## Examples

```bash
# hbm collection ls
NAME                RESOURCES
collection1         resource1
collection2         resource2
```

Filtering

```bash
# hbm cluster ls -f "elem=resource1"
NAME                RESOURCES
collection1         resource1
```

## Related information

* [collection_add](collection_add.md)
* [collection_find](collection_find.md)
* [collection_rm](collection_rm.md)
