# Error Handling

</br>

## List of Contents:

### 1. [Don’t just check errors, handle them gracefully](#content-1)

### 2. [Error handling and Go](#content-2)

### 3. [REST API Error Handling in Go: Behavioral Type Assertion](#content-3)

### 4. [Go Error Handling (Part 3) — The errors Package](#content-4)

### 5. [Decorating Go Errors](#content-5)

### 6. [Golang Microservices: Working and Dealing with Errors](#content-6)

</br>

---

## Contents:

## [Don’t just check errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully) <span id="content-1"></span>

We wanted there to be a single way to do error handling, something that we could teach all Go programmers by rote, just as we might teach mathematics, or the alphabet.

However, there is no single way to handle errors. Instead, Go's error handling can be classified into three core strategies.

### 1. Sentinel errors

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

### 2. Error types

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

### 3. Opaque errors

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

### Don’t just check errors, handle them gracefully

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

### Annotating errors

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

### Only handle errors once

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

</br>

---

## [Error handling and Go](https://blog.golang.org/error-handling-and-go) <span id="content-2"></span>

- Standard usage:

```go
f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
// do something with the open *File f
```

- If you want to create custom error 1

```go
// errorString is a trivial implementation of error.
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}
```

- Error message that probably you want to consider

```go
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    // implementation
}
```

- Formatting error message, this returns error type

```go
fmt.Errorf("math: square root of negative number %g", f)
```

- Another example of custom error

```go
type SyntaxError struct {
    msg    string // description of error
    Offset int64  // error occurred after reading Offset bytes
}

func (e *SyntaxError) Error() string { return e.msg }
```

- Simplify repetitive error handling

```go
func init() {
    http.HandleFunc("/view", viewRecord)
}

func viewRecord(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
    record := new(Record)
    if err := datastore.Get(c, key, record); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    if err := viewTemplate.Execute(w, record); err != nil {
        http.Error(w, err.Error(), 500)
    }
}

type appHandler func(http.ResponseWriter, *http.Request) error

func viewRecord(w http.ResponseWriter, r *http.Request) error {
  c := appengine.NewContext(r)
  key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
  record := new(Record)
  if err := datastore.Get(c, key, record); err != nil {
      return err
  }
  return viewTemplate.Execute(w, record)
}

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if err := fn(w, r); err != nil {
    http.Error(w, err.Error(), 500)
  }
}
```

</br>

---

## [REST API Error Handling in Go: Behavioral Type Assertion](https://medium.com/@ozdemir.zynl/rest-api-error-handling-in-go-behavioral-type-assertion-509d93636afd) <span id="content-3"></span>

### Introduction

- Robust Go applications should deal with errors gracefully.

### General guidelines

- Always check errors. Never ignore them unless you have a very, very good reason.
- Always log error details somewhere. Errors are the most valuable information we have, to fix bugs, potential failures and performance issues.
- Add breadcrumbs to error logs, such as Client IP, request headers/body, user information/events. The more data you have, the easier it is to debug the problem.
- Handle errors only once. Wrap them with additional context and return to caller function if necessary. It is much more maintainable to handle (taking an action on error such as logging) them in a specific exit point/middleware.
- Differentiate client and server errors. Don’t overthink and come up with hundreds of different error types. Keep it simple. It is almost always enough to have just two generic types: Client Errors (4xx), which means something was probably wrong with the request data and can be corrected by the client. Server Errors (5xx), are unexpected errors and they usually point out that there are some bugs needs to be fixed in the code.
- Use HTTP response status codes to indicate something went wrong (400 Bad Request, 404 Not Found…). Although there are different status codes you can use for different situations, they are usually not enough to describe the error on their own. It is best to add more context to the response body.
- Define a good error response structure from the beginning. Whether it is a simple JSON response with only detail and status or a complex one with domain , type , track_id , title , helpUrl ; stick with it. Try not to have different structures on different endpoints. Be consistent as much as possible. You can follow RFC7807.
- Make your Client Error messages human readable. Use code to identify errors, messages are for users. Don’t scare the user off with complex error messages, they are best to be descriptive and easy to grasp.
- Don’t share the details/stack trace of Server Errors with the client. They might contain your code logic and secrets. Log them on a secure platform with a unique id and return this id to the client. So that you can find relevant records/metrics easily when a customer contacts you with this trace id.
- Document your errors both for developers and users. Try to be as descriptive as possible with error details. It is better to suggest possible reasons/solutions for users in the documentation.

### Some background — Errors in Go

- By convention, errors are the last return value from functions, that implements the built-in interface error:
  ```go
  type error interface {
      Error() string
  }
  ```
- Custom error:
  ```go
  type CustomError string
  func (err CustomError) Error() string {
      return string(err)
  }
  ```
- Explicitly check errors:
  ```go
  f, err := os.Open(filename)
  if err != nil {
      // Handle the error ...
  }
  ```

### Handling errors in a sample Go REST API project

- Example:

  ```go
  package main

  import (
  	"encoding/json"
  	"fmt"
  	"io/ioutil"
  	"log"
  	"net/http"
  )

  type loginSchema struct {
  	Username string `json:"username"`
  	Password string `json:"password"`
  }

  func loginUser(username string, password string) (bool, error) {...}

  func loginHandler(w http.ResponseWriter, r *http.Request) {
  	if r.Method != http.MethodPost {
  		w.WriteHeader(405) // Return 405 Method Not Allowed.
  		return
  	}
  	// Read request body.
  	body, err := ioutil.ReadAll(r.Body)
  	if err != nil {
  		log.Printf("Body read error, %v", err)
  		w.WriteHeader(500) // Return 500 Internal Server Error.
  		return
  	}

  	// Parse body as json.
  	var schema loginSchema
  	if err = json.Unmarshal(body, &schema); err != nil {
  		log.Printf("Body parse error, %v", err)
  		w.WriteHeader(400) // Return 400 Bad Request.
  		return
  	}

  	ok, err := loginUser(schema.Username, schema.Password)
  	if err != nil {
  		log.Printf("Login user DB error, %v", err)
  		w.WriteHeader(500) // Return 500 Internal Server Error.
  		return
  	}

  	if !ok {
  		log.Printf("Unauthorized access for user: %v", schema.Username)
  		w.WriteHeader(401) // Wrong password or username, Return 401.
  		return
  	}
  	w.WriteHeader(200) // Successfully logged in.
  }

  func main() {
  	http.HandleFunc("/login/", loginHandler)
  	log.Fatal(http.ListenAndServe(":8080", nil))
  }
  ```

- Rewriting the implemenation:
  ![New](https://miro.medium.com/max/700/1*bkPy-jPUV5F5j9w7ECy5UA.png)
- Remember that we want to have two different main error types: Client Error for 4xx errors and Server Error (or Internal Error) for 5xx. We can declare interfaces based on the behavior we expect from these two types and use type assertion on rootHandler to make some decisions about the error.
- Client erro:
  ```go
  // ClientError is an error whose details to be shared with client.
  type ClientError interface {
  	Error() string
  	// ResponseBody returns response body.
  	ResponseBody() ([]byte, error)
  	// ResponseHeaders returns http status code and headers.
  	ResponseHeaders() (int, map[string]string)
  }
  ```
- Explanation for above:
  - `ResponseBody() ([]byte, error)` : Returns JSON response body of the error (title, message, error code…) in bytes. (\*Getting response body as bytes from one method is not the best solution, see Further Improvements section.)
  - `ResponseHeaders() (int, map[string]string)` : Returns HTTP status code (4xx, 5xx) and headers (content type, no-cache…) of response.
  - `Error()` string , this is necessary to make every ClientError and error at the same time.
- Example:

  ```go
  // HTTPError implements ClientError interface.
  type HTTPError struct {
  	Cause  error  `json:"-"`
  	Detail string `json:"detail"`
  	Status int    `json:"-"`
  }

  func (e *HTTPError) Error() string {
  	if e.Cause == nil {
  		return e.Detail
  	}
  	return e.Detail + " : " + e.Cause.Error()
  }

  // ResponseBody returns JSON response body.
  func (e *HTTPError) ResponseBody() ([]byte, error) {
  	body, err := json.Marshal(e)
  	if err != nil {
  		return nil, fmt.Errorf("Error while parsing response body: %v", err)
  	}
  	return body, nil
  }

  // ResponseHeaders returns http status code and headers.
  func (e *HTTPError) ResponseHeaders() (int, map[string]string) {
  	return e.Status, map[string]string{
  		"Content-Type": "application/json; charset=utf-8",
  	}
  }

  func NewHTTPError(err error, status int, detail string) error {
  	return &HTTPError{
  		Cause:  err,
  		Detail: detail,
  		Status: status,
  	}
  }
  ```

- HTTPError has all the information we need to log the error and return a proper HTTP response to the client:
  - Cause : Original error (unmarshall errors, network errors…) which caused this HTTP error, set it to nil if there isn’t any.
  - Detail : message to return in JSON response. Ex: { "detail": "Wrong password" } .
  - Status : HTTP response status code. Ex: 400, 401, 405…
- The final version:

  ```go
  package main

  import (
  	"encoding/json"
  	"fmt"
  	"io/ioutil"
  	"log"
  	"net/http"
  )

  type loginSchema struct {
  	Username string `json:"username"`
  	Password string `json:"password"`
  }

  // Return `true`, nil if given user and password exists in database.
  func loginUser(username string, password string) (bool, error) {...}

  // Use as a wrapper around the handler functions.
  type rootHandler func(http.ResponseWriter, *http.Request) error

  func loginHandler(w http.ResponseWriter, r *http.Request) error {
  	if r.Method != http.MethodPost {
  		return NewHTTPError(nil, 405, "Method not allowed.")
  	}

  	body, err := ioutil.ReadAll(r.Body) // Read request body.
  	if err != nil {
  		return fmt.Errorf("Request body read error : %v", err)
  	}

  	// Parse body as json.
  	var schema loginSchema
  	if err = json.Unmarshal(body, &schema); err != nil {
  		return NewHTTPError(err, 400, "Bad request : invalid JSON.")
  	}

  	ok, err := loginUser(schema.Username, schema.Password)
  	if err != nil {
  		return fmt.Errorf("loginUser DB error : %v", err)
  	}

  	if !ok { // Authentication failed.
  		return NewHTTPError(nil, 401, "Wrong password or username.")
  	}
  	w.WriteHeader(200) // Successfully authenticated. Return access token?
  	return nil
  }

  // rootHandler implements http.Handler interface.
  func (fn rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  	err := fn(w, r) // Call handler function
  	if err == nil {
  		return
  	}
  	// This is where our error handling logic starts.
  	log.Printf("An error accured: %v", err) // Log the error.

  	clientError, ok := err.(ClientError) // Check if it is a ClientError.
  	if !ok {
  		// If the error is not ClientError, assume that it is ServerError.
  		w.WriteHeader(500) // return 500 Internal Server Error.
  		return
  	}

  	body, err := clientError.ResponseBody() // Try to get response body of ClientError.
  	if err != nil {
  		log.Printf("An error accured: %v", err)
  		w.WriteHeader(500)
  		return
  	}
  	status, headers := clientError.ResponseHeaders() // Get http status code and headers.
  	for k, v := range headers {
  		w.Header().Set(k, v)
  	}
  	w.WriteHeader(status)
  	w.Write(body)
  }

  func main() {
  	// Notice rootHandler.
  	http.Handle("/login/", rootHandler(loginHandler))
  	log.Fatal(http.ListenAndServe(":8080", nil))
  }
  ```

</br>

---

## [Go Error Handling (Part 3) — The errors Package](https://sher-chowdhury.medium.com/go-error-handling-part-3-the-errors-package-1cb73f6eb0ce) <span id="content-4"></span>

- On the other hand, the errors package doesn’t require any of this extra setup; instead, you can go straight to creating error variables using the package’s errors.New() function.
  ![](https://miro.medium.com/max/1000/1*lmUD4CMR5_79bAF2CySGWA.png)
- It’s best practice for functions to always return an error value.

</br>

---

## [Decorating Go Errors](https://medium.com/spectro-cloud/decorating-go-error-d1db60bb9249) <span id="content-5"></span>

### Introduction

- Go treats the error as a value with a predefined type, technically an interface.
- However, writing a multi-layered architecture application and exposing the features with APIs demands error treatment with much more contextual information than just a value.

### Custom Type

- As we will override the default Go error type we have to start with a custom error type which will be interpreted within the application, and is also of Go error type. Hence we will introduce new custom error interface composing the Go error:
  ```go
  type GoError struct {
      error
  }
  ```

### Contextual Data

- When we say the error is a value in Go, it is of string value - any type which has the Error() string function implemented qualifies for the error type.
- Treating string values as errors complicates the error interpretations across the layers, as interpreting the error string is not the right approach. So let’s decouple the string with the error code.
  ```go
  type GoError struct {
    error
    Code    string
  }
  ```
- Now the error interpretation will be based on the error Code rather than the string. Let’s further decouple the error string with the contextual data which allows internationalization with thei18N package
  ```go
  type GoError struct {
    error
    Code    string
    Data    map[string]interface{}
  }
  ```
- Data contains the contextual data to construct the error string. The error string can be templatized with the data:
  ```text
  //i18N def
  "InvalidParamValue": "Invalid parameter value '{{.actual}}', expected '{{.expected}}' for '{{.name}}'"
  ```

### Cause

- The error can occur in any layer and it is necessary to provide the option for each layer to interpret the error and further wrap with additional contextual information without losing the original error value.
- The GoError can be further decorated with the Causes which will hold the entire error stack.
  ```go
  type GoError struct {
    error
    Code    string
    Data    map[string]interface{}
    Causes  []error
  }
  ```
- Causes is an array type if it has to hold multiple error data and is set to the base error type to include the third-party error for the cause within the application.
- Tagging the layer component will help to identify the layer where the error has occurred, and unnecessary error wraps can be avoided.
- For example, if the error component of servicetype occurs in the service layer, then the wrapping error might not be required
- Component information checks will help to prevent exposing the errors which a user shouldn’t be informed about, like Database errors.
  ```go
  type GoError struct {
    error
    Code      string
    Data      map[string]interface{}
    Causes    []error
    Component ErrComponent
  }
  type ErrComponent string
  const (
    ErrService  ErrComponent = "service"
    ErrRepo     ErrComponent = "repository"
    ErrLib      ErrComponent = "library"
  )
  ```

### Response Type

- Adding an error response type will support the error categorization for easy interpretation. For example, the errors can be categorized with response types like NotFound, and this information can be used for errors like DbRecordNotFound , ResourceNotFound , UserNotFound, and so on.
- Example:

  ```go
  type GoError struct {
    error
    Code         string
    Data         map[string]interface{}
    Causes       []error
    Component    ErrComponent
    ResponseType ResponseErrType
  }
  type ResponseErrType string

  const (
    BadRequest    ResponseErrType = "BadRequest"
    Forbidden     ResponseErrType = "Forbidden"
    NotFound      ResponseErrType = "NotFound"
    AlreadyExists ResponseErrType = "AlreadyExists"
  )
  ```

### Retry

- Snippet:
  ```go
  type GoError struct {
    error
    Code         string
    Message      string
    Data         map[string]interface{}
    Causes       []error
    Component    ErrComponent
    ResponseType ResponseErrType
    Retryable    bool
  }
  ```

### GoError Interface

- Error checking can be simplified by defining an explicit error interface definition with the implementation of GoError:

  ```go
  package goerr
  type Error interface {
    error

    Code() string
    Message() string
    Cause() error
    Causes() []error
    Data() map[string]interface{}
    String() string
    ResponseErrType() ResponseErrType
    SetResponseType(r ResponseErrType) Error
    Component() ErrComponent
    SetComponent(c ErrComponent) Error
    Retryable() bool
    SetRetryable() Error
  }
  ```

### Error Abstraction

- With the above-mentioned decorations, it is important to build the abstraction over an error and keep these decorations in a single place and provide the reusability of the error function:
  ```go
  func ResourceNotFound(id, kind string, cause error) GoError {
    data := map[string]interface{}{"kind": kind, "id": id}
    return GoError{
        Code:         "ResourceNotFound",
        Data:         data,
        Causes:       []error{cause},
        Component:    ErrService,
        ResponseType: NotFound,
        Retryable:    false,
    }
  }
  ```
- This error function abstracts the ResourceNotFound and the developer will use this function instead of constructing the new error object every time:
  ```go
  //UserService
  user, err := u.repo.FindUser(ctx, userId)
  if err != nil {
    if err.ResponseType == NotFound {
        return ResourceNotFound(userUid, "User", err)
    }
    return err
  }
  ```

</br>

---

## [Golang Microservices: Working and Dealing with Errors](https://www.youtube.com/watch?v=uQOfXL6IFmQ) <span id="content-6"></span>

### Errors in Go

- Use errors.Is

  ```go
  if err == io.ErrUnexpectedEOF // Before
  if errors.Is(err, io.ErrUnexpectedEOF) // After
  ```

- Use errors.As

  ```go
  if e, ok := err.(*os.PathError); ok // Before
  var e *os.PathError  // After
  if errors.As)err, &e)
  ```

- Use fmt.Errorf("message: %w", err)
- Use errors.Unwrap

</br>

---

## References:

- https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
- https://www.youtube.com/watch?v=lsBF58Q-DnY
- https://blog.golang.org/error-handling-and-go
- https://medium.com/@ozdemir.zynl/rest-api-error-handling-in-go-behavioral-type-assertion-509d93636afd
- https://medium.com/spectro-cloud/decorating-go-error-d1db60bb9249
- https://www.youtube.com/watch?v=uQOfXL6IFmQ
