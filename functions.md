# Functions

`main` is special - used as an entry point of an application in the `main` module.

```go
package main

import "fmt"

func add(x, y int) int { // same as '(x int, y int)'
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

Functions have:
* 0-n typed parameters
* 0-n typed return values
