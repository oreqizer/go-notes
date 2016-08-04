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

### Arguments

Functions receive a **copy** of an argument. To modify the argument, we need to pass a *pointer* to the value.

The exception is reference-based types like `slice`, `map`, `function`, `pointer` and `channel`.

### Variadic arguments

`...` allows variadic arguments *of the same type*. It captures them to a slice.

```go
func foo(params ...int) {
    fmt.Println(len(params))
}
```

We can also spread a slice when suffixing `...`:

```go
nums := []int{1, 2, 3}

foo(nums...) // 3
```

### Function values

Functions are *first class values* - they can be passed around:

```go
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
```

### Function closures

Closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

`adder` returns a closure. Closure is bound to its own `sum` variable:

```go
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// both 'pos' and 'neg' have their own 'sum' variable
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
```
