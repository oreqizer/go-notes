# Maps

Contains **key/value** pairs.

* zero value is `nil`, no keys can be added
* `map[K]V` has keys of type `K` and values of type `V`
* `make` returns an initialized map of the given type
* values accessed with `[]`, like this: `m[key]`

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
```

### Map literals

Like `struct` literals, but *keys* are required:

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
```

If the top-level type is just a type name, it can be omitted:

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
```

### Mutating maps

* reading: `elem = m[key]`
* writing: `m[key] = elem`
* deleting: `delete(m, key)`

Testing if a key exists:

```go
elem, ok := m[key]
// elem is the value if ok, else it's the value's type's zero value
// ok is a bool, if the key exists
```

## Exercise

Implement `WordCount`. It should return a map of the counts of each “word” in the string `s`. The `wc.Test` function runs a test suite against the provided function and prints success or failure.

```go
package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word] = wordMap[word] + 1
	}
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
```
