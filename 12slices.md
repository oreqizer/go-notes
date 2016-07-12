# Slices

Slices are dynamically-sized *views* of arrays. They are much more common than arrays. They can be of any type, also of other slices.

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

### Make

`make` allocates a zeroed array and returns a slice that refers to that array. Used for creating dynamically-sized arrays.

```go
a := make([]int, 5)  // len(a) is 5
```

*Capacity* can also be provided as a third argument:

```go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

### Append

`append` is used for appending new elements to a slice.

Definition `func append(s []T, vs ...T) []T`:

* **s** is the original slice
* **vs** are the new values to append

If the *slice*'s capacity is not sufficient, a new *array* is allocated, and a pointer to the new array is returned.

### Range

The `range` form of the `for` loop iterates over a slice or a map. First value returned when iterating is the *index*, second is the *value*:

```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

If we want to ignore the index, simply name the variable `_`:

```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for _, v := range pow {
		fmt.Printf("%d, ", v) // 1, 2, 4, 8, 16, 32, 64, 128,
	}
}
```

## Exercise

Implement `Pic`. It should return a slice of length `dy`, each element of which is a slice `of` dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

```go
package main

import "golang.org/x/tour/pic"

func fillRow(row []uint8, dx, y int) {
	for x := 0; x < dx; x++ {
		row[x] = uint8((x+y)/2)
		// row[x] = uint8(x*y)
		// row[x] = uint8(x^y)
	}
}

func Pic(dx, dy int) [][]uint8 {
	full := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		full[y] = make([]uint8, dx)
		fillRow(full[y], dx, y)
	}
	return full
}

func main() {
	pic.Show(Pic)
}
```
