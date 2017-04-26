---
description: The user ls command description and usage
keywords:
- user, ls
menu:
  main:
    parent: smn_cli
title: user ls
---

# hbm user ls
***

```markdown
List whitelisted users

Usage:
  hbm user ls [flags]

Aliases:
  ls, list

Flags:
  -f, --filter value   Filter output based on conditions provided (default [])
```

## Filtering

The filtering flag (`-f` or `--filter`) format is a `key=value` pair. If there is more
than one filter, then pass multiple flags (e.g. `--filter "foo=bar" --filter "bif=baz"`)

The currently supported filters are:

* name (user's name)
* elem (username)

## Examples

```bash
# hbm user ls
NAME                GROUPS
user1               group1
user2               group2
```

Filtering

```bash
# hbm user ls -f "elem=group1"
NAME                GROUPS
user1               group1
```

## Related information

* [user_add](user_add.md)
* [user_find](user_find.md)
* [user_member](user_member.md)
* [user_rm](user_rm.md)
