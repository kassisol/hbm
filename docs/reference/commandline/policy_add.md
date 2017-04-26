---
description: The policy add command description and usage
keywords:
- policy, add
menu:
  main:
    parent: smn_cli
title: policy add
---

# hbm policy add
***

```markdown
Add policy

Usage:
  hbm policy add [name] [flags]

Flags:
  -c, --cluster string      Set cluster
      --collection string   Set collection
  -g, --group string        Set group
```

## Examples

```bash
# hbm policy add --group group1 --cluster cluster1 --collection collection1 policy1
# hbm policy ls
NAME                GROUP               CLUSTER             COLLECTION
policy1             group1              cluster1            collection1
```

## Related information

* [policy_find](policy_find.md)
* [policy_ls](policy_ls.md)
* [policy_rm](policy_rm.md)
