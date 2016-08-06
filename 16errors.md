# Errors

Go's errors are *values*. The `error` type is a built-in interface similar to `fmt.Stringer`:

```go
type error interface {
    Error() string
}
```

Errors are commonly returned as a *second value* of a function:

```go
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```

## Exercise

`Sqrt` should return a non-nil `error` value when given a negative number, as it doesn't support complex numbers.

```go
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for {
		delta := z
		z = z - (z * z - x)/(2*z)
		if math.Abs(delta - z) < 0.01 {
			break
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
```
