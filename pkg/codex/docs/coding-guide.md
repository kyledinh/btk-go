# Coding Guide
> Opininated guide for coding in this repository. 

* Correctness > Capability > Clarity
* Parse > Process > Publish 
* Dream > Document > Design > Demo > Deliver 

**Table of Contents** 
- [Configurations](#configuration)
- [Guidelines](#guidelines)
   - [Testing](#testing)
   - [Logging](#logging)
   - [Docker](#docker)
- [Performance](#performance)
- [Style](#style)
- [Patterns](#patterns)
   - [Errors](#errors)
- [Linting](#linting)
- [Tooling](#tooling)
- [References](#references)
   - [Shell Guide](#shell-guide)

<hr>

## Configuration 

- Viper Config: https://github.com/spf13/viper

<hr>

## Guidelines 

### Testing 

- Table Driven Testing
- Unit Testing
- Integration Testing
- Performance Testing

### Logging 

- spew: https://pkg.go.dev/github.com/davecgh/go-spew
- zap: https://github.com/uber-go/zap
- https://www.datadoghq.com/blog/go-logging/

### Docker

<hr>

## Performance 

### Monitoring/Metrics

<hr>

## Style 

<hr>

## Patterns 

### Errors


There are few options for declaring errors.
Consider the following before picking the option best suited for your use case.

- Does the caller need to match the error so that they can handle it?
  If yes, we must support the [`errors.Is`] or [`errors.As`] functions
  by declaring a top-level error variable or a custom type.
- Is the error message a static string,
  or is it a dynamic string that requires contextual information?
  For the former, we can use [`errors.New`], but for the latter we must
  use [`fmt.Errorf`] or a custom error type.
- Are we propagating a new error returned by a downstream function?
  If so, see the [section on error wrapping](#error-wrapping).

[`errors.Is`]: https://golang.org/pkg/errors/#Is
[`errors.As`]: https://golang.org/pkg/errors/#As

| Error matching? | Error Message | Guidance                            |
|-----------------|---------------|-------------------------------------|
| No              | static        | [`errors.New`]                      |
| No              | dynamic       | [`fmt.Errorf`]                      |
| Yes             | static        | top-level `var` with [`errors.New`] |
| Yes             | dynamic       | custom `error` type                 |

[`errors.New`]: https://golang.org/pkg/errors/#New
[`fmt.Errorf`]: https://golang.org/pkg/fmt/#Errorf

For example,
use [`errors.New`] for an error with a static string.
Export this error as a variable to support matching it with `errors.Is`
if the caller needs to match and handle this error.

<table>
<thead><tr><th>No error matching</th><th>Error matching</th></tr></thead>
<tbody>
<tr><td>

```go
// package foo

func Open() error {
  return errors.New("could not open")
}

// package bar

if err := foo.Open(); err != nil {
  // Can't handle the error.
  panic("unknown error")
}
```

</td><td>

```go
// package foo

var ErrCouldNotOpen = errors.New("could not open")

func Open() error {
  return ErrCouldNotOpen
}

// package bar

if err := foo.Open(); err != nil {
  if errors.Is(err, foo.ErrCouldNotOpen) {
    // handle the error
  } else {
    panic("unknown error")
  }
}
```

</td></tr>
</tbody></table>

For an error with a dynamic string,
use [`fmt.Errorf`] if the caller does not need to match it,
and a custom `error` if the caller does need to match it.

<table>
<thead><tr><th>No error matching</th><th>Error matching</th></tr></thead>
<tbody>
<tr><td>

```go
// package foo

func Open(file string) error {
  return fmt.Errorf("file %q not found", file)
}

// package bar

if err := foo.Open("testfile.txt"); err != nil {
  // Can't handle the error.
  panic("unknown error")
}
```

</td><td>

```go
// package foo

type NotFoundError struct {
  File string
}

func (e *NotFoundError) Error() string {
  return fmt.Sprintf("file %q not found", e.File)
}

func Open(file string) error {
  return &NotFoundError{File: file}
}


// package bar

if err := foo.Open("testfile.txt"); err != nil {
  var notFound *NotFoundError
  if errors.As(err, &notFound) {
    // handle the error
  } else {
    panic("unknown error")
  }
}
```

</td></tr>
</tbody></table>

Compose an error message from the function call with a prefined error message in the moxerr packagCompose an error message from the function call with a prefined error message in the moxerr package

<table>
<thead><tr><th>Wrapping Error Message</th></tr></thead>
<tbody>
<tr><td>

```go
// package moxerr
var (
	ErrResourceNotFound = errors.New("resource not found")
	ErrCLIAction        = errors.New("cli action failed to execute")
)

type WrappedError struct {
	Message string
	MoxErr  *error
}

func (we *WrappedError) Error() string {
	return fmt.Sprintf("message: %s", we.Message)
}

func NewWrappedError(message string, err *error) *WrappedError {
	return &WrappedError{
		Message: message,
		MoxErr:  err,
	}
}
```

</td></tr>
</tbody></table>


* https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully


<hr>

## Linting 

- errcheck: https://github.com/kisielk/errcheck
- goimports: https://godoc.org/golang.org/x/tools/cmd/goimports
- golint: https://github.com/golang/lint
- golangci-lint: https://github.com/golangci/golangci-lint
- govet: https://golang.org/cmd/vet/
- staticcheck: https://staticcheck.io/

<hr>

## Tooling 

- Delve: https://github.com/go-delve/delve
- Go Tools (pprof) video: https://www.youtube.com/watch?v=uBjoTxosSys

<hr>

## References

### Shell Guide

* https://www.hexnode.com/blogs/the-ultimate-guide-to-mac-shell-scripting/
* https://news.learnenough.com/macos-bash-zshell
