---
title: "hbm group ls"
description: "The group ls command description and usage"
keywords: [ "group", "ls" ]
date: "2017-01-27"
menu:
  main:
    parent: smn_cli
---

```markdown
List whitelisted groups

Usage:
  hbm group ls [flags]

Aliases:
  ls, list

Flags:
  -f, --filter value   Filter output based on conditions provided (default [])
```

## Filtering

The filtering flag (`-f` or `--filter`) format is a `key=value` pair. If there is more
than one filter, then pass multiple flags (e.g. `--filter "foo=bar" --filter "bif=baz"`)

The currently supported filters are:

* name (group's name)
* elem (user)

## Examples

```bash
# hbm group ls
NAME                USERS
group1              user1
group2              user2
```

Filtering

```bash
# hbm group ls -f "elem=user1"
NAME                USERS
group1              user1
```

## Related information

* [group_add](group_add.md)
* [group_find](group_find.md)
* [group_rm](group_rm.md)
