package errors

import (
	"fmt"
	"io"
)

type StackError struct {
	error
	*stack
}

func (w *StackError) Cause() error { return w.error }

// Unwrap provides compatibility for Go 1.13 error chains.
func (w *StackError) Unwrap() error { return w.error }

func (w *StackError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", w.Cause())
			w.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, w.Error())
	case 'q':
		fmt.Fprintf(s, "%q", w.Error())
	}
}
