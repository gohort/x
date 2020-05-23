# x

General purpose custom errors library.

This library is mostly useful for when categorizing errors in a router to
Identify the proper status code.

Using this package is safer than assigning your errors status codes directly
to be transport agnostic.

```go
package main

import (
    "fmt"

    "github.com/gohort/x"
)

var (
    ErrNotFound = x.NewError("not found")
    ErrTicketing = x.NewError("ticketing system")
    ErrReading = x.NewError("reading")
)

// x has a custom Errorf for categorizing errors. This is helpful when
// developing an API with a backend business logic.
func getTicket(key string) (*Value, error) {
    val, err := find(key)
    if err != nil {
        return nil, x.Errorf(ErrTicketing, "find: %w", err)
    }
    // ...
}

// errors can be compared using errors.Is and be unwrapped.
func handler(w http.ResponseWriter, _ *http.Request) {
    if err := handle(); err != nil {
        if errors.Is(err, ErrNotFound) {
            // ...
        }
    }
    // ...
}

// Errors can be used standalone like a standard error.
func find(key string) (*Value, error) {
    if valueNotFound(key) {
        return nil, ErrNotFound
    }
    // ...
}
```