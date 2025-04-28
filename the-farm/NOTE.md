The error interface

Error handling is not done via exceptions in Go. Instead, errors are normal values of types that implement the built-in `error` interface. The `error` interface is very minimal. It contains only one method `Error()` that returns the error message as a string.

```go
type error interface {
  Error() string
}
```