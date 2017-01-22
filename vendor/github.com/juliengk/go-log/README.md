# Go Log

## Example

```
package main

import (
	"github.com/juliengk/go-log"
	"github.com/juliengk/go-log/driver"
)

func main() {
	l, _ := log.NewDriver("standard", nil)

	l.WithFields(driver.Fields{
		"user": "root",
	}).Info("some info")
}
```
