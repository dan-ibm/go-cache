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
	"time"
)

func main() {
	mapCache := cache.New()


	// You can omit duration argument, by default, its 10 seconds (cache.DefaultDuration)
	mapCache.Set("userId", 42, time.Second * 5)
	// You can change time sleep value, to check if expiration works
	time.Sleep(time.Second * 3)
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


