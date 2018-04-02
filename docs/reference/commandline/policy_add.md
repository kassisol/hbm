---
title: "hbm policy add"
description: "The policy add command description and usage"
keywords: [ "policy", "add" ]
date: "2017-01-27"
menu:
  docs:
    parent: "hbm_cli_policy"
github_edit: "https://github.com/kassisol/hbm/edit/master/docs/reference/commandline/policy_add.md"
---

```markdown
Add policy

Usage:
  hbm policy add [name] [flags]

Flags:
  -c, --collection string   Set collection
  -g, --group string        Set group
  -h, --help                help for add
```

## Examples

```bash
# hbm policy add --group group1 --collection collection1 policy1
# hbm policy ls
NAME                GROUP               COLLECTION
policy1             group1              collection1
```

## Related information

* [policy_find](policy_find.md)
* [policy_ls](policy_ls.md)
* [policy_rm](policy_rm.md)
