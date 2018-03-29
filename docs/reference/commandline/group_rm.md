---
title: "hbm group rm"
description: "The group rm command description and usage"
keywords: [ "group", "rm" ]
date: "2017-01-27"
menu:
  main:
    parent: smn_cli
---

```markdown
Remove group from the whitelist

Usage:
  hbm group rm [name] [flags]

Aliases:
  rm, remove
```

## Examples

```bash
# hbm group ls
NAME                USERS
group1              user1
group2              user2
# hbm group rm group1
# hbm group ls
NAME                USERS
group2              user2
```

## Related information

* [group_add](group_add.md)
* [group_find](group_find.md)
* [group_ls](group_ls.md)
