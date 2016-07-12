# Structs

A `struct` is a collection of fields. It's fields are accessed with a `.`.

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
    v := Vertex{1, 2}
	fmt.Println(v) // {1, 2}
    v.X = 4
    fmt.Println(v.X) // 4
}
```

### Struct pointers

Accessing field through a *struct pointer* implicitly dereferences the field's value:

```go
v := Vertex{1, 2}
p := &v
// p.X is the same as (*p).X or v.X
```

### Struct literals

We can construct a `struct` either by specifying values of the fields in the correct order (X, Y in this case), or we can use the `key: value` syntax.

Prefixing the struct declaration with a `&` returns a pointer instead of the value.

```go
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
```
