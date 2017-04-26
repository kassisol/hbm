---
description: The user rm command description and usage
keywords:
- user, rm
menu:
  main:
    parent: smn_cli
title: user rm
---

# hbm user rm
***

```markdown
Remove user from the whitelist

Usage:
  hbm user rm [flags]

Aliases:
  rm, remove
```

## Examples

```bash
# hbm user ls
NAME                GROUPS
user1               group1
user2               group2
# hbm user rm user1
# hbm user ls
NAME                GROUPS
user2               group2
```

## Related information

* [user_add](user_add.md)
* [user_find](user_find.md)
* [user_ls](user_ls.md)
* [user_member](user_member.md)
