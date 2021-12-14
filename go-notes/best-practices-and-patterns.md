# Best Practices and Patterns

</br>

## List of Contents:
### 1. [7 Code Patterns in Go I Can’t Live Without](#content-1)
### 2. [Rules pattern in Golang](#content-2)
### 3. [Practical SOLID in Golang: Single Responsibility Principle](#content-3)
### 4. [Practical SOLID in Golang: Open/Closed Principle](#content-4)
### 5. [Practical SOLID in Golang: Liskov Substitution Principle](#content-5)
### 6. [Practical SOLID in Golang: Interface Segregation Principle](#content-6)
### 7. [Practical SOLID in Golang: Dependency Inversion Principle](#content-7)


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

## [Practical SOLID in Golang: Single Responsibility Principle](https://levelup.gitconnected.com/practical-solid-in-golang-single-responsibility-principle-20afb8643483) <span id="content-3"></span>


### When we do not respect Single Responsibility
- The Single Responsibility Principle (SRP) states that each software module should have one and only one reason to change.
- Today, SRP has a wide range, where it touches different aspects of the software. We can use its purpose in classes, functions, modules. And, naturally, in Go, we can use it in a struct.
  ```go
  type EmailService struct {
    db           *gorm.DB
    smtpHost     string
    smtpPassword string
    smtpPort     int
  }

  func NewEmailService(db *gorm.DB, smtpHost string, smtpPassword string, smtpPort int) *EmailService {
    return &EmailService{
      db:           db,
      smtpHost:     smtpHost,
      smtpPassword: smtpPassword,
      smtpPort:     smtpPort,
    }
  }

  func (s *EmailService) Send(from string, to string, subject string, message string) error {
    email := EmailGorm{
      From:    from,
      To:      to,
      Subject: subject,
      Message: message,
    }

    err := s.db.Create(&email).Error
    if err != nil {
      log.Println(err)
      return err
    }
    
    auth := smtp.PlainAuth("", from, s.smtpPassword, s.smtpHost)
    
    server := fmt.Sprintf("%s:%d", s.smtpHost, s.smtpPort)
    
    err = smtp.SendMail(server, auth, from, []string{to}, []byte(message))
    if err != nil {
      log.Println(err)
      return err
    }

    return nil
  }
  ```
- Explanation for the above code:
  - Even if it looks fine, we realize that this code breaks every aspect of SRP when we scratch the surface.
  - The responsibility of EmailService is not just to send emails but to store an email message into DB and send it via SMTP protocol.
  - The word "and" is bold with purpose. Using such an expression does not look like the case where we describe a single responsibility.
  - As soon as describing the responsibility of some code struct requires the usage of the word "and", it already breaks the Single Responsibility Principle.
  - In our example, we broke SRP on many code levels. The first one is on the level of function. Function Send is responsible for both storing a message in the database and sending email over SMTP protocol.
  - The second level is a struct EmailService. As we already concluded, it also has two responsibilities, storing inside the database and sending emails.
  - What are the consequences of such a code?
    - When we change a table structure or type of storage, we need to change a code for sending emails over SMTP.
    - When we want to integrate Mailgun or Mailjet, we need to change a code for storing data in the MySQL database.
    - If we choose different integration of sending emails in the application, each integration needs to have a logic to store data in the database.
    - Suppose we decide to split the application's responsibility into two teams, one for maintaining a database and the second one for integrating email providers. In that case, they will work on the same code.
    - This service is practically untestable with unit tests.

### How we do respect Single Responsibility
- To split the responsibilities in this case and make code blocks that have just one reason to exist, we should define a struct for each of them.
- It practically means to have a separate struct for storing data in some storage and a different struct for sending emails by using some integration with email providers. 
  ```go
  type EmailGorm struct {
    gorm.Model
    From    string
    To      string
    Subject string
    Message string
  }

  type EmailRepository interface {
    Save(from string, to string, subject string, message string) error
  }

  type EmailDBRepository struct {
    db *gorm.DB
  }

  func NewEmailRepository(db *gorm.DB) EmailRepository {
    return &EmailDBRepository{
      db: db,
    }
  }

  func (r *EmailDBRepository) Save(from string, to string, subject string, message string) error {
    email := EmailGorm{
      From:    from,
      To:      to,
      Subject: subject,
      Message: message,
    }

    err := r.db.Create(&email).Error
    if err != nil {
      log.Println(err)
      return err
    }

    return nil
  }

  type EmailSender interface {
    Send(from string, to string, subject string, message string) error
  }

  type EmailSMTPSender struct {
    smtpHost     string
    smtpPassword string
    smtpPort     int
  }

  func NewEmailSender(smtpHost string, smtpPassword string, smtpPort int) EmailSender {
    return &EmailSMTPSender{
      smtpHost:     smtpHost,
      smtpPassword: smtpPassword,
      smtpPort:     smtpPort,
    }
  }

  func (s *EmailSMTPSender) Send(from string, to string, subject string, message string) error {
    auth := smtp.PlainAuth("", from, s.smtpPassword, s.smtpHost)

    server := fmt.Sprintf("%s:%d", s.smtpHost, s.smtpPort)

    err := smtp.SendMail(server, auth, from, []string{to}, []byte(message))
    if err != nil {
      log.Println(err)
      return err
    }

    return nil
  }

  type EmailService struct {
    repository EmailRepository
    sender     EmailSender
  }

  func NewEmailService(repository EmailRepository, sender EmailSender) *EmailService {
    return &EmailService{
      repository: repository,
      sender:     sender,
    }
  }

  func (s *EmailService) Send(from string, to string, subject string, message string) error {
    err := s.repository.Save(from, to, subject, message)
    if err != nil {
      return err
    }

    return s.sender.Send(from, to, subject, message)
  }
  ```
- Here we provide two new structs. The first one is EmailDBRepository as an implementation for the EmailRepository interface. It includes support for persisting data in the underlying database.
- The second structure is EmailSMTPSender that implements the EmailSender interface. This struct is responsible for only email sending over SMPT protocol.
- Finally, the new EmailService contains interfaces from above and delegates the request for email sending.
- Here, that is not the case. EmailService does not hold the responsibility for storing and sending emails. It delegates them to the structs below. Its responsibility is to delegate requests for processing emails to the underlying services.
- There is a difference between holding and delegating responsibility. If an adaptation of a particular code can remove the whole purpose of responsibility, we talk about holding. If that responsibility still exists even after removing a specific code, then we talk about delegation.

### Some more examples
- Example:
  ```go
  import "github.com/dgrijalva/jwt-go"

  func extractUsername(header http.Header) string {
    raw := header.Get("Authorization")
    parser := &jwt.Parser{}
    token, _, err := parser.ParseUnverified(raw, jwt.MapClaims{})
    if err != nil {
      return ""
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
      return ""
    }

    return claims["username"].(string)
  }
  ```
- Instead of restructuring a way to describe the function's purpose, we should engage more in restructuring the method itself. Below you can find a proposal for a new code:
  ```go
  func extractUsername(header http.Header) string {
    raw := extractRawToken(header)
    claims := extractClaims(raw)
    if claims == nil {
      return ""
    }
    
    return claims["username"].(string)
  }

  func extractRawToken(header http.Header) string {
    return header.Get("Authorization")
  }

  func extractClaims(raw string) jwt.MapClaims {
    parser := &jwt.Parser{}
    token, _, err := parser.ParseUnverified(raw, jwt.MapClaims{})
    if err != nil {
      return nil
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
      return nil
    }
    
    return claims
  }
  ```
- Now we have two new functions. The first one, extractRawToken, contains the responsibility for extracting a raw JWT token from the HTTP header. If we change a key in the header that holds a token, we should touch just one method.
- Example 2:
  ```go
  type User struct {
    db *gorm.DB
    Username string
    Firstname string
    Lastname string
    Birthday time.Time
    //
    // some more fields
    //
  }

  func (u User) IsAdult() bool {
    return u.Birthday.AddDate(18, 0, 0).Before(time.Now())
  }

  func (u *User) Save() error {
    return u.db.Exec("INSERT INTO users ...", u.Username, u.Firstname, u.Lastname, u.Birthday).Error
  }
  ```
- The example above shows the typical implementation of the pattern Active Record. In this case, we also added a business logic inside the User struct, not just storing data into the database.

**[⬆ back to top](#list-of-contents)**

</br>

---

## [Practical SOLID in Golang: Open/Closed Principle](https://levelup.gitconnected.com/practical-solid-in-golang-open-closed-principle-1dd361565452) <span id="content-4"></span>

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

## [Practical SOLID in Golang: Liskov Substitution Principlet](https://levelup.gitconnected.com/practical-solid-in-golang-liskov-substitution-principle-e0d2eb9dd39) <span id="content-5"></span>

### When we do not respect The Liskov Substitution
- Definition:
  > Let Φ(x) be a property provable about objects x of type T. Then Φ(y) should be true for objects y of type S where S is a subtype of T.
- A more practical definition:
  > If S is a subtype of T, then objects of type T in a program may be replaced with objects of type S without altering any of the desirable properties of that program.
- If ObjectA is an instance of ClassA, and ObjectB is an instance of ClassB, and ClassB is a subtype of ClassA — if we use ObjectB instead ObjectA somewhere in the code, the functionality of the application must not be broken.
- Example:
  ```go
  type User struct {
    ID uuid.UUID
    //
    // some fields
    //
  }

  type UserRepository interface {
    Update(ctx context.Context, user User) error
  }

  type DBUserRepository struct {
    db *gorm.DB
  }

  func (r *DBUserRepository) Update(ctx context.Context, user User) error {
    return r.db.WithContext(ctx).Delete(user).Error
  }
  ```
- Here we can see one code example. And, to be honest, I hardly could find worse and more dummy than this one. Like, instead of updating the User in the database, like the Update method says, it deletes it.
- But, hey, this is the point. We can see an interface, UserRepository. Following the interface, we have a struct, DBUserRepository. Although this struct implements the initial interface — it does not do what the interface claims it should.
- It breaks the functionality of the interface instead of following the expectation. Here is the point of LSP in Go: struct must not violate the purpose of the interface.
- Less ridiculous example:
  ```go
  type UserRepository interface {
    Create(ctx context.Context, user User) (*User, error)
    Update(ctx context.Context, user User) error
  }

  type DBUserRepository struct {
    db *gorm.DB
  }

  func (r *DBUserRepository) Create(ctx context.Context, user User) (*User, error) {
    err := r.db.WithContext(ctx).Create(&user).Error
    return &user, err
  }

  func (r *DBUserRepository) Update(ctx context.Context, user User) error {
    return r.db.WithContext(ctx).Save(&user).Error
  }

  type MemoryUserRepository struct {
    users map[uuid.UUID]User
  }

  func (r *MemoryUserRepository) Create(_ context.Context, user User) (*User, error) {
    if r.users == nil {
      r.users = map[uuid.UUID]User{}
    }
    user.ID = uuid.New()
    r.users[user.ID] = user
    
    return &user, nil
  }

  func (r *MemoryUserRepository) Update(_ context.Context, user User) error {
    if r.users == nil {
      r.users = map[uuid.UUID]User{}
    }
    r.users[user.ID] = user

    return nil
  }
  ```
- Here we have new UserRepository interface and its two implementations: DBUserRepository and MemoryUserRepository. As we can see, MemoryUserRepository does need theContext argument, but it is still there to respect the interface.
- Here problem begins. We adapted MemoryUserRepository to support the interface, even though this intention is unnatural. Consequently, we can switch data sources in our application, where one source is not permanent storage.
- The purpose of the Repository pattern is to represent the interface to the underlying permanent data storage, like a database. It should not play the role of cache system, like here where we store Users in memory.
- Geometry example:
  ```go
  type ConvexQuadrilateral interface {
    GetArea() int
  }

  type Rectangle interface {
    ConvexQuadrilateral
    SetA(a int)
    SetB(b int)
  }

  type Oblong struct {
    Rectangle
    a int
    b int
  }

  func (o *Oblong) SetA(a int) {
    o.a = a
  }

  func (o *Oblong) SetB(b int) {
    o.b = b
  }

  func (o Oblong) GetArea() int {
    return o.a * o.b
  }

  type Square struct {
    Rectangle
    a int
  }

  func (o *Square) SetA(a int) {
    o.a = a
  }

  func (o Square) GetArea() int {
    return o.a * o.a
  }

  func (o *Square) SetB(b int) {
    //
    // should it be o.a = b ?
    // or should it be empty?
    //
  }
  ```
- This interface defines only one method, GetArea. As a subtype of ConvexQuadrilateral we can define an interface Rectangle. This subtype has two sides involving its area, so we must provide SetA and SetB.
- The next is actual implementations. The first one is Oblong, which can have wider width or wider height. In geometry, it is any rectangle that is not square. Implementing the logic for this struct is easy.
- The second subtype of Rectangle is Square. In geometry, a square is a subtype of a rectangle, but if we follow this in software development, we can only make an issue in our implementation.
- Square has all four equal sides. So, that makes SetB obsolete. To respect the subtyping we had chosen initially, we realized that our code has obsolete methods. The same issue is if we pick a slightly different path:
  ```go
  type ConvexQuadrilateral interface {
    GetArea() int
  }

  type EquilateralRectangle interface {
    ConvexQuadrilateral
    SetA(a int)
  }

  type Oblong struct {
    EquilateralRectangle
    a int
    b int
  }

  func (o *Oblong) SetA(a int) {
    o.a = a
  }

  func (o *Oblong) SetB(b int) {
    // where is this method defined?
    o.b = b
  }

  func (o Oblong) GetArea() int {
    return o.a * o.b
  }

  type Square struct {
    EquilateralRectangle
    a int
  }

  func (o *Square) SetA(a int) {
    o.a = a
  }

  func (o Square) GetArea() int {
    return o.a * o.a
  }
  ```
- In the example above, instead of Rectangle, we introduced the EquilateralRectangle interface. In geometry, that should be a rectangle with all four equal sides.
- In this case, when our interface defines only the SetA method, we avoided obsolete code in our implementation. Still, this breaks LSP, as we introduced an additional method SetB for Oblong, without which we can not calculate the area, even that our interface says we can.
- So, we already started catching the idea of The Liskov Substitution Principle in Go. So we can summarize what can go wrong if we break it:
  - It provides a false shortcut for implementation.
  - It can cause obsolete code.
  - It can damage the expected code execution.
  - It can break the desired use case.
  - It can cause an unmaintainable interface structure.

### How we do respect The Liskov Substitution
- So, let us first jump into fixing the issue for different implementations of the UserRepository interface:
  ```go
  type UserRepository interface {
    Create(ctx context.Context, user User) (*User, error)
    Update(ctx context.Context, user User) error
  }

  type MySQLUserRepository struct {
    db *gorm.DB
  }

  type CassandraUserRepository struct {
    session *gocql.Session
  }

  type UserCache interface {
    Create(user User)
    Update(user User)
  }

  type MemoryUserCache struct {
    users map[uuid.UUID]User
  }
  ```
- In this example, we split the interface into two, with clear purpose and signatures of different methods. Now, we have the UserRepository interface and the UserCache interface.
- UserRepository purpose is now definitely to permanently store user data into some storage. For it, we prepared concrete implementations like MySQLUserRepository and CassandraUserRepository.
- On the other hand, we have the UserCache interface with a clear understanding that we need it for temporarily keeping user data in some cache. As Concrete implementation, we can use MemoryUserCache.
- Example:
  ```go
  type ConvexQuadrilateral interface {
    GetArea() int
  }

  type EquilateralQuadrilateral interface {
    ConvexQuadrilateral
    SetA(a int)
  }

  type NonEquilateralQuadrilateral interface {
    ConvexQuadrilateral
    SetA(a int)
    SetB(b int)
  }

  type NonEquiangularQuadrilateral interface {
    ConvexQuadrilateral
    SetAngle(angle float64)
  }

  type Oblong struct {
    NonEquilateralQuadrilateral
    a int
    b int
  }

  type Square struct {
    EquilateralQuadrilateral
    a int
  }

  type Parallelogram struct {
    NonEquilateralQuadrilateral
    NonEquiangularQuadrilateral
    a     int
    b     int
    angle float64
  }

  type Rhombus struct {
    EquilateralQuadrilateral
    NonEquiangularQuadrilateral
    a     int
    angle float64
  }
  ```
- In this case, we introduced three new interfaces: EquilateralQuadrilateral (a quadrilateral with all four equal sides), NonEquilateralQuadrilateral (a quadrilateral with two pairs of equal sides), and NonEquiangularQuadrilateral (a quadrilateral with two pairs of equal angles).
- Now, we can define a Square, with only the SetA method, Oblong with both SetA and SetB, and Parallelogram with all of them plus SetAngle. So, we did not use subtyping here, but more like features.
- With both fixed examples, we restructured our code so it could always meet expectations from the end-user. It also removes obsolete methods and does not break any of them. The code is now stable.

### Conclusion
- The Liskov Substitution Principle teaches us what the correct way of subtyping is. We should never make forced polymorphism, even that it mirrors the real-world situtaions.


**[⬆ back to top](#list-of-contents)**

</br>

---

## [Practical SOLID in Golang: Interface Segregation Principle](https://betterprogramming.pub/7-code-patterns-in-go-i-cant-live-without-f46f72f58c4b) <span id="content-6"></span>

### When we do not respect The Interface Segregation
- Keep interfaces small so that users don’t end up depending on things they don’t need.
- Snippet:
  ```go
  type User interface {
    AddToShoppingCart(product Product)
    IsLoggedIn() bool
    Pay(money Money) error
    HasPremium() bool
    HasDiscountFor(product Product) bool
    //
    // some additional methods
    //
  }

  type Guest struct {
    cart ShoppingCart
    //
    // some additional fields
    //
  }

  func (g *Guest) AddToShoppingCart(product Product) {
    g.cart.Add(product)
  }

  func (g *Guest) IsLoggedIn() bool {
    return false
  }

  func (g *Guest) Pay(Money) error {
    return errors.New("user is not logged in")
  }

  func (g *Guest) HasPremium() bool {
    return false
  }

  func (g *Guest) HasDiscountFor(Product) bool {
    return false
  }

  type NormalCustomer struct {
    cart   ShoppingCart
    wallet Wallet
    //
    // some additional fields
    //
  }

  func (c *NormalCustomer) AddToShoppingCart(product Product) {
    c.cart.Add(product)
  }

  func (c *NormalCustomer) IsLoggedIn() bool {
    return true
  }

  func (c *NormalCustomer) Pay(money Money) error {
    return c.wallet.Deduct(money)
  }

  func (c *NormalCustomer) HasPremium() bool {
    return false
  }

  func (c *NormalCustomer) HasDiscountFor(Product) bool {
    return false
  }

  type PremiumCustomer struct {
    cart     ShoppingCart
    wallet   Wallet
    policies []DiscountPolicy
    //
    // some additional fields
    //
  }

  func (c *PremiumCustomer) AddToShoppingCart(product Product) {
    c.cart.Add(product)
  }

  func (c *PremiumCustomer) IsLoggedIn() bool {
    return true
  }

  func (c *PremiumCustomer) Pay(money Money) error {
    return c.wallet.Deduct(money)
  }

  func (c *PremiumCustomer) HasPremium() bool {
    return true
  }

  func (c *PremiumCustomer) HasDiscountFor(product Product) bool {
    for _, p := range c.policies {
      if p.IsApplicableFor(c, product) {
        return true
      }
    }
    
    return false
  }

  type UserService struct {
    //
    // some fields
    //
  }

  func (u *UserService) Checkout(ctx context.Context, user User, product Product) error {
    if !user.IsLoggedIn() {
      return errors.New("user is not logged in")	
    }
    
    var money Money
    //
    // some calculation
    //
    if user.HasDiscountFor(product) {
      //
      // apply discount
      //
    }
    
    return user.Pay(money)
  }
  ```
- The actual implementations for this interface are three structs. The first one is the Guest struct. It should be a User who is not logged in on our system, but at least they can add a Product to the ShoppingCart.
- The second implementation is the NormalCustomer. It can do whatever Guest can, plus to buy a Product.
- The third implementation is thePremiumCustomer, which can use all features of our system.
- Now, look at all of those three structs. Only PremiumCustomer requires all three methods.
- Maybe we can assign all of them to NormalCustomer, but definitely, we hardly need more than two for Guest.
- Methods HasPremium and HasDiscountFor do not have any sense for Guest. If that struct represents the User who is not logged in, why would we even consider implementing methods for discounts?
- And we did all of this to add generalization inside UserService to handle all types of Users in the same place, with the same code. But, because of that, we need to implement a bunch of unused methods.
- So, to have better generalization, we got:
  - Many structs with unused methods.
  - Methods that we need somehow to mark so that others do not use them.
  - Much additional code for unit testing.
  - Unnatural polymorphism.

### How we do respect The Interface Segregation
- Build interfaces around the minimal cohesive group of features.
- The only needed is to define a minimal interface that delivers a complete set of features.
- Snippet:
  ```go
  type User interface {
    AddToShoppingCart(product Product)
    //
    // some additional methods
    //
  }

  type LoggedInUser interface {
    User
    Pay(money Money) error
    //
    // some additional methods
    //
  }

  type PremiumUser interface {
    LoggedInUser
    HasDiscountFor(product Product) bool
    //
    // some additional methods
    //
  }

  type Guest struct {
    cart ShoppingCart
    //
    // some additional fields
    //
  }

  func (g *Guest) AddToShoppingCart(product Product) {
    g.cart.Add(product)
  }

  type NormalCustomer struct {
    cart   ShoppingCart
    wallet Wallet
    //
    // some additional fields
    //
  }

  func (c *NormalCustomer) AddToShoppingCart(product Product) {
    c.cart.Add(product)
  }

  func (c *NormalCustomer) Pay(money Money) error {
    return c.wallet.Deduct(money)
  }

  type PremiumCustomer struct {
    cart     ShoppingCart
    wallet   Wallet
    policies []DiscountPolicy
    //
    // some additional fields
    //
  }

  func (c *PremiumCustomer) AddToShoppingCart(product Product) {
    c.cart.Add(product)
  }

  func (c *PremiumCustomer) Pay(money Money) error {
    return c.wallet.Deduct(money)
  }

  func (c *PremiumCustomer) HasDiscountFor(product Product) bool {
    for _, p := range c.policies {
      if p.IsApplicableFor(c, product) {
        return true
      }
    }

    return false
  }

  type UserService struct {
    //
    // some fields
    //
  }

  func (u *UserService) Checkout(ctx context.Context, user User, product Product) error {
    loggedIn, ok := user.(LoggedInUser)
    if !ok {
      return errors.New("user is not logged in")
    }

    var money Money
    //
    // some calculation
    //
    if premium, ok := loggedIn.(PremiumUser); ok && premium.HasDiscountFor(product)  {
      //
      // apply discount
      //
    }

    return loggedIn.Pay(money)
  }
  ```
- Now, instead of one, we have three interfaces. PremiumUser embeds LoggedInUser, which embeds User. In addition, each of them introduces one method.
- The User now represents only customers who are still not authenticated on our platform. For such type, we know they can use features of ShoppingCart.
- The new LoggedInUser interface represents all our authenticated customers, and thePremiumUser interface represents all authenticated customers with a paid premium account.
- As you see in the UserService, instead of using methods with the boolean result, we just clarify the subtype of the User interface. If User implements LoggedInUser, we know that we talk about the authenticated customer.
- Also, if User implements PremiumUser, we know that we talk about the customer with a premium account. So, by casting, we already check for some business rules.

### Some more examples
- Adding small interfaces to make them more straightforward but still implementing them together in the same struct does not make too much sense.
- Snippet:
  ```go
  // too much splitting
  type UserWithFirstName interface {
    FirstName() string
  }

  type UserWithLastName interface {
    LastName() string
  }

  type UserWithFullName interface {
    FullName() string
  }

  // optimal splitting
  type UserWithName interface {
    FirstName() string
      LastName() string
      FullName() string
  }
  ```
- Obviously, if a customer is registered on our platform, they will need to provide their first and last name for billing purposes. So, our User will need both FirstName and LastName methods, and with that, naturally, FullName.
- A good example of one-method interface:
  ```go
  package io

  type Reader interface {
    Read(p []byte) (n int, err error)
  }

  type Writer interface {
    Write(p []byte) (n int, err error)
  }

  type Closer interface {
    Close() error
  }

  type Seeker interface {
    Seek(offset int64, whence int) (int64, error)
  }

  type WriteCloser interface {
    Writer
    Closer
  }

  type ReadWriteCloser interface {
    Reader
    Writer
    Closer
  }

  //.... and so on
  ```
- The Interface Segregation Principle is the fourth SOLID principle, and it stands for the letter I in the word SOLID. It teaches us to make our interfaces as tiny as possible.


**[⬆ back to top](#list-of-contents)**

</br>

---

## [Practical SOLID in Golang: Dependency Inversion Principle](https://levelup.gitconnected.com/practical-solid-in-golang-dependency-inversion-principle-8cbd4eed484b) <span id="content-7"></span>

### When we do not respect The Dependency Inversion
- High-level modules should not depend on low-level modules. Both should depend on abstractions. Abstractions should not depend on details. Details should depend on abstractions.
- First, we should accept Abstraction as an OOP concept. We use such a concept to expose essential behaviors and hide details of their implementation.
- High-level modules in the Go context are software components used on the top of the application, like code used for presentation. It can also be a code close to the top level, like code for business logic or some use case components.
- On the other hand, low-level software components are mostly small code pieces that support the higher level. They hide technical details about different infrastructural integrations. For example, that can be a struct that keeps the logic for retrieving data from the database, sending an SQS message, fetching a value from Redis, or sending an HTTP request to an external API.
- Snippet:
  ```go
  // infrastructure layer

  type UserRepository struct {
    db *gorm.DB
  }

  func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{
      db: db,
    }
  }

  func (r *UserRepository) GetByID(id uint) (*domain.User, error) {
    user := domain.User{}
    err := r.db.Where("id = ?", id).First(&user).Error
    if err != nil {
      return nil, err
    }

    return &user, nil
  }

  // domain layer

  type User struct {
    ID uint `gorm:"primaryKey;column:id"`
    // some fields
  }

  // application layer

  type EmailService struct {
    repository *infrastructure.UserRepository
    // some email sender
  }

  func NewEmailService(repository *infrastructure.UserRepository) *EmailService {
    return &EmailService{
      repository: repository,
    }
  }

  func (s *EmailService) SendRegistrationEmail(userID uint) error {
    user, err := s.repository.GetByID(userID)
    if err != nil {
      return err
    }
    // send email
    return nil
  }
  ```
- In the code snippet from above, we defined a high-level component, EmailService. This struct belongs to the application layer, and it is responsible for sending an email to fresh registered customers.
- So, it appears that our high-level component, EmailService, depends on the low-level component, UserRepository. Practically, without defining connection to the database, we can not initiate our use case struct.
- Unit test snippet:
  ```go
  import (
    "testing"
    // some dependencies
    "github.com/DATA-DOG/go-sqlmock"
    "github.com/stretchr/testify/assert"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
  )

  func TestEmailService_SendRegistrationEmail(t *testing.T) {
    db, mock, err := sqlmock.New()
    assert.NoError(t, err)

    dialector := mysql.New(mysql.Config{
      DSN:        "dummy",
      DriverName: "mysql",
      Conn:       db,
    })
    finalDB, err := gorm.Open(dialector, &gorm.Config{})
    
    repository := infrastructure.NewUserRepository(finalDB)
    service := NewEmailService(repository)
    //
    // a lot of code to define mocked SQL queries
    //
    // and then actual test
  }
  ```
- Mocking in Go relies on the usage of interfaces, for which we can define a mocked implementation, but we can not do the same for structs.
- So, we can not mock UserRepository, as it is a struct.
- But what about abstractions that rely on details? Let us check the code below:
  ```go
  // domain layer

  type User struct {
    ID uint `gorm:"primaryKey;column:id"`
    // some fields
  }

  type UserRepository interface {
    GetByID(id uint) (*User, error)
  }
  ```
- Look at the User struct. It still poses a definition for mapping to the database.
And, even if such a struct is inside the domain layer, it still possesses infrastructural details.
- Our new interface UserRepository (abstraction) depends on the User struct with the database scheme (details), and we still break DIP.
- Any change to the database engine or table structure affects our highest levels.

### How we do respect The Dependency Inversion
- High-level modules should not depend on low-level modules. Both should depend on abstractions. Abstractions should not depend on details. Details should depend on abstractions.
- We should define some abstraction (an interface) on which both of our components, EmailService and UserRepository, will depend.
- In addition, such abstraction should not rely on any technical details (like Gorm object).
- Snippet:
  ```go
  // infrastructure layer

  type UserGorm struct {
    // some fields
  }

  func (g UserGorm) ToUser() *domain.User {
    return &domain.User{
      // some fields
    }
  }

  type UserDatabaseRepository struct {
    db *gorm.DB
  }

  var _ domain.UserRepository = &UserDatabaseRepository{}

  /*
  type UserRedisRepository struct {
    
  }
  type UserCassandraRepository struct {
  }
  */

  func NewUserDatabaseRepository(db *gorm.DB) UserRepository {
    return &UserDatabaseRepository{
      db: db,
    }
  }

  func (r *UserDatabaseRepository) GetByID(id uint) (*domain.User, error) {
    user := UserGorm{}
    err := r.db.Where("id = ?", id).First(&user).Error
    if err != nil {
      return nil, err
    }

    return user.ToUser(), nil
  }

  // domain layer

  type User struct {
    // some fields
  }

  type UserRepository interface {
    GetByID(id uint) (*User, error)
  }

  // application layer

  type EmailService struct {
    repository domain.UserRepository
    // some email sender
  }

  func NewEmailService(repository domain.UserRepository) *EmailService {
    return &EmailService{
      repository: repository,
    }
  }

  func (s *EmailService) SendRegistrationEmail(userID uint) error {
    user, err := s.repository.GetByID(userID)
    if err != nil {
      return err
    }
    // send email
    return nil
  }
  ```
- In the new code structure, we can see the UserRepository interface as a component that depends on the User struct, and both of them are inside the domain layer.
- The User struct does not reflect the database scheme anymore, but we use the UserGorm struct for that. That struct is on the infrastructure layer. It provides a method ToUser that maps it to the actual User struct.
- In this scenario, we can use UserGorm as part of the details used inside UserDatabaseRepository, as actual implementation for UserRepository.
- Inside the infrastructure layer, we can define as many implementations for UserRepository, as we want. That can be UserFileRepository or UserCassandraRepository, for example.
- The high-level component (EmailService) depends on abstraction — it contains a field with the type UserRepository. Still, how does low-level component depends on abstraction?
- In Go, structs implement interfaces implicitly. That means we do not need to add a code where UserDatabaseRepository explicitly implements UserRepository, but we can add a check with a blank identifier.
- Unit test snippet:
  ```go
  import (
    "errors"
    "testing"
  )

  type GetByIDFunc func(id uint) (*User, error)

  func (f GetByIDFunc) GetByID(id uint) (*User, error) {
    return f(id)
  }

  func TestEmailService_SendRegistrationEmail(t *testing.T) {
    service := NewEmailService(GetByIDFunc(func(id uint) (*User, error) {
      return nil, errors.New("error")
    }))
    //
    // and just to call the service
  }
  ```

### Some more examples
- Example 1:
  ```go
  type User struct {
    // some fields
  }

  type UserJSON struct {
    // some fields
  }

  func (j UserJSON) ToUser() *User {
    return &User{
      // some fields
    }
  }

  func GetUser(id uint) (*User, error) {
    filename := fmt.Sprintf("user_%d.json", id)
    data, err := ioutil.ReadFile(filename)
    if err != nil {
      return nil, err
    }
    
    var user UserJSON
    err = json.Unmarshal(data, &user)
    if err != nil {
      return nil, err
    }
    
    return user.ToUser(), nil
  }
  ```
- Full snippet:
  ```go
  type User struct {
  // some fields
  }

  type UserJSON struct {
    // some fields
  }

  func (j UserJSON) ToUser() *User {
    return &User{
      // some fields
    }
  }

  func GetUserFile(id uint) (io.Reader, error) {
    filename := fmt.Sprintf("user_%d.json", id)
    file, err := os.Open(filename)
    if err != nil {
      return nil, err
    }

    return file, nil
  }

  func GetUserHTTP(id uint) (io.Reader, error) {
    uri := fmt.Sprintf("http://some-api.com/users/%d", id)
    resp, err := http.Get(uri)
    if err != nil {
      return nil, err
    }

    return resp.Body, nil
  }

  func GetDummyUser(userJSON UserJSON) (io.Reader, error) {
    data, err := json.Marshal(userJSON)
    if err != nil {
      return nil, err
    }

    return bytes.NewReader(data), nil
  }

  func GetUser(reader io.Reader) (*User, error) {
    data, err := ioutil.ReadAll(reader)
    if err != nil {
      return nil, err
    }

    var user UserJSON
    err = json.Unmarshal(data, &user)
    if err != nil {
      return nil, err
    }

    return user.ToUser(), nil
  }
  ```


**[⬆ back to top](#list-of-contents)**

</br>

---

## References:
- https://betterprogramming.pub/7-code-patterns-in-go-i-cant-live-without-f46f72f58c4b
- https://medium.com/@michalkowal567/rules-pattern-in-golang-425765f3c646
- https://levelup.gitconnected.com/practical-solid-in-golang-single-responsibility-principle-20afb8643483
- https://levelup.gitconnected.com/practical-solid-in-golang-open-closed-principle-1dd361565452
- https://levelup.gitconnected.com/practical-solid-in-golang-interface-segregation-principle-f272c2a9a270