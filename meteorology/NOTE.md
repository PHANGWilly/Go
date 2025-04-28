`Stringer` is an interface for defining the string format of values.

The interface consists of a single String method:

```go
type Stringer interface {
    String() string
}
```

Types that want to implement this interface must have a `String()` method that returns a human-friendly string representation of the type. The `fmt` package (and many others) will look for this method to format and print values.