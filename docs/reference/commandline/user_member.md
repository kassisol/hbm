---
title: "hbm user member"
description: "The user member command description and usage"
keywords: [ "user", "member" ]
date: "2017-01-27"
menu:
  docs:
    parent: "hbm_cli_user"
github_edit: "https://github.com/kassisol/hbm/edit/master/docs/reference/commandline/user_member.md"
---

```markdown
Manage user membership to group

Usage:
  hbm user member [group] [name] [flags]

Flags:
  -a, --add      Add user to group
  -r, --remove   Remove user from group
```

## Examples

### Add a user to a group
```bash
# hbm user ls
NAME                GROUPS
user1
# hbm user member --add group1 user1
# hbm user ls
NAME                GROUPS
user1               group1
```

### Remove a user from a group
```bash
# hbm user ls
NAME                GROUPS
user1               group1
# hbm user member --remove group1 user1
# hbm user ls
NAME                GROUPS
user1
```

## Related information

* [user_add](user_add.md)
* [user_find](user_find.md)
* [user_ls](user_ls.md)
* [user_rm](user_rm.md)
