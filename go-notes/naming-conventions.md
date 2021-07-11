# Naming Conventions

</br>

## List of Contents:
### 1. [Naming Conventions in Go: Short but Descriptive](#content-1)
### 2. [What's in a name?](#content-2)
### 3. [GoLang Naming Rules and Conventions](#content-3)
### 4. [Package names](#content-4)
### 5. [Style guideline for Go packages](#content-5)


</br>

---

## Contents:

## [Naming Conventions in Go: Short but Descriptive](https://betterprogramming.pub/naming-conventions-in-go-short-but-descriptive-1fa7c6d2f32a) <span id="content-1"></span>

> “There are only two hard things in Computer Science: cache invalidation and naming things.” — Phil Karlton

The purpose for conventions is to make life easier. To make sure won't wonder what this variable for and this function does.

### MixedCaps (CamelCase)
When writing variable, function, or interface, if we don't need to expose it to other packages, then keep it local.
If in future use, we need to expose this entity to other packages then make it importable.
`camelCase` for local naming entity and `CamelCase` for global entitiy.

### Interface names
> “By convention, one-method interfaces are named by the method name plus an -er suffix or similar modification to construct an agent noun: Reader, Writer, Formatter, CloseNotifier etc.” — Go’s official documentation

This rule only applies to single method interface, the rule of thumb is `MethodName + er = InterfaceName`.
If our interface contains more than one method, then it is up to us for its name. But keep it reasonable and clear. 

### Getters
> “There’s nothing wrong with providing getters and setters yourself, and it’s often appropriate to do so, but it’s neither idiomatic nor necessary to put Get into the getter's name.”
```go
owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}
```

### Unwritten Rules
1. Shorter variable names </br>
Our variable name should be short, but descriptive
- Single-letter identifier </br>
Used for local variables with limited scope, e.g naming index or value in for loop.
- Shorthand name </br>
Use obvious name for our variable, avoid using ambiguous name.
Such as`pid // Bad (does it refer to podID or personID or productID?)`
  
2. Unique Names </br>
Use uppercase letter to name acronyms such as API, HTTP, ID, or DB. (e.g don't name it userId but userID)

3. Line length </br>
Avoid long lines

</br>

---

## [What's in a name?](https://talks.golang.org/2014/names.slide#1) <span id="content-2"></span>

Names matter because to write a better code, we also have to make sure our code is clear and readable.

A good name is:
- Consistent
- Short (easy to type)
- Accurate (easy to understand, what we write explain itself)

A rule of thumb
The greater the distance between a name's declaration and its uses,
the longer the name should be. If we define some variable at the top line, but we use it somewhere several lines at the bottom, make sure our name longer but better distinction.

Use MixedCase!

### Local variables
Common variable/type combinations may use really short names:

Prefer i to index. </br>
Prefer r to reader. </br>
Prefer b to buffer. </br>

### Parameters
Function parameters are like local variables,
but they also serve as documentation.

Where the types are descriptive, they should be short:
```go
func AfterFunc(d Duration, f func()) *Timer

func Escape(w io.Writer, s []byte)
```
Where the types are more ambiguous, the names may provide documentation:
```go
func Unix(sec, nsec int64) Time

func HasPrefix(s, prefix []byte) bool
```

### Return values
Return values on exported functions should only be named for documentation purposes.
```go
func Copy(dst Writer, src Reader) (written int64, err error)

func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)
```

### Receivers
Receivers are a special kind of argument.

By convention, they are one or two characters that reflect the receiver type,
because they typically appear on almost every line:

```go
func (b *Buffer) Read(p []byte) (n int, err error)

func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request)

func (r Rectangle) Size() Point
```

Receiver names should be consistent across a type's methods.
(Don't use r in one method and rdr in another.)

### Exported package-level names
Don't use stutter naming such as strings.StringReader. That's why we have bytes.Buffer and strings.Reader,
not bytes.ByteBuffer and strings.StringReader.

### Interface Types

Interfaces that specify just one method are usually just that function name with 'er' appended to it.
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

### Errors
Error types should be of the form FooError:
```go
type ExitError struct {
...
}
```
Error values should be of the form ErrFoo:
```go
var ErrFormat = errors.New("image: unknown format")
```


### Packages
Choose package names that lend meaning to the names they export. 
Steer clear of util, common, and the like.

### Import paths
The last component of a package path should be the same as the package name.

`"compress/gzip" // package gzip` </br>
Avoid stutter in repository and package paths:

`"code.google.com/p/goauth2/oauth2" // bad; my fault` </br>
For libraries, it often works to put the package code in the repo root:

`"github.com/golang/oauth2" // package oauth2`</br>
Also avoid upper case letters (not all file systems are case sensitive).

</br>

---

## [GoLang Naming Rules and Conventions](https://medium.com/@kdnotes/golang-naming-rules-and-conventions-8efeecd23b68) <span id="content-3"></span>

### Files
1. Go follows a convention where source files are all lower case with underscore separating multiple words.
2. Compound file names are separated with _
3. File names that begin with “.” or “_” are ignored by the go tool
4. Files with the suffix _test.go are only compiled and run by the go test tool.

### Constants
Constant should use all capital letters and use underscore _ to separate words.

</br>

---

## [Package names](https://blog.golang.org/package-names) <span id="content-4"></span>

Go code is organized into packages. Good package names make code better. A package's name provides context for its contents, making it easier for clients to understand what the package is for and how to use it. The name also helps package maintainers determine what does and does not belong in the package as it evolves. Well-named packages make it easier to find the code you need.

Good package names are short and clear. They are lower case, with no under_scores or mixedCaps. They are often simple nouns, such as:
- time (provides functionality for measuring and displaying time)
- list (implements a doubly linked list)
- http (provides HTTP client and server implementations)

Abbreviate judiciously. Package names may be abbreviated when the abbreviation is familiar to the programmer. Widely-used packages often have compressed names:
- strconv (string conversion)
- syscall (system call)
- fmt (formatted I/O)

On the other hand, if abbreviating a package name makes it ambiguous or unclear, don't do it.
Don't steal good names from the user. Avoid giving a package a name that is commonly used in client code. For example, the buffered I/O package is called bufio, not buf, since buf is a good variable name for a buffer.

### Naming package contents
1. Avoid repetition </br>
http package use http.Server not http.HTTPServer. Don't use stutter naming
2. Avoid meaningless package names </br>
Packages named util, common, or misc provide clients with no sense of what the package contains. This makes it harder for clients to use the package and makes it harder for maintainers to keep the package focused.
Bad example:
```go
package util
func NewStringSet(...string) map[string]bool {...}
func SortStringSet(map[string]bool) []string {...}

// How to use it
set := util.NewStringSet("c", "a", "b")
fmt.Println(util.SortStringSet(set))
```
Better approach
```go
package stringset
func New(...string) map[string]bool {...}
func Sort(map[string]bool) []string {...}

//How to use it
set := stringset.New("c", "a", "b")
fmt.Println(stringset.Sort(set))
```

### Avoid unnecessary package name collisions
While packages in different directories may have the same name, packages that are frequently used together should have distinct names. This reduces confusion and the need for local renaming in client code. For the same reason, avoid using the same name as popular standard packages like io or http.

</br>

---

## [Style guideline for Go packages](https://rakyll.org/style-packages/) <span id="content-5"></span>

### 1. Don't export from main </br>
An identifier may be exported to permit access to it from another package.

Main packages are not importable, so exporting identifiers from main packages is unnecessary. Don’t export identifiers from a main package if you are building the package to a binary.

### 2. Package naming
- Lowercase only </br>
  Package names should be lowercase. Don’t use snake_case or camelCase in package names.
- Short, but representative names </br>
  Package names should be short, but should be unique and representative. Users of the package should be able to grasp its purpose from just the package’s name.
  
### 3. Clean import paths
Avoid exposing your custom repository structure to your users. Align well with the GOPATH conventions. Avoid having src/, pkg/ sections in your import paths. </br>
`github.com/user/repo/src/httputil   // DON'T DO IT, AVOID SRC!!`

### 4. No plurals
In go, package names are not plural. This is surprising to programmers who came from other languages and are retaining an old habit of pluralizing names. Don’t name a package httputils, but httputil!

### 5. Renames should follow the same rules
```go
import (
    gourl "net/url"

    "myother.com/url"
)
```

</br>

---

## References
- https://betterprogramming.pub/naming-conventions-in-go-short-but-descriptive-1fa7c6d2f32a
- https://talks.golang.org/2014/names.slide#19
- https://medium.com/@kdnotes/golang-naming-rules-and-conventions-8efeecd23b68
- https://blog.golang.org/package-names
- https://rakyll.org/style-packages/