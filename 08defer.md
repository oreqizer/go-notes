# Defer

`defer` defers the execution of a function until the surrounding function returns.

```go
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```

`defer` calls are pushed to a stack, thus executed in a *LIFO* manner:

```go
package main

import "fmt"

func main() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Print(i)
	}
	fmt.Println("done")
}
```

Output:
```
counting
done
9876543210
```
