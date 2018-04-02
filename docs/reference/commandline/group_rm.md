---
title: "hbm group rm"
description: "The group rm command description and usage"
keywords: [ "group", "rm" ]
date: "2017-01-27"
menu:
  docs:
    parent: "hbm_cli_group"
github_edit: "https://github.com/kassisol/hbm/edit/master/docs/reference/commandline/group_rm.md"
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
