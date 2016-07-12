# Arrays

In Go, arrays have fixed length.

* `[n]T` is an array of values of type T

```go
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1]) // Hello world
	fmt.Println(a) // [Hello world]
}
```

We can initialize an array with values:

```go
package main

import "fmt"

func main() {
    primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes) // [2 3 5 7 11 13]
}
```
