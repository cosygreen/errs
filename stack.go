package errs

import (
	"errors"

	"github.com/tehsphinx/cstack"
)

// WithStack adds a call stack to the error.
func WithStack(err error) error {
	if err == nil {
		return nil
	}

	const skipStack = 3

	return stackError{
		err:   err,
		stack: cstack.CallStack(skipStack),
	}
}

type stackError struct {
	err   error
	stack cstack.Stack
}

func (s stackError) Error() string {
	return s.err.Error()
}

func (s stackError) Unwrap() error {
	return s.err
}

// GetStack returns the stack from the error chain if there was one added using WithStack.
func GetStack(err error) (cstack.Stack, bool) {
	var r stackError
	if !errors.As(err, &r) {
		return nil, false
	}
	return r.stack, true
}

// FormatStack implements marshalling of the error stack.
//
// Usage with zerolog:
//
//	zerolog.ErrorStackMarshaler = func(err error) interface{} {
//		stack := FormatStack(err)
//		if stack == "" {
//			return nil
//		}
//		return stack
//	}
func FormatStack(err error) string {
	st, ok := GetStack(err)
	if !ok {
		return ""
	}

	return st.DefaultFormat()
}

// StackFrameInfo implements marshalling of the error stack as a slice of frames.
//
// Usage with zerolog:
//
//	zerolog.ErrorStackMarshaler = func(err error) interface{} {
//		return StackInfo(err)
//	}
func StackFrameInfo(err error) []cstack.FrameInfo {
	st, ok := GetStack(err)
	if !ok {
		return nil
	}

	return st.StackInfo()
}
