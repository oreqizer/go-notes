# Switch

Just like everywhere else, with some little differences:

* `case` body breaks automatically (unless it ends with a `fallthrough`)
* `case` can has multiple conditions separated with `,`
* `switch` can has initialization expression as a `for` or `if`

```go
package main

import (
    "fmt"
    "runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os { // init expression
	case "darwin":
		fmt.Println("OS X.")
	case "freebsd", "linux": // multiple choice
		fmt.Println("Nerd OS.")
	default:
		fmt.Printf("%s.", os)
	}

	switch x := 1; x {
	case 1:
		fmt.Print("One. ")
        fallthrough // executes the next `case` automatically
	case 2, 3, 4, 5:
		fmt.Print("Two. Three, four, five!") // One. Two. Three, four, five!
	}
}
```

Function calls in a `case` are evaluated only if the `switch` gets there (no condition was met before).

Switch without a condition is like `switch true {` - simply checks for the first truthy value in the `case` statements:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```

Useful for writing long `if`-`else` chains.
