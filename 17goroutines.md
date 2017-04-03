# Goroutines

Each concurrently executing activity is a *goroutine*. They are spawned using the `go` statement before a function call.

By default there's only one goroutine - one that calls the `main` function.

## Channels

Goroutines are the activities, *channels* are the connections between them. They are conduits for a particular type:

* `chan int` - a channel with values of type `int`
* `chan string` - a channel with values of type `string`

Creating a channel:

```go
ch := make(chan int) // ch has type 'chan int'
```

A channel is a *reference* to the data structure created by `make`. Its zero value is `nil`.

**Operations:**

Channels use the `<-` operator for both *send* and *receive* actions. Received values can be discarded.

```go
ch <- x // send 'x' to the channel
x = <-ch // receive from the channel and store to 'x'
<-ch // receive from the channel and discard the value
close(ch) // close the channel
```

The `close` function closes the channel. Send actions trigger `panic`, receive actions take any leftover values. If no more values are left, receive will return zero values for the type.

### Buffered channels

To create a *buffered* channel, simply pass a second argument to the `make` function that will indicate the channel's capacity. Zero means an *unbuffered* channel, as if we passed nothing.

```go
ch = make(chan int)    // unbuffered channel
ch = make(chan int, 0) // unbuffered channel
ch = make(chan int, 3) // buffered channel with capacity 3
```

The built-in `cap` function can be used to determine a channel's capacity, and `len` can be used to get the number of currently buffered elements:

```go
ch := make(chan int, 3)
fmt.Println(cap(ch)) // 3
fmt.Println(len(ch)) // 0
ch <- 13
ch <- 37
fmt.Println(len(ch)) // 2
```

**Differences:**

Unbuffered (*synchronous*) channels:
* *sending* things first blocks the goroutine until the value was *received*
* *receiving* a value first blocks the goroutine until the value was *sent*

Buffered channels:
* *sending* only blocks the goroutine if the channel is **full**
* *receiving* only blocks the goroutine if the channel is **empty**

### Channel iteration

To check if a channel is *drained*, we can accept a second `ok` parameter:

```go
x, ok := <-ch
```

`ok` is `true` if the channel still has values, `false` if it's drained. Another option is to use `range`:

```go
for x := range ch {
    // do something with channel values
}
```

### Unidirectional channels

Channel type can be declared *unidirectional*, either only for sending or receiving:

```go
func counter(out chan<- int) {
    for x := 0; x < 100; x++ {
        out <- x
    }
    close(out)
}

func printer(in <-chan int) {
    for v := range in {
        fmt.Println(v)
    }
}

func main() {
    nums := make(chan int)
    go counter(nums)
    printer(nums)
}
```

Not obeying the channel direction is detected at compile time.

### Select

A way to *multiplex* channels. A `select` waits until the communication for some `case` is ready to proceed, then executes its statements and moves on.

```go
select {
case <- ch1:
    // wait for ch1, discard value
case x := <- ch2:
    // wait for ch2, save value to x
case ch3 <- y:
    // send to ch3
default:
    // optional fallback
}
```

An empty-bodied `select {}` waits forever.

> **Note:** If multiple cases are ready at the same time, `select` picks one **randomly**.
