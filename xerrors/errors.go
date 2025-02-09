package xerrors

import (
	"errors"
	"fmt"
)

func New(err error) StackTraceError {
	return newStackTraceError(err)
}

func Wrap(err error, msg string) StackTraceError {
	return newStackTraceError(fmt.Errorf("%w; %s", err, msg))
}

func Wrapf(err error, format string, args ...any) StackTraceError {
	return newStackTraceError(fmt.Errorf("%w; %s", err, fmt.Sprintf(format, args...)))
}

func GetErrorWithStack(err error) string {
	var stErr StackTraceError
	if ok := errors.As(err, &stErr); ok {
		return stErr.errorWithStack()
	}
	return err.Error()
}
