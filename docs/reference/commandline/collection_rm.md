---
description: The collection rm command description and usage
keywords:
- collection, rm
menu:
  main:
    parent: smn_cli
title: collection rm
---

# hbm collection rm
***

```markdown
Remove collection from the whitelist

Usage:
  hbm collection rm [name] [flags]

Aliases:
  rm, remove
```

## Examples

```bash
# hbm collection ls
NAME                RESOURCES
collection1         resource1
collection2         resource2
# hbm collection rm collection1
# hbm collection ls
NAME                RESOURCES
collection2         resource2
```

## Related information

* [collection_add](collection_add.md)
* [collection_find](collection_find.md)
* [collection_ls](collection_ls.md)
