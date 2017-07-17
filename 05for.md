# For loop

Just like C, except no parentheses `( )` around the top expression, and braces `{ }` are required.

The initialization expression parts are:

* init statement
* condition
* post statement

```go
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

### While

The *init* and *post* statements are optional (it becomes C's `while`):

```go
package main

import "fmt"

func main() {
	sum := 1
    	// we could keep semicolons:
    	// for ; sum < 1000; {
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```

### Infinite

Writing only `for` makes an infinite loop:

```go
package main

import "fmt"

func main() {
	for {
        	fmt.Println("oreqizer is 1337")
	}
}
```
