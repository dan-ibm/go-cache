Go Cache
=========================
Go Cache helps you to simply store data as key value. You can set, get or delete data by using key

See it in action:

## Example:

```go
package main

import (
	"fmt"
	"github.com/dan-ibm/go-cache"
)

func main() {
	mapCache := cache.New()

	mapCache.Set("userId", 42)
	userId, err := mapCache.Get("userId")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(userId)
	}

	if err = mapCache.Delete("userId"); err != nil {
		fmt.Println(err)
	}
	if userId, err = mapCache.Get("userId"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(userId)
	}
}

```


