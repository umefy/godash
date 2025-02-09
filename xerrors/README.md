# xerrors

`xerrors` package is an error package that can used to track the stack trace error.

```go
// create new error
xerrors.New(errors.New("some error"))

// wrap error
xerrors.Wrap(err. "customer message")

// wrapf error
xerrors.Wrapf(err, "%s some message", "This is")

// Get stack trace
xerrors.GetErrorWithStack(err) // the error should be created or wrapped by xerrors
```
