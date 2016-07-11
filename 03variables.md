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

**Numeric constants** are high-precision values. An untyped `const` takes the type needed by it's context.

```go
package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)
```

*Note:*
An int can store at maximum a 64-bit integer, and sometimes less.
