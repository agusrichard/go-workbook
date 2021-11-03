# Best Practices and Patterns

</br>

## List of Contents:
### 1. [7 Code Patterns in Go I Can’t Live Without](#content-1)
### 2. [Rules pattern in Golang](#content-2)
### 3. [Practical SOLID in Golang: Open/Closed Principle](#content-3)


</br>

---

## Contents:

## [7 Code Patterns in Go I Can’t Live Without](https://betterprogramming.pub/7-code-patterns-in-go-i-cant-live-without-f46f72f58c4b) <span id="content-1"></span>

### Use Maps as a Set
- We often need to check the existence of something. For example, we might want to check if a file path/URL/ID has been visited before. In these cases, we can use map[string]struct{}. For example:
  ![Example](https://miro.medium.com/max/1400/1*-0GYVejiTBawRZL_FKuiaw.png)
- Using an empty struct, struct{}, means we don’t want the value part of the map to take up any space. 
- Sometimes people use map[string]bool, but benchmarks have shown that map[string]struct{} perform better both in memory and time.
  
### Using chan struct{} to Synchronize Goroutines
- In the following case, the channel carries a data type struct{}, which is an empty struct that takes up no space.
  ![Example](https://miro.medium.com/max/1400/1*FsG09La74LUt_7i5X1SrZg.png)


### Use Close to Broadcast
- Example:
  ![Example](https://miro.medium.com/max/700/1*wK0RRlIdJTD7-eLItTGtIA.png)
- Note that closing a channel to broadcast a signal works with any number of goroutines, so close(quit) also applies in the previous example.


### Use a Nil Channel to Block a Select Case
- Sometimes we need to disable certain cases in the select statement, like in the following function, which reads from an event source and sends events to the dispatch channel.
- Initial example:
  ![Use a Nil Channel to Block a Select Case](https://miro.medium.com/max/1400/1*EHmt4rQ0b7axtlxAb6T3Aw.png)
- Things we want to improve:
  - disable case s.dispatchC when len(pending) == 0 so that the code won’t panic
  - disable case s.eventSource when len(pending) >= maxPending to avoid allocating too much memory
- Solution:
  ![Solution](https://miro.medium.com/max/2400/1*W0SfNmdH1IvAZ5BkdfRQmQ.png)
- The trick here is to use an extra variable to turn on/off the original channel, and then put that variable to use in the select case.
  ![Diagram](https://miro.medium.com/max/700/1*md35tLW4O3va0MCd1EhD_Q.png)


### Non-blocking Read From a Channel
- Example:
  ![Example](https://miro.medium.com/max/700/1*0eijA0u9emxaUVTWoOaFuA.png)


### Anonymous Struct
- Sometimes we just want a container to hold a group of related values, and this kind of grouping won’t appear anywhere else. 
- In these cases, we don’t care about its type. In Python, we might create a dictionary or tuple for these kinds of situations. In Go, you can create an anonymous struct for this kind of situation.
- Example:
  ![Example](https://miro.medium.com/max/688/1*g1kYZ5caQJ_Eyij_h-wj4w.png)
- Note that struct {...} is the type of the variable Config — now you can access your config values via Config.Timeout.
- Using in a test:
  ![Example 1](https://miro.medium.com/max/700/1*gSivMFyN_jZZ7ClfkPPNgw.png)
  ![Example 2](https://miro.medium.com/max/700/1*adnvw2JCHhO_ZgnjtflOQw.png)
- There are definitely more scenarios that you might find anonymous structs handy. For example, when you want to parse the following JSON, you might define an anonymous struct with nested anonymous structs so that you can parse it with the encoding/json library.


### Wrapping Options With Functions
- Optional arguments in Python:
  ![Optional arguments](https://miro.medium.com/max/1400/1*nEQOY876RwVFLsQ-dJgJbQ.png)
- My favorite way of achieving this in Go is to wrap these options (port, proxy) using functions. That is, we can construct functions to apply our option values, which are stored in the function’s closure.
- How we can do the similar thing in Go:
  ![Similar thing in Go](https://miro.medium.com/max/2400/1*piiERiPSMD0X_tJi6aNFtg.png)


**[⬆ back to top](#list-of-contents)**

</br>

---

## [Rules pattern in Golang](https://medium.com/@michalkowal567/rules-pattern-in-golang-425765f3c646) <span id="content-1"></span>

### Introduction
- First example of code. Certainly bad practice:
  ```go
  func CalculateDiscount(c Customer) float64 {
    var discount float64
    if c.BirthDate.Before(time.Now().AddDate(-65,0,0)) {
      // senior discount 5%
      discount = 0.05
    }
    if c.BirthDate.Day() == time.Now().Day() && c.BirthDate.Month() == time.Now().Month() {
      // birthday discount 10%
      discount = math.Max(discount, 0.10)
    }
    if !c.FirstPurchaseDate.IsZero() {
      if c.FirstPurchaseDate.Before(time.Now().AddDate(-1,0,0)){
        // 1 year loyal customer, 10%
        discount = math.Max(discount, 0.10)
        if c.FirstPurchaseDate.Before(time.Now().AddDate(-5,0,0)) {
          // 5 years loyal 12%
          discount = math.Max(discount, 0.12)
          if c.FirstPurchaseDate.Before(time.Now().AddDate(-10,0,0)) {
            // 10 years loyal 20%
            discount = math.Max(discount, 0.2)
          }
        }

        if c.BirthDate.Day() == time.Now().Day() && c.BirthDate.Month() == time.Now().Month() {
          // birthday discount +10%
          discount += 0.1
        }
      }
    } else {
      // first time purchase discount of 15%
      discount = math.Max(discount, 0.15)
    }
    if c.IsVeteran {
      // veterans get 10%
      discount = math.Max(discount, 0.1)
    }

    return discount
  }
  ```

### Rules design pattern
- We are going to create a discount calculator, which is going to calculate the discount based on rules. This approach will reduce the complexity and duplication of this calculating logic.
- First of all, We need to create some abstraction, i.e. DiscountCalculator.
  ```go
  type DiscountCalculator struct {
  }
  ```
- We need to create a Rule interface also:
  ```go
  type Rule interface {
    CalculateDiscount(c Customer) float64
  }
  ```
- Bits of code:
  ```go
  func (c Customer) IsBirthday() bool {
    return c.BirthDate.Day() == time.Now().Day() && c.BirthDate.Month() == time.Now().Month()
  }
  ```
- Bits of code:
  ```go
  /* I know this struct is empty, but struct size depends only on the fields that are inside this struct, and we want to implementinterface. */
  type BirthdayDiscountRule struct {
  }

  func (b BirthdayDiscountRule) CalculateDiscount(c Customer) float64 {
    var discount float64

    if c.IsBirthday() {
        discount = 0.1
    }

    return discount
  }
  ```
- A bit of code:
  ```go

  type LoyalCustomerRule struct {
    yearsAsCustomer int
    discount        float64
  }

  func NewLoyalCustomerRule(yearsAsCustomer int, discount float64) *LoyalCustomerRule {
    return &LoyalCustomerRule{yearsAsCustomer: yearsAsCustomer, discount: discount}
  }

  func (l LoyalCustomerRule) CalculateDiscount(c Customer) float64 {
    var discount float64
    date := time.Now()

    if c.HasBeenLoyalForYears(l.yearsAsCustomer, date) {
      birthdayRule := BirthdayDiscountRule{}

      discount = l.discount + birthdayRule.CalculateDiscount(c)
    }
    
    return discount
  }
  ```
- A bit of code:
  ```go
  func NewDiscountCalculator() *DiscountCalculator {
    rules := []Rule {
      BirthdayDiscountRule{},
      SeniorDiscountRule{},
      VeteranDiscountRule{},
      NewLoyalCustomerRule(1, 0.1),
      NewLoyalCustomerRule(5, 0.12),
      NewLoyalCustomerRule(10, 0.2),
    }
    
    return &DiscountCalculator{rules: rules}
  }

  func (dc *DiscountCalculator) CalculateDiscount(c Customer) float64 {
    var discount float64
    
    for _, rule := range dc.rules {
      discount = math.Max(rule.CalculateDiscount(c), discount)
    }
    
    return discount
  }
  ```

**[⬆ back to top](#list-of-contents)**

</br>

---

## [Practical SOLID in Golang: Open/Closed Principle](https://levelup.gitconnected.com/practical-solid-in-golang-open-closed-principle-1dd361565452) <span id="content-3"></span>

### When we do not respect the Open/Closed Principle
- You should be able to extend the behavior of a system without having to modify that system.
- By checking the code example below, we can see what it means for some structures not to respect this principle and possible consequences.
  ```go
  import (
    "net/http"

    "github.com/ahmetb/go-linq"
    "github.com/gin-gonic/gin"
  )

  type PermissionChecker struct {
    //
    // some fields
    //
  }

  func (c *PermissionChecker) HasPermission(ctx *gin.Context, name string) bool {
    var permissions []string
    switch ctx.GetString("authType") {
    case "jwt":
      permissions = c.extractPermissionsFromJwt(ctx.Request.Header)
    case "basic":
      permissions = c.getPermissionsForBasicAuth(ctx.Request.Header)
    case "applicationKey":
      permissions = c.getPermissionsForApplicationKey(ctx.Query("applicationKey"))
    }
    
    var result []string
    linq.From(permissions).
      Where(func(permission interface{}) bool {
        return permission.(string) == name
      }).ToSlice(&result)

    return len(result) > 0
  }

  func (c *PermissionChecker) getPermissionsForApplicationKey(key string) []string {
    var result []string
    //
    // extract JWT from the request header
    //
    return result
  }

  func (c *PermissionChecker) getPermissionsForBasicAuth(h http.Header) []string {
    var result []string
    //
    // extract JWT from the request header
    //
    return result
  }

  func (c *PermissionChecker) extractPermissionsFromJwt(h http.Header) []string {
    var result []string
    //
    // extract JWT from the request header
    //
    return result
  }
  ```
- Here we have the primary method, HasPermission, which checks if permission with specific names is associated with the data within the Context.
- If we respect The Single Responsibility Principle, PermissionChecker is responsible for deciding if permission is inside the Context, and it does not have anything with the authorization process.
- Suppose we want to extend the authorization logic and add some new flow, such as keeping user data in session or using Digest Authorization. In that case, we need to make adaptations in PermissionChecker as well.
- Such implementation brings the explosion of issues:
  - PermissionChecker keeps the logic initially handled somewhere else.
  - Any adaptation of authorization logic, which may be a different module, requires adaptation in PermissionChecker.
  - To add a new way of extracting permissions, we always need to modify PermissionChecker.
  - Logic inside PermissionChecker inevitably grows with each new authorization flow.
  - Unit testing for PermissionChecker contains too many technical details about different extractions of permission.

### How we do respect The Open/Closed Principle
- The Open/Closed Principle says that software structures should be open for extension but closed for modification.
- In Object-Oriented Programming, we support such extensions by using different implementations for the same interface. In other words, we use polymorphism.
  ```go
  type PermissionProvider interface {
    Type() string
    GetPermissions(ctx *gin.Context) []string
  }

  type PermissionChecker struct {
    providers []PermissionProvider
    //
    // some fields
    //
  }

  func (c *PermissionChecker) HasPermission(ctx *gin.Context, name string) bool {
    var permissions []string
    for _, provider := range c.providers {
      if ctx.GetString("authType") != provider.Type() {
        continue
      }
      
      permissions = provider.GetPermissions(ctx)
      break
    }

    var result []string
    linq.From(permissions).
      Where(func(permission interface{}) bool {
        return permission.(string) == name
      }).ToSlice(&result)

    return len(result) > 0
  }
  ```
- Now, we introduct an interface to be implemented by e.g JwtPermissionProvider, or ApiKeyPermissionProvider, or AuthBasicPermissionProvider.

### Some more examples
- Example:
  ```go
  type PermissionProvider interface {
    Type() string
    GetPermissions(ctx *gin.Context) []string
  }

  type PermissionChecker struct {
    //
    // some fields
    //
  }

  func (c *PermissionChecker) HasPermission(ctx *gin.Context, provider PermissionProvider, name string) bool {
    permissions := provider.GetPermissions(ctx)

    var result []string
    linq.From(permissions).
      Where(func(permission interface{}) bool {
        return permission.(string) == name
      }).ToSlice(&result)

    return len(result) > 0
  }
  ```
- Example:
  ```go
  func GetCities(sourceType string, source string) ([]City, error) {
    var data []byte
    var err error

    if sourceType == "file" {
      data, err = ioutil.ReadFile(source)
      if err != nil {
        return nil, err
      }
    } else if sourceType == "link" {
      resp, err := http.Get(source)
      if err != nil {
        return nil, err
      }

      data, err = ioutil.ReadAll(resp.Body)
      if err != nil {
        return nil, err
      }
      defer resp.Body.Close()
    }

    var cities []City
    err = yaml.Unmarshal(data, &cities)
    if err != nil {
      return nil, err
    }

    return cities, nil
  }
  ```
- Example (better way):
  ```go
  type DataReader func(source string) ([]byte, error)

  func ReadFromFile(fileName string) ([]byte, error) {
    data, err := ioutil.ReadFile(fileName)
    if err != nil {
      return nil, err
    }

    return data, nil
  }

  func ReadFromLink(link string) ([]byte, error) {
    resp, err := http.Get(link)
    if err != nil {
      return nil, err
    }

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      return nil, err
    }
    defer resp.Body.Close()

    return data, nil
  }

  func GetCities(reader DataReader, source string) ([]City, error) {
    data, err := reader(source)
    if err != nil {
      return nil, err
    }

    var cities []City
    err = yaml.Unmarshal(data, &cities)
    if err != nil {
      return nil, err
    }

    return cities, nil
  }
  ```
- As you can see in the solution from above, in Go, we may define a new type that embeds the function. Here we described a new type, DataReader, representing a function type for reading raw data from some source.
- New methods ReadFromFile and ReadFromLink are actual implementations of the DataReader type.
- The GetCities method expects the actual implementation of DataReader as an argument, which then executes inside the function body and takes raw data.


**[⬆ back to top](#list-of-contents)**

</br>

---

## References:
- https://betterprogramming.pub/7-code-patterns-in-go-i-cant-live-without-f46f72f58c4b
- https://medium.com/@michalkowal567/rules-pattern-in-golang-425765f3c646
- https://levelup.gitconnected.com/practical-solid-in-golang-open-closed-principle-1dd361565452