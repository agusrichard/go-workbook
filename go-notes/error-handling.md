# Error Handling

> Don’t just check errors, handle them gracefully

## Contents:

### [Don’t just check errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)

We wanted there to be a single way to do error handling, something that we could teach all Go programmers by rote, just as we might teach mathematics, or the alphabet.

However, there is no single way to handle errors. Instead, Go's error handling can be classified into three core strategies.

#### 1. Sentinel errors
```go
if err == ErrSomething { … }
```

Examples include values like io.EOF or low level errors like the constants in the syscall package, like syscall.ENOENT.
Using sentinel values is the least flexible error handling strategy, as the caller must compare the result to predeclared value using the equality operator. This presents a problem when you want to provide more context, as returning a different error would will break the equality check.

Instead the caller will be forced to look at the output of the error‘s Error method to see if it matches a specific string.

Never inspect the output of error.Error. As an aside, I believe you should never inspect the output of the error.Error method. The Error method on the error interface exists for humans, not code. 

Never the less, comparing the string form of an error is, in my opinion, a code smell, and you should try to avoid it.

Sentinel errors create a dependency between two packages
By far the worst problem with sentinel error values is they create a source code dependency between two packages. As an example, to check if an error is equal to io.EOF, your code must import the io package.

Conclusion: avoid sentinel errors

#### 2. Error types
```go
if err, ok := err.(SomeType); ok { … }
```

```go
type MyError struct {
        Msg string
        File string
        Line int
}

func (e *MyError) Error() string { 
        return fmt.Sprintf("%s:%d: %s”, e.File, e.Line, e.Msg)
}

return &MyError{"Something happened", “server.go", 42}
```
Because MyError error is a type, callers can use type assertion to extract the extra context from the error.
```go
err := something()
switch err := err.(type) {
case nil:
        // call succeeded, nothing to do
case *MyError:
        fmt.Println(“error occurred on line:”, err.Line)
default:
// unknown error
}
```

An excellent example of this is the os.PathError type which annotates the underlying error with the operation it was trying to perform, and the file it was trying to use.
```go
// PathError records an error and the operation
// and file path that caused it.
type PathError struct {
        Op   string
        Path string
        Err  error // the cause
}

func (e *PathError) Error() string
```

Problems with error types. If your code implements an interface whose contract requires a specific error type, all implementors of that interface need to depend on the package that defines the error type.

Conclusion: avoid error types

While error types are better than sentinel error values, because they can capture more context about what went wrong, error types share many of the problems of error values.

#### 3. Opaque errors
In my opinion this is the most flexible error handling strategy as it requires the least coupling between your code and caller.

I call this style opaque error handling, because while you know an error occurred, you don’t have the ability to see inside the error. As the caller, all you know about the result of the operation is that it worked, or it didn’t.

This is all there is to opaque error handling–just return the error without assuming anything about its contents. If you adopt this position, then error handling can become significantly more useful as a debugging aid.

```go
import “github.com/quux/bar”

func fn() error {
        x, err := bar.Foo()
        if err != nil {
                return err
        }
        // use x
}
```

> Assert errors for behaviour, not type

In this case rather than asserting the error is a specific type or value, we can assert that the error implements a particular behaviour. Consider this example:
```go
type temporary interface {
	Temporary() bool
}

// IsTemporary returns true if err is temporary.
func IsTemporary(err error) bool {
    te, ok := err.(temporary)
    return ok && te.Temporary()
}
```

If the error does not implement the temporary interface; that is, it does not have a Temporary method, then then error is not temporary. The key here is this logic can be implemented without importing the package that defines the error or indeed knowing anything about err‘s underlying type–we’re simply interested in its behaviour.

#### Don’t just check errors, handle them gracefully

This brings me to a second Go proverb that I want to talk about; don’t just check errors, handle them gracefully. Can you suggest some problems with the following piece of code?
```go
func AuthenticateRequest(r *Request) error {
        err := authenticate(r.User)
        if err != nil {
                return err
        }
        return nil
}
```

An obvious suggestion is that the five lines of the function could be replaced with
```go
return authenticate(r.User)
```

But this is the simple stuff that everyone should be catching in code review. More fundamentally the problem with this code is I cannot tell where the original error came from.
If authenticate returns an error, then AuthenticateRequest will return the error to its caller, who will probably do the same, and so on. At the top of the program the main body of the program will print the error to the screen or a log file, and all that will be printed is: No such file or directory. 

There is no information of file and line where the error was generated. There is no stack trace of the call stack leading up to the error.

Donovan and Kernighan’s The Go Programming Language recommends that you add context to the error path using fmt.Errorf

```go
func AuthenticateRequest(r *Request) error {
err := authenticate(r.User)
if err != nil {
return fmt.Errorf("authenticate failed: %v", err)
}
return nil
}
```
But as we saw earlier, this pattern is incompatible with the use of sentinel error values or type assertions, because converting the error value to a string, merging it with another string, then converting it back to an error with fmt.Errorf breaks equality and destroys any context in the original error.

#### Annotating errors

```go
// Wrap annotates cause with a message.
func Wrap(cause error, message string) error

// Cause unwraps an annotated error.
func Cause(err error) error
```

Consider this function:
```go
func ReadFile(path string) ([]byte, error) {
        f, err := os.Open(path)
        if err != nil {
                return nil, errors.Wrap(err, "open failed")
        } 
        defer f.Close()
 
        buf, err := ioutil.ReadAll(f)
        if err != nil {
                return nil, errors.Wrap(err, "read failed")
        }
        return buf, nil
}
```

```go
func ReadConfig() ([]byte, error) {
        home := os.Getenv("HOME")
        config, err := ReadFile(filepath.Join(home, ".settings.xml"))
        return config, errors.Wrap(err, "could not read config")
}
 
func main() {
        _, err := ReadConfig()
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
}
```

If we replace `fmt.Println()` with `errors.Println()`We will get stack trace.

Now we’ve introduced the concept of wrapping errors to produce a stack, we need to talk about the reverse, unwrapping them. This is the domain of the errors.Cause function.
```go
// IsTemporary returns true if err is temporary.
func IsTemporary(err error) bool {
        te, ok := errors.Cause(err).(temporary)
        return ok && te.Temporary()
}
```

In operation, whenever you need to check an error matches a specific value or type, you should first recover the original error using the errors.Cause function.

#### Only handle errors once

If you make less than one decision, you’re ignoring the error. As we see here, the error from w.Write is being discarded.
```go
func Write(w io.Writer, buf []byte) {
        w.Write(buf)
}
```

But making more than one decision in response to a single error is also problematic.
```go
func Write(w io.Writer, buf []byte) error {
        _, err := w.Write(buf)
        if err != nil {
                // annotated error goes to log file
                log.Println("unable to write:", err)
 
                // unannotated error returned to caller
                return err
        }
        return nil
}
```
```go
func Write(w io.Write, buf []byte) error {
        _, err := w.Write(buf)
        return errors.Wrap(err, "write failed")
}
```

For maximum flexibility I recommend that you try to treat all errors as opaque. In the situations where you cannot do that, assert errors for behaviour, not type or value.

Minimise the number of sentinel error values in your program and convert errors to opaque errors by wrapping them with errors.Wrap as soon as they occur.

Finally, use errors.Cause to recover the underlying error if you need to inspect it.

## References:
- https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully