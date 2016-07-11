# Variables

`var` declares a list of variables, with it's type at the end.

```go
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
```

They can also be initialized (with type being inferred, if not specified):

```go
package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
```

**Inside of a function:**

When in a function, we can use `:=` shorthand for `var` declaration with implicit type.

```go
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```

### Constants

Declared using `const`. Cannot be declared using `:=`. They can be **string**, **boolean**, or **numeric**.

A good practice is to name them *uppercase*.

```go
package main

import "fmt"

const Pi = 3.14 // Pi is exported

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
```
