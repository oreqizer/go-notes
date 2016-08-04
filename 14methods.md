# Methods

Special functions that have a *receiver* argument.

**Receiver:**

* is a type/interface
* has its own argument list before function name
* it's type **must** be defined in the same package

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```

### Pointer receivers

By default, functions receive a **copy** of a value (unless it's a *reference type*). Pointer receivers allow modifying the *receiver* - they are more common.

Pointer arguments **must** take a pointer:

```go
var v Vertex
ScaleFunc(v)  // Compile error!
ScaleFunc(&v) // OK
```

Pointer receivers can take a *pointer* or a *value*:

```go
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```

The same applies in reverse - value receivers can take a pointer, arguments can't.

**BEWARE**: if a *value receiver* receives a *pointer*, modifying it **will** modify the pointer's underlying value.
