# Go Mount

## Example

```
package main

import (
	"fmt"
	"log"

	"github.com/juliengk/go-mount"
)

func main() {
	entries, err := mount.New()
	if err != nil {
		log.Fatal(err)
	}

	entry, err := entries.Find("/sys")
	if err != nil {
		log.Fatal(err)
	}

	result := entry.FindOption("nosuid")

	fmt.Println(result)
}
```
