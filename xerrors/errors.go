package xerrors

import (
	"errors"
	"fmt"
)

func New(err error) stackTraceError {
	return newStackTraceError(err)
}

func Wrap(err error, msg string) stackTraceError {
	return newStackTraceError(fmt.Errorf("%w; %s", err, msg))
}

func Wrapf(err error, format string, args ...any) stackTraceError {
	return newStackTraceError(fmt.Errorf("%w; %s", err, fmt.Sprintf(format, args...)))
}

func GetErrorWithStack(err error) string {
	var stErr stackTraceError
	if ok := errors.As(err, &stErr); ok {
		return stErr.errorWithStack()
	}
	return err.Error()
}
