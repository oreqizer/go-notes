# Interfaces

An *interface type* is a set of method signatures. They are implemented implicitly - we only need to define all the methods of the interface's signature.

### Interface value

Interfaces consist of:

* interface **value**
* concrete **type**

An interface value holds a value of a specific underlying concrete type.

Calling a method on an interface *value* executes the method of the same name on its underlying *type*.

```go
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

### Nil underlying value

* if a concrete value is `nil`, the method is called with a `nil` receiver
* interface value that holds a `nil` concrete value is itself non-`nil`

### Nil interface value

A `nil` interface value holds neither value nor concrete type.

Calling a method on a `nil` interface is a run-time error because there is no type inside the interface tuple to indicate which *concrete* method to call.

### Empty interface

The *interface type* that specifies zero methods is known as the *empty interface*:

```go
interface{}
```

An empty interface may hold values of any type. (Every type implements at least *zero methods*)

Empty interfaces are used for values of unknown type. `fmt.Print` takes any number of arguments of type `interface{}`.

### Type assertion

A *type assertion* provides access to an interface value's underlying concrete value.

```go
t := i.(T)
```

* asserts that the *interface value* `i` holds the *concrete type* `T`
* assigns the underlying `T` value to the variable `t`. If `i` does not hold a `T`, the statement will trigger a panic

Testing if an interface value holds a type:

```go
t, ok := i.(T)
```

* if yes, `t` will be the underlying value and `ok` will be `true`
* else, `t` is a *zero value* of the type and `ok` will be `false`

### Type switch

Like a multiple *type assertion*. Type `T` is replaced with the keyword `type`:

```go
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
```

### Conventions

If an interface has one method, it's name is usually the method's name with *er* suffix.

*Stringer*

Any `type` implementing *Stringer* can represent itself as a `string`.

```go
type Stringer interface {
    String() string
}
```

*Reader*

The `io` package specifies the `io.Reader` interface, which represents the read end of a stream of data. It has many implementations in the standard library like *files, network connections, compressors, ciphers*, and others.

```go
type Reader interface {
	Read(b []byte) (n int, err error)
}
```

`Read` populates the given **byte slice** with data and returns the number of bytes populated and an `error` value. It returns an `io.EOF` error when the stream ends.

## Exercise

Define your own `Image` type, implement the necessary [methods](https://golang.org/pkg/image/#Image), and call `pic.ShowImage`.

```go
package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{255, 255, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
```
