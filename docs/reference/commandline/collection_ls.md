---
title: "hbm collection ls"
description: "The collection ls command description and usage"
keywords: [ "collection", "ls" ]
date: "2017-01-27"
menu:
  docs:
    parent: "hbm_cli_collection"
github_edit: "https://github.com/kassisol/hbm/edit/master/docs/reference/commandline/collection_ls.md"
---

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
