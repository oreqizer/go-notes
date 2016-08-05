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
