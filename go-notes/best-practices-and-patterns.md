# Best Practices and Patterns

</br>

## List of Contents:
### 1. [7 Code Patterns in Go I Can’t Live Without](#content-1)
### 1. [Rules pattern in Golang](#content-2)


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

## References:
- https://betterprogramming.pub/7-code-patterns-in-go-i-cant-live-without-f46f72f58c4b
- https://medium.com/@michalkowal567/rules-pattern-in-golang-425765f3c646