---
title: "hbm policy rm"
description: "The policy rm command description and usage"
keywords: [ "policy", "rm" ]
date: "2017-01-27"
menu:
  main:
    parent: smn_cli
---

```markdown
Remove policy

Usage:
  hbm policy rm [name] [flags]

Aliases:
  rm, remove
```

## Examples

```bash
# hbm policy ls
NAME                GROUP               COLLECTION
policy1             group1              collection1
policy2             group2              collection2
# hbm policy rm policy1
# hbm policy ls
NAME                GROUP               COLLECTION
policy2             group2              collection2
```

## Related information

* [policy_add](policy_add.md)
* [policy_find](policy_find.md)
* [policy_ls](policy_ls.md)
