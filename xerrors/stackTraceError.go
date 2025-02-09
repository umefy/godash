package xerrors

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
)

const packageName = "godash/xerrors"

type StackTraceError struct {
	err   error
	stack string
}

func (st StackTraceError) Error() string {
	return st.err.Error()
}

func (st StackTraceError) Unwrap() error {
	return st.err
}

func (st StackTraceError) errorWithStack() string {
	return fmt.Sprintf("%v\nStackTrace: %s", st.err, st.stack)
}

func filterStackTrace(stack []byte) string {
	lines := strings.Split(string(stack), "\n")
	var filteredLines []string
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.Contains(line, "runtime/debug") || strings.Contains(line, packageName) || strings.Contains(line, "[running]") {
			continue
		}
		filteredLines = append(filteredLines, line)
	}
	return strings.Join(filteredLines, "\n")
}

func newStackTraceError(err error) StackTraceError {
	var stErr StackTraceError
	if ok := errors.As(err, &stErr); ok {
		return StackTraceError{
			err:   err,
			stack: stErr.stack,
		}
	}
	stack := debug.Stack()

	return StackTraceError{
		err:   err,
		stack: filterStackTrace(stack),
	}
}
