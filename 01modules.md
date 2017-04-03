# Modules

### Package

A single **folder**. Consists of *1-n* source files. They share package scope, meaning they can see each other's globals, even if they are *not exported* (are *lowercase*).

```go
package main
```

`main` is special - indicated entry point of a program. Package names are lowercase names of the modules to be imported.

### Import

**Simple import:**

```go
package main

import "fmt"
import "math/rand"
```

**Factored import** (preferred):

```go
package main

import (
	"fmt"
	"math/rand"
)
```

**Named import:**

Useful, when packages of the same name are imported.

```go
package main

import f "fmt" // fmt is now f
```

We can combine these:

```go
package main

import (
	f "fmt"
	"math/rand"
)
```

### Export

Everything scoped globally starting with **uppercase** letter is exported. We access package exports like so:

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi) // 3.141592653589
}
```

### Init:

A special `init` function can be defined in a file that runs at after the package has been initialized:

```go
func init() {
	http.HandleFunc("/api", handler)
}
```
