# Slices

Slices are dynamically-sized *views* of arrays. They are much more common than arrays.

* `[]T` is a slice with elements of type `T`
* `a[0:5]` creates a slice with first 5 elements of the array `a`

### Slices are like references to arrays

Changing a slice's element changes the value for all the slices sharing the reference.

```go
package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2] // is [John Paul]
	b := names[1:3] // is [George Ringo]

	b[0] = "XXX"
	fmt.Println(a, b) // [John XXX] [XXX George]
	fmt.Println(names) // [John XXX George Ringo]
}
```

### Slice literals

Like array literal without the length:

* builds the array
* references it

so doing `s := []bool{true, true, false}` is the same as:

```go
a := [3]bool{true, true, false}
s := a[0:3]
```

### Slice defaults

* `0` for the low bound
* *length* for the high bound

These slice expressions are thus equivalent:

```go
a := [3]bool{true, true, false}
a[0:3]
a[:3]
a[0:]
a[:]
```

### Length and capacity

Length:

* function `len`
* shows the number of elements the slice contains

Capacity:

* function `cap`
* number of elements from the 1st element of the **slice** until the last element of the underlying **array**.

Length can be extended up to the *capacity* of the slice.

```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s) // len=0 cap=6 []

	// Extend its length.
	s = s[:4]
	printSlice(s) // len=4 cap=6 [2 3 5 7]

	// Drop its first two values.
	s = s[2:]
	printSlice(s) // len=2 cap=4 [5 7]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

### Nil slice

The zero value of a slice is `nil`.

* has 0 length and capacity
* has no underlying array
