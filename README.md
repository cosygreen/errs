# errs

The `errs` package is meant to be used in combination with the default errors package.
It adds error types that can be used in combination e.g. to have a call stack in the error chain.

## Usage

```go
package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tehsphinx/cstack"
	"github.com/tehsphinx/errs"
)

func main() {
	err1 := createErr()
	err2 := wrapErr()

	stackStr := errs.MarshalStack(err1)
	stackSlice := errs.StackFrames(err2)

	jsonBts, _ := json.Marshal(stackSlice)
	fmt.Println(stackStr)
	fmt.Println(string(jsonBts))
}

func createErr() error {
	return errs.New("new error with stack")
}

func wrapErr() error {
	err := errors.New("wrapped with stack: error without stack")
	return errs.WithStack(err)
}
```
