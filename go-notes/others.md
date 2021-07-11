# Other Notes

</br>

## List of Contents:
### 1. [Go mod](#content-1)
### 2. [What “accept interfaces, return structs” means in Go](#content-2)
### 3. [Preemptive Interface Anti-Pattern in Go](#content-3)


</br>

---

## Contents:

## [Go mod](https://golang.org/doc/code) <span id="content-1"></span>

</br>

- We don't have to declare the module path belonging to a repository.
- A module can be defined locally without belonging to a repository.
- `go install` command builds the module, producing an executable binary.
- The binaries are install to the bin subdirectory of the default GOPATH.
- The easiest way to make your module available for others to use is usually to make its module path match the URL for the repository.
- `go mod tidy` command adds missing module requirements for imported packages.
- `go clean -modcache`: remove all downloaded modules.

</br>

---

## [What “accept interfaces, return structs” means in Go](https://medium.com/@cep21/what-accept-interfaces-return-structs-means-in-go-2fe879e25ee8) <span id="content-2"></span>


- > All problems in computer science can be solved by another level of indirection, except of course for the problem of too many indirections
  > >David J. Wheeler
- Interfaces abstract away from structures in Go
- Tt doesn’t make sense to create this complexity until it’s needed
- > Always [abstract] things when you actually need them, never when you just foresee that you need them.
- You can control the return values of a function, but you can't control the input type.
- That's why it's better for us to accept interface, instead of concrete types.
- Another aspect of simplification is removing unnecessary detail.
- If you don't need some recipes to make something, then don't list it on your need-list.
- Check this snippet:
```go
type Database struct{ }
func (d *Database) AddUser(s string) {...}
func (d *Database) RemoveUser(s string) {...}
func NewUser(d *Database, firstName string, lastName string) {
  d.AddUser(firstName + lastName)
}
```
- On the above code, we define database to have 2 methods. But on NewUser database job is just to add new user. No need to add RemoveUser
- This is probably the better way:
```go
type DatabaseWriter interface {
  AddUser(string)
}
func NewUser(d DatabaseWriter, firstName string, lastName string) {
  d.AddUser(firstName + lastName)
}
```

</br>

---

## [Preemptive Interface Anti-Pattern in Go](https://medium.com/@cep21/preemptive-interface-anti-pattern-in-go-54c18ac0668a) <span id="content-3"></span>
- Interfaces are a way to describe behavior
- Preempttive interfaces are when developer codes to an interface before an actual need arises.
- Example:
```go
type Auth interface {
  GetUser() (User, error)
}
type authImpl struct {
  // ...
}
func NewAuth() Auth {
  return &authImpl
}
```
- You have to change the code if you use this
```java
// auth.java
public class Auth {
  public boolean canAction() {
    // ...
  }
}
// logic.java
public class Logic {
  public void takeAction(Auth a) {
    // ...
  }
}
```
- For example, you want to take any objects in takeAction as long as it has canAction inside it. How it would be?
- Better code in java
```java
// auth.java
public interface Auth {
  public boolean canAction()
}
// authimpl.java
class AuthImpl implements Auth {
}
// logic.java
public class Logic {
  public void takeAction(Auth a) {
    // ...
  }
}
```
- (Personal notes) It'e better to pass a pointer of a struct rather than the struct itself. It makes sure that we check for its nullity.
- Go uses implicit interface, which means concrete objects (structs) don't need to explicitly defined that they are using this interface. It's different from explicit interface like in Java.
- Usually you don't need preemptive interface in go.
- Go is at its most powerful when interface definitions are small.
- In the standard library, most interface definitions are a single method
- Accepting interfaces gives your API the greatest flexibility and returning structs allows the people reading your code to quickly navigate to the correct function
- Unnecessary abstraction creates unnecessary complication. Don’t over complicate code until it’s needed.

</br>

---

## References:
- https://medium.com/@cep21/what-accept-interfaces-return-structs-means-in-go-2fe879e25ee8
- https://medium.com/@cep21/preemptive-interface-anti-pattern-in-go-54c18ac0668a