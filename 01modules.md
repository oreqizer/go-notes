# Modules

### Package:

```go
package main
```

`main` is special - indicated entry point of a program. Package names are lowercase names of the modules to be imported.

### Import:

Simple import:

```go
package main

import "fmt"
import "math/rand"
```

Factored import (preferred):

```go
package main

import (
	"fmt"
	"math/rand"
)
```

### Export:

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
