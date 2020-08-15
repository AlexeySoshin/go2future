# Go2Future
Deferred value implementation in Go2   

Provides a higher-level concurrency construct, that eliminates the need for `go` statement and channels in most common use-cases.

## Usage
Creating a new Future:

```go 
f1 := NewFuture(func() (string, error) {
    time.Sleep(1000)
    return "F1", nil
})
```

Waiting for its result:
```go
result := f1.Get()
```

Discarding the Future in case it's not needed:
```go
f3 := NewFuture(func() (string, error) {
    time.Sleep(100)
    fmt.Println("I'm done!")
    return "F3", nil
})
f3.Cancel()
```

## Installation
```bash
./install.sh
```
This will set your Go to development version that supports generics and provides you with `go2go` tool


# License 
MIT