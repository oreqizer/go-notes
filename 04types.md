# Types

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

These are all the available basic types. Type declaration may be *factored* as imports:

```go
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
```

### Zero values

Variables declared without an explicit initial value are given their *zero value*.

The zero value is:

* `0` for numeric types
* `false` for the boolean type
* `""` (the empty string) for strings

### Type conversions

The expression `T(v)` converts the value `v` to the type `T`.

```go
var i int = 42 // i := 42
var f float64 = float64(i) // f := float64(i)
var u uint = uint(f) // u := uint(f)
```

In Go, all conversions are explicit.

*Note:*
The **int**, **uint**, and **uintptr** types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use **int** unless you have a specific reason to use a sized or unsigned integer type.
