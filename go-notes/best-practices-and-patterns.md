# Best Practices and Patterns

</br>

## [List of Contents:](#list-of-contents)
### 1. [7 Code Patterns in Go I Can’t Live Without](#content-1)
### 2. [Rules pattern in Golang](#content-2)
### 3. [Practical DDD in Golang: Repository](#content-3)
### 4. [Wrappers and decorators in Golang](#content-4)


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


## [Practical DDD in Golang: Repository](https://levelup.gitconnected.com/practical-ddd-in-golang-repository-d308c9d79ba7) <span id="content-3"></span>

### The Anti-Corruption Layer
- As the domain layer is the one on the bottom and does not communicate with others, we define the Repository inside it but as an interface.
- Example:
  ```go
  import (
    "context"

    "github.com/google/uuid"
  )

  type Customer struct {
    ID uuid.UUID
    //
    // some fields
    //
  }

  type Customers []Customer

  type CustomerRepository interface {
    GetCustomer(ctx context.Context, ID uuid.UUID) (*Customer, error)
    SearchCustomers(ctx context.Context, specification CustomerSpecification) (Customers, int, error)
    SaveCustomer(ctx context.Context, customer Customer) (*Customer, error)
    UpdateCustomer(ctx context.Context, customer Customer) (*Customer, error)
    DeleteCustomer(ctx context.Context, ID uuid.UUID) (*Customer, error)
  }
  ```
- That interface we call a Contract that defines the method signatures we can call inside our domain.
- As we defined Repository as such interface, we can use it everywhere inside the domain layer.
- It always expects and returns us our Entities, in this case, Customer and Customers (I like to define such particular collections in Go to attach different methods to them).
- The Entity Customer does not hold any information about the type of storage below: there is no Go tag defining JSON structure, Gorm columns, or anything similar. For that, we must use the infrastructure layer.
  ```go
  // domain layer

  type CustomerRepository interface {
    GetCustomer(ctx context.Context, ID uuid.UUID) (*Customer, error)
    SearchCustomers(ctx context.Context, specification CustomerSpecification) (Customers, int, error)
    SaveCustomer(ctx context.Context, customer Customer) (*Customer, error)
    UpdateCustomer(ctx context.Context, customer Customer) (*Customer, error)
    DeleteCustomer(ctx context.Context, ID uuid.UUID) (*Customer, error)
  }

  // infrastructure layer

  import (
    "context"

    "github.com/google/uuid"
    "gorm.io/gorm"
  )

  type CustomerGorm struct {
    ID   uint   `gorm:"primaryKey;column:id"`
    UUID string `gorm:"uniqueIndex;column:uuid"`
    //
    // some fields
    //
  }

  func (c CustomerGorm) ToEntity() (model.Customer, error) {
    parsed, err := uuid.Parse(c.UUID)
    if err != nil {
      return Customer{}, err
    }
    
    return model.Customer{
      ID: parsed,
      //
      // some fields
      //
    }, nil
  }

  type CustomerRepository struct {
    connection *gorm.DB
  }

  func (r *CustomerRepository) GetCustomer(ctx context.Context, ID uuid.UUID) (*model.Customer, error) {
    var row CustomerGorm
    err := r.connection.WithContext(ctx).Where("uuid = ?", ID).First(&row).Error
    if err != nil {
      return nil, err
    }
    
    customer, err := row.ToEntity()
    if err != nil {
      return nil, err
    }
    
    return &customer, nil
  }
  //
  // other methods
  //
  ```
- In the example, you see two different structures, Customer and CustomerGorm. The first one is Entity, where we want to keep our business logic, some domain invariants, and rules. It does not know anything about the underlying database.
- The second structure is a Data Transfer Object, which defines how our data is transferred from and to storage. This structure does not have any other responsibility but to map the database’s data to our Entity.
- The division of those two structures is the fundamental point for using Repository as Anti-Corruption layer in our application. It makes sure that technical details of table structure do not pollute our business logic.
- What are the consequences here? First, it is the truth that we need to maintain two types of structures, one for business logic, one for storage. 
- In addition, I insert the third structure as well, the one I use as Data Transfer Object for my API.
- Example:
  ```go
  // domain layer

  type Customer struct {
    ID      uuid.UUID
    Person  *Person
    Company *Company
    Address Address
  }

  type Person struct {
    SSN       string
    FirstName string
    LastName  string
    Birthday  Birthday
  }

  type Birthday time.Time

  type Company struct {
    Name               string
    RegistrationNumber string
    RegistrationDate   time.Time
  }

  type Address struct {
    Street   string
    Number   string
    Postcode string
    City     string
  }

  // infrastructure layer

  type CustomerGorm struct {
    ID        uint         `gorm:"primaryKey;column:id"`
    UUID      string       `gorm:"uniqueIndex;column:id"`
    PersonID  uint         `gorm:"column:person_id"`
    Person    *PersonGorm  `gorm:"foreignKey:PersonID"`
    CompanyID uint         `gorm:"column:company_id"`
    Company   *CompanyGorm `gorm:"foreignKey:CompanyID"`
    Street    string       `gorm:"column:street"`
    Number    string       `gorm:"column:number"`
    Postcode  string       `gorm:"column:postcode"`
    City      string       `gorm:"column:city"`
  }

  func (c CustomerGorm) ToEntity() (model.Customer, error) {
    parsed, err := uuid.Parse(c.UUID)
    if err != nil {
      return model.Customer{}, err
    }

    return model.Customer{
      ID:      parsed,
      Person:  c.Person.ToEntity(),
      Company: c.Company.ToEntity(),
      Address: Address{
        Street:   c.Street,
        Number:   c.Number,
        Postcode: c.Postcode,
        City:     c.City,
      },
    }, nil
  }

  type PersonGorm struct {
    ID        uint      `gorm:"primaryKey;column:id"`
    SSN       string    `gorm:"uniqueIndex;column:ssn"`
    FirstName string    `gorm:"column:first_name"`
    LastName  string    `gorm:"column:last_name"`
    Birthday  time.Time `gorm:"column:birthday"`
  }

  func (p *PersonGorm) ToEntity() *model.Person {
    if p == nil {
      return nil
    }

    return &model.Person{
      SSN:       p.SSN,
      FirstName: p.FirstName,
      LastName:  p.LastName,
      Birthday:  Birthday(p.Birthday),
    }
  }

  type CompanyGorm struct {
    ID                 uint      `gorm:"primaryKey;column:id"`
    Name               string    `gorm:"column:name"`
    RegistrationNumber string    `gorm:"column:registration_number"`
    RegistrationDate   time.Time `gorm:"column:registration_date"`
  }

  func (c *CompanyGorm) ToEntity() *model.Company {
    if c == nil {
      return nil
    }

    return &model.Company{
      Name:               c.Name,
      RegistrationNumber: c.RegistrationNumber,
      RegistrationDate:   c.RegistrationDate,
    }
  }

  func NewRow(customer model.Customer) CustomerGorm {
    var person *PersonGorm
    if customer.Person != nil {
      person = &PersonGorm{
        SSN:       customer.Person.SSN,
        FirstName: customer.Person.FirstName,
        LastName:  customer.Person.LastName,
        Birthday:  time.Time(customer.Person.Birthday),
      }
    }

    var company *CompanyGorm
    if customer.Company != nil {
      company = &CompanyGorm{
        Name:               customer.Company.Name,
        RegistrationNumber: customer.Company.RegistrationNumber,
        RegistrationDate:   customer.Company.RegistrationDate,
      }
    }

    return CustomerGorm{
      UUID:     uuid.NewString(),
      Person:   person,
      Company:  company,
      Street:   customer.Address.Street,
      Number:   customer.Address.Number,
      Postcode: customer.Address.Postcode,
      City:     customer.Address.City,
    }
  }
  ```
- Still, besides whole this maintenance, it brings new value to our code. We can provide our Entities inside the domain layer in a way that describes our business logic the best. We do not limit them with the storage we use.

### Persistence
- The second feature of the Repository is Persistence. We define the logic for sending our data into the storage below to keep it permanently, update, or even delete.
- Example:
  ```go
  func NewRow(customer Customer) CustomerGorm {
    return CustomerGorm{
      UUID: uuid.NewString(),
      //
      // some fields
      //
    }
  }

  type CustomerRepository struct {
    connection *gorm.DB
  }

  func (r *CustomerRepository) SaveCustomer(ctx context.Context, customer Customer) (*Customer, error) {
    row := NewRow(customer)
    err := r.connection.WithContext(ctx).Save(&row).Error
    if err != nil {
      return nil, err
    }

    customer, err = row.ToEntity()
    if err != nil {
      return nil, err
    }

    return &customer, nil
  }
  //
  // other methods
  //
  ```
- Sometimes we decide to have unique identifiers that we want to create within an application. In such cases, the Repository is the right place. In the example above, you can see that we generate a new UUID before creating the database record.
- We can do this with integers if we want to avoid auto-incrementing from the database engine. In any case, if we do not wish to rely on database keys, we should create them inside Repository.
  ```go
  type CustomerRepository struct {
    connection *gorm.DB
  }

  func (r *CustomerRepository) CreateCustomer(ctx context.Context, customer Customer) (*Customer, error) {
    tx := r.connection.Begin()
    defer func() {
      if r := recover(); r != nil {
        tx.Rollback()
      }
    }()

    if err := tx.Error; err != nil {
      return nil, err
    }

    //
    // some code
    //

    var total int64
    var err error
    if customer.Person != nil {
      err = tx.Model(PersonGorm{}).Where("ssn = ?", customer.Person.SSN).Count(&total).Error
    } else if customer.Person != nil {
      err = tx.Model(CompanyGorm{}).Where("registration_number = ?", customer.Person.SSN).Count(&total).Error
    }
    if err != nil {
      tx.Rollback()
      return nil, err
    } else if total > 0 {
      tx.Rollback()
      return nil, errors.New("there is already such record in DB")
    }
    
    //
    // some code
    //
    
    err = tx.Save(&row).Error
    if err != nil {
      tx.Rollback()
      return nil, err
    }

    err = tx.Commit().Error
    if err != nil {
      tx.Rollback()
      return nil, err
    }

    customer := row.ToEntity()

    return &customer, nil
  }
  ```
- The other thing we want the Repository to use for is transactions. Whenever we want to persist some data and execute many queries that work on the same extensive set of tables, it is an excellent time to define a transaction, which we should deliver inside the Repository.
- In the example from above, we are checking the uniqueness of Person or Company. If they exist, we return an error. All of that we can define as part of a single transaction, and if something fails there, we can roll back it.
- Here Repository is a perfect place for such code. It is good that we can also make our inserts more straightforward in the future so that we will not need transactions at all. In that case, we do not change a Contract of the Repository, but only the code inside.

### Types of Repositories
- As mentioned, we can use MongoDB or Cassandra. We can use a Repository for keeping our cache, and in that case, it would be Redis, for example. It can even be REST API or configurational file.
- Example:
  ```go
  // redis repository

  type CustomerRepository struct {
    client *redis.Client
  }

  func (r *CustomerRepository) GetCustomer(ctx context.Context, ID uuid.UUID) (*Customer, error) {
    data, err := r.client.Get(ctx, fmt.Sprintf("user-%s", ID.String())).Result()
    if err != nil {
      return nil, err
    }

    var row CustomerJSON
    err = json.Unmarshal([]byte(data), &row)
    if err != nil {
      return nil, err
    }
    
    customer := row.ToEntity()

    return &customer, nil
  }

  // API

  type CustomerRepository struct {
    client *http.Client
    baseUrl string
  }

  func (r *CustomerRepository) GetCustomer(ctx context.Context, ID uuid.UUID) (*Customer, error) {
    resp, err := r.client.Get(path.Join(r.baseUrl, "users", ID.String()))
    if err != nil {
      return nil, err
    }
    
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      return nil, err
    }
    defer resp.Body.Close()

    var row CustomerJSON
    err = json.Unmarshal(data, &row)
    if err != nil {
      return nil, err
    }

    customer := row.ToEntity()

    return &customer, nil
  }
  ```
- Now we can see the real benefit of having a split between our business logic and technical details. We keep the same interface for our Repository, so our domain layer can always use it.
- So, your Repository Contract should always deal with your business logic, but your Repository implementation must use internal structures that you can map later to Entities.


**[⬆ back to top](#list-of-contents)**

</br>

---

## [Wrappers and decorators in Golang](https://levelup.gitconnected.com/wrappers-and-decorators-in-golang-c8cbe1359932) <span id="content-4"></span>


### Introduction
- The most basic form of function wrapping is passing the function you want to wrap to another function as parameter.
- Example:
  ```go
  package main
  import (
    "fmt"
  )
  func main() {
    print(1, 2, sum)
  }
  func sum(a, b int) int {
    return a + b
  }
  func print(a, b int, f func(int, int) int) {
    fmt.Printf("inputs: %d, %d\n", a, b)
    res := f(a, b)
    fmt.Printf("result: %d\n", res)
  }
  ```
- Another option is to use closures to implement similar functionality. Closure functions are function created within another function’s scope.
- Example:
  ```go
  package main
  import (
  "fmt"
  )
  func main() {
  a := 1
  b := 2
  fmt.Println("Sum:")
  print(a, b, sum)
  fmt.Println("\nMultiply:")
  print(a, b, multiply)
  msg:="variable from outer scope"
  s := func(x, y int) int {
    fmt.Println(msg)
    return x - y
  }
  fmt.Println("\nSubtract:")
  print(a, b, s)
  }
  func sum(a, b int) int {
  return a + b
  }
  func multiply(a, b int) int {
  return a * b
  }
  func print(a, b int, f func(int, int) int) {
  fmt.Printf("inputs: %d, %d\n", a, b)
  res := f(a, b)
  fmt.Printf("result: %d\n", res)
  }
  ```

### Dependency injection
- In cases where we need to receive an extra dependency we can wrap the function we want to execute in another function that receives the dependency as a parameter, returning a function with the exact signature we need.

### The decorator pattern
- Example:
  ```go
  package cmd

  import (
    "log"
    "os"

    "github.com/RicardoLinck/decorators/cache"
    "github.com/RicardoLinck/decorators/service"
  )

  type runner interface {
    GetData(keys ...string) error
  }

  type defaultRunner struct {
    url string
  }

  func (d *defaultRunner) GetData(keys ...string) error {
    c := service.NewClient(d.url)

    cc := cache.NewCachedDataGetter(c)
    for _, k := range keys {
      log.Print(cc.GetData(k))
    }

    return nil
  }

  type dryRunner struct {
    runner
  }

  func (d *dryRunner) GetData(keys ...string) error {
    log.Default().SetOutput(os.Stdout)
    return d.runner.GetData(keys...)
  }

  type fileRunner struct {
    runner
    filePath string
  }

  func (f *fileRunner) GetData(keys ...string) error {
    file, err := os.Create(f.filePath)
    if err != nil {
      return err
    }

    log.Default().SetOutput(file)
    return f.runner.GetData(keys...)
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