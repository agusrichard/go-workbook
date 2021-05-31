# Coding Conventions and Guidelines


## Intro
Coding standards are a series of procedures for a particular programming language specifying a programming style, the methods, and different procedures.
A coding standard makes sure that all the developers working on the project are following certain specified guidelines. The code can be easily understood and proper consistency is maintained.
> **The finished program code should look like that it has been written by a single developer, in a single session.**

### In general, best practices of better code:
- Code comments and proper documentation </br>
  It is advisable to start every method or routine with the comment specifying what a routine, method or a function does, about its various parameters, its return values, errors and exceptions (if any).
- Use indentation </br>
  There is no particular style to be followed in general. We can make our own style, but make sure that we follow this style consistently.
- Avoid commenting on obvious things. </br>
  Comments are good, but too much of them will lead us to chaos. We should avoid write comments on obvious things, or if our code is self-explained then no need to add comments.
- Grouping code </br>
  Remember that high cohesion code is good. Small and simple code is easy to manage and considered to be better.
- Proper and consistent scheme for naming </br>
  Choose one naming scheme and stick with it for the rest of the project.
- Principles of DRY </br>
  The good practice is to write our own code and don't copy-paste too frequently.
- Avoid deep nesting structure </br>
  Complex logic can be considered as lack of clarity and hard to understand.
- Use short line length </br>
  This makes sure our code easy and comfortable to read.
- Proper organizations of files and folders.
- Refactoring code </br>
  We have to implement the Open/Close Principle which basically states our code is closed for modification but open for extension.
  

## Project Structure

There are several goals we should achieve when structuring our project:
- Consistent
- Easy to understand, navigate, reason about (**make sense**)
- Easy to change
- Loosely coupled
- Easy to test
- As simple as possible, but no simplistic
- Design reflects exactly how the software works
- Structure reflects the design exactly

Then, the (kinda) best way to structure our project is by following Group by functions (layered architecure).
This project structure pattern is suitable for medium to large applications without having to overkill and use more complex structure.

We divide each package into its own functionalities, e.g. `package handlers` is responsible for HTTP handler for all endpoints.

### How to implement the Structure:

The list of packages (which means they have their own folders)
- handlers </br>
  We put all HTTP handler methods inside this folder. Then divide it into its own file. For example, all handlers for `vessels` functionality have to be put in this file which means we have to create `vessels.go` file.
  The main task of `handlers` is to parse the data for further processing and tidy up the response, so it will be readable for the user.
- usecases </br>
  All business logic implementations go here. Which means, this is the package where the data from user is processed.
  This package might include processing step before the data goes to the database or aggregate the queried data or combine several functionalities into one bucket.
  The bottom line is we have to put the business logic here and nowhere else.
- repositories </br>
  This package serves as the layer where our application makes contact with the database.
  Its main task is only database related functionalities, such as querying, inserting, updating, and deleting.
  Don't put excessive logic in here (e.g., control flow or conditional statements).
  If it is related to the database, like error handling for database querying, then it is fine.
- models </br>
  Any data entities or blueprints are belong in here. For example, if we have user data, then we have to make its blueprint as `struct` (with optional tags).
  We are also allowed to write helper methods, related to the entity. (e.g., if user need to be validated then it makes sense if the struct has method Validate() or IsValid())
- utilities </br>
  Any additional helpers are stored in this package.
- docs </br>
  For documentations (obviously).
  
If we need new packages to comprehend the growing complexity of our codebase, then feel free to add new one.
This structure should make our life easy, not giving us another headaches.

This is how our project structure would look like (example): </br>
```go
.
├── config
│   ├── config.go
│   └── db.go
├── docker-compose.yaml
├── Dockerfile
├── entities
│   ├── config.entities.go
│   ├── response.entities.go
│   └── tweet.entities.go
├── go.mod
├── go.sum
├── handlers
│   ├── tweet_handler.go
│   └── tweet_handler_mocked_test.go
├── main.go
├── mocks
│   ├── TweetRepository.go
│   └── TweetUsecase.go
├── postgres.Dockerfile
├── README.md
├── repositories
│   ├── tweet_repository.go
│   └── tweet_repository_test.go
├── sample.env
├── server
│   ├── handlers.go
│   ├── repositories.go
│   ├── setup.go
│   └── usecases.go
├── sql
│   └── tweet.sql
├── tmp
│   ├── air.log
│   └── main
├── usecases
│   ├── tweet_usecase.go
│   └── tweet_usecase_test.go
└── utils
├── server_http.go
└── truncate_table.go
```

## Naming Conventions

> Good naming is like a good joke. If you have to explain it, it’s not funny.
— Dave Cheney

### 1. Folder name
- Use plural noun to name a directory or folder, if this folder contains several files **with the same functionalities**.
  (e.g. handlers, usecases, repositories since these folders contain their own funcionalities)
- Use singluar noun to name a directory or folder, if this folder contains several files **with the name describes how the files are grouped beyond their type**.
  (e.g. folder post contain sample data, functions, tupes and etc)
  
### 2. Package name
- Lowercase only </br>
  Package names should be lowercase. Don’t use snake_case or camelCase in package names.
- Short, but representative names </br>
  Package names should be short, but should be unique and representative. Users of the package should be able to grasp its purpose from just the package’s name.
- Singular
  We should not use plural noun for the name of a package
- Avoid generic names </br>
  Avoid using package names of util, misc, common since these names could not explain what they contain in a decent way.
  
Example:
```go
package http_utils // BAD
package httpUtils // BAD
package httputils // BAD
package httputil // GOOD
```

### 3. File name
- Go follows a convention where source files are all lower case with underscore separating multiple words.
- Compound file names are separated with _
- File names that begin with “.” or “_” are ignored by the go tool
- Files with the suffix _test.go are only compiled and run by the go test tool.

### 4. Variables
- Common variable/type combinations may use really short names </br>
  Prefer i to index. </br>
  Prefer r to reader. </br>
  Prefer b to buffer. </br>
  ```go
    for i, v := range lists {
        ...
    }
  ```

- Avoid using global variables, it's better to create global Getter and Setter </br>
  - Bad practice:
  ```go
    // Inside database.go
    ...
  
    var DB *sql.DB
  
    ... (some mutation to variable DB)
  
  ```
  - Better way
  ```go
    var db *sql.DB  

    // Better way (with local db variable)
    func GetDB() *sql.DB {
        return db
    }
  ```

### 5. Functions
- Use MixedCase(CamelCase)
- If we don't need to expose this function to other packages, then keep it local.


### 6. Parameters
- If the type could explain what it is, then use short name
```go
func UpdateUser(u *models.User) error {...}
```
- If the type could not explain what it is, then use descriptive name
```go
func Sleep(duration time.Time) error {...}
```
- The key point is use short and descriptive name

### 7. Receivers
- By convention, they are one or two characters that reflect the receiver type,
because they typically appear on almost every line </br>
```go
func (b *Buffer) Read(p []byte) (n int, err error)
```
- Keep the receiver name short, since the type of the receiver is oftenly could explain itself.

### 8. Return values
Return values on exported functions should only be named for documentation purposes.
```go
func Copy(dst Writer, src Reader) (written int64, err error)

func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)
```

### 9. Exported package-level names
Don't use stutter naming such as strings.StringReader. That's why we have bytes.Buffer and strings.Reader,
not bytes.ByteBuffer and strings.StringReader.

### 10. Interface Types

Interfaces that specify just one method are usually just that function name with 'er' appended to it.
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

### 10.  Errors
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

More about errors on later section

### 11. Constants
Constant should use all capital letters and use underscore _ to separate words.

## Error Handling

> Don’t just check errors, handle them gracefully

Three way to handle errors:
### 1. Sentinel errors
```go
if err == ErrSomething { … }
```

With this strategy we have to evaluate the type of the error, or the string representation of the error.
Certainly not a good way since we are comparing the content of the error.

### 2. Error types (custom error)
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

With error types, we can pass another properties which an error probably has. The cons are we need to import the error types resulting in avoidable dependencies and complex error handling.

### 3. Opaque Errors

With opaque error, we just have to return it to the caller without assuming what it is.
The bad side is we don't know where the error starts, which resulting in tedious debugging sesion.

As a work around, we have to use errors package. We just need two functions, they are Wrap and Cause.

The main task of errors.Wrap is to wrap our error and add another additional context (resulting in error stack trace).
Meanwhile, errors.Cause is reverse operation or Wrap, if we want to access the error we can use errors.Cause

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
                errors.Println(err)
                os.Exit(1)
        }
}
```

### Practices about Error Handling:
1. Only handle errors once </br>
If we make less than one decision, we’re ignoring the error. As we see here, the error from w.Write is being discarded.
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
2. Handle and write error decently </br>
Start by using Opaque error when writing and error, then handle this error as the above code suggests.
If we think error types or sentinel errors are better, then go for it by keep following the good practice.
   
## Testing

> A code that cannot be tested is flawed

### 1. Create test file in the same folder containing our main program/logic </br>
If we create a file `vessel.go` in a package `repository`. Then the name of the test file is `vessel_test.go`.
Both `vessel.go` and `vessel_test.go` are in the same folder (in this case `repositories`).

### 2. Naming our test function </br>
The rule of thumb of naming test function is Test + [function name] + _ + [short description about testcase] + _ + [Positive/Negative Test]

Positive test cases ensure that users can perform appropriate actions when using valid data.

Negative test cases are performed to try to “break” the software by performing invalid (or unacceptable) actions, or by using invalid data.

```go
// file_test.go

func TestCreateUser_ValidInput_Positive(t *testing.T) {...}
```

For test, it seemed like we broke our own rule defined previously. But the reason for this naming scheme is to add verbosity when running the tests.
An example of exception of a rule.

### 3. Using assertions
If we don't need setup or teardown functions when running the tests, then just use normal testing function and assertions provided by testify.

```go
// file_test.go

func TestCreateUser_ValidInput_Positive(t *testing.T) {
	assert.Equal(t, 10, 10, 'create a verbose test message here')
	
	// or
	assert := assert.New(t)
	assert.Equal(10, 10, 'create a verbose test message here')
}
```

### 4. Using suites
We need suites if we have to run our test cases as a group. For example, when we need setup and teardown functions for our tests.
```go
type tweetRepositorySuite struct {
	suite.Suite
	repository TweetRepository
	cleanupExecutor utils.TruncateTableExecutor
}

// This function will only run once, before all tests
func (suite *tweetRepositorySuite) SetupSuite() {
	configs := config.GetConfig()
	db := config.ConnectDB(configs)
	repository := InitializeTweetRepository(db)

	suite.repository = repository

	suite.cleanupExecutor = utils.InitTruncateTableExecutor(db)
}

// This function will run after each test, we can use this for a clean up process
func (suite *tweetRepositorySuite) TearDownTest() {
	defer suite.cleanupExecutor.TruncateTable([]string{"tweets"})
}

// This function will run before each test
func (suite *tweetRepositorySuite) SetupTest() {...}

// This function will run after all tests in a suite have run
func (suite *tweetRepositorySuite) TearDownSuite() {...}

// Example of using suite to assert something
func (suite *tweetRepositorySuite) TestCreateTweet_Positive() {
  tweet := entities.Tweet{
    Username: "username",
    Text: "text",
  }
  
  err := suite.repository.CreateTweet(&tweet)
  suite.NoError(err, "no error when create tweet with valid input")
}

// The main function to run all our tests
func TestTweetRepository(t *testing.T) {
  suite.Run(t, new(tweetRepositorySuite))
}
```

### 5. Using mocks

It's better for us to mock interface rather than concrete struct (with some methods).
First, if we do this then we implement the 5th principle of SOLID, which is Dependency Inversion Principle.
Second, it's easier to mock interface. By using mockery, we can generate the mocked methods, without having to define them manually.

For example we have this interface:
```go
type TweetRepository interface {
	GetAllTweets() (*[]entities.Tweet, error)
	GetTweetByID(id int) (*entities.Tweet, error)
	SearchTweetByText(text string) (*[]entities.Tweet, error)
	CreateTweet(tweet *entities.Tweet) error
	UpdateTweet(tweet *entities.Tweet) error
	DeleteTweet(id int) error
}
```

To generate mocked methods of this interface, we just have to run `mockery --name=InterfaceName --recursive=true`.
For example, `mockery --name=TweetUsecase --recursive=true`

```go
// file_test.go
type tweetUsecaseSuite struct {
  suite.Suite
  repository *mocks.TweetRepository
  usecase TweetUsecase
  cleanupExecutor utils.TruncateTableExecutor
}

func (suite *tweetUsecaseSuite) SetupTest() {
  repository := new(mocks.TweetRepository)
  usecase := InitializeTweetUsecase(repository)
  
  
  suite.repository = repository
  suite.usecase = usecase
}

func (suite *tweetUsecaseSuite) TestCreateTweet_Positive() {
  tweet := entities.Tweet{
    Username: "username",
    Text: "text",
  }
  
  suite.repository.On("CreateTweet", &tweet).Return(nil)
  
  err := suite.usecase.CreateTweet(&tweet)
  suite.Nil(err, "err is a nil pointer so no error in this process")
  suite.repository.AssertExpectations(suite.T())
}

func TestTweetUsecase(t *testing.T) {
  suite.Run(t, new(tweetUsecaseSuite))
}
```

## Additional Guidelines
### 1. Use pointer when passing around struct between functions. </br>
By doing this, we can't check nullity of a struct easily. We don't have to compare it to empty struct of some type.
What we have to do is to check whether it returns pointer or nil.
```go
func someFunction(inp *input) *output {...}

func otherFunction() {
  result := someFunction()
  
  // No need for doing this, since result is a pointer 
  if result == output{} {...} // Note that this will result in error type checking

  // We can't access properties of a nil pointer
  if result == nil {...} 
  
  // we can access the properties of a struct here
  if result != nil {...}
}
```
### 2. Always check nullity if some function returns pointer.
If we directly assume some structs have some properties, then directly access them.
```go
type output struct {
  prop1 int
  prop2 string
}

func someFunction(inp *input) *output {
  return nil
}

func otherFunction() {
  result := someFunction()
  
  // this will result in error nil pointer dereference, because we try to access nil pointer
  fmt.Println(result.prop1)
  
  // therefore we have to check for the nullity first
  if result == nil {
    // can't access the result's properties here	
  } else { 
    // we can access the result's properties here
  }
  
  // 
}
```
### 3. Use defer as soon as something needs to be done when the function stack returns.

```go
func AmazingFunction() {
  result, err := SomeFunction()
  if err != nil {
    // error handling
  }
  defer result.Close()  // in this example, the function call returns, we close the result (it could be connection, file or any object which has method Close()
  
  // ... do something
}
```

Three rules for defer statement:
1. A deferred function's arguments are evaluated when the defer statement is evaluated.
```go
func a() {
    i := 0
    defer fmt.Println(i)
    i++
    return
}
```

This returns 0 instead of 1. Because defer is evaluated right away, but is called later on when the function returns.

2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.
```go
func b() {
    for i := 0; i < 4; i++ {
        defer fmt.Print(i)
    }
}
```

This prints "3210"

3. Deferred functions may read and assign to the returning function's named return values.
```go
func c() (i int) {
    defer func() { i++ }()
    return 1
}
```


## References:
### 1. Coding standards
  - https://www.multidots.com/importance-of-code-quality-and-coding-standard-in-software-development/#:~:text=Coding%20standards%20help%20in%20the,and%20thereby%20reduce%20the%20errors.&text=If%20the%20coding%20standards%20are%20followed%2C%20the%20code%20is%20consistent,at%20any%20point%20in%20time.
  - https://medium.com/leafgrowio-engineering/why-is-coding-standards-important-319fce79d1a4
### 2. Project structure in go
  - https://www.youtube.com/watch?v=oL6JBUk6tj0
  - https://tutorialedge.net/golang/go-project-structure-best-practices/
  - https://www.wolfe.id.au/2020/03/10/how-do-i-structure-my-go-project/
### 3. Naming Conventions
  - https://betterprogramming.pub/naming-conventions-in-go-short-but-descriptive-1fa7c6d2f32a
  - https://talks.golang.org/2014/names.slide#19
  - https://medium.com/@kdnotes/golang-naming-rules-and-conventions-8efeecd23b68
  - https://blog.golang.org/package-names
  - https://rakyll.org/style-packages/
  - https://logansbailey.com/plural-vs-singular-directory-names#:~:text=Directory%20names%20should%20be%20singular,and%20performing%20actions%20on%20them
### 4. Error Handling
  - https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
  - https://www.youtube.com/watch?v=lsBF58Q-DnY
  - https://blog.golang.org/error-handling-and-go
### 5. Testing
  - https://github.com/stretchr/testify
  - https://github.com/vektra/mockery
  - https://tutorialedge.net/golang/improving-your-tests-with-testify-go/
  - https://www.guru99.com/unit-test-vs-integration-test.html#:~:text=Unit%20Testing%20test%20each%20part,see%20they%20are%20working%20fine.&text=Unit%20Testing%20is%20executed%20by,performed%20by%20the%20testing%20team.
  - https://codeburst.io/unit-testing-for-rest-apis-in-go-86c70dada52d
  - https://medium.com/@benbjohnson/structuring-tests-in-go-46ddee7a25c
  - https://www.youtube.com/watch?v=_NKQX-TdNMc
  - https://netmind.net/en/positive-vs-negative-test-cases/#:~:text=Positive%20test%20cases%20ensure%20that,or%20by%20using%20invalid%20data
  - https://github.com/agusrichard/go-workbook/tree/master/restapi-test-app
### 6. Idiomatic Go
  - https://www.youtube.com/watch?v=yeetIgNeIkc
  - https://blog.golang.org/defer-panic-and-recover#:~:text=A%20defer%20statement%20pushes%20a,perform%20various%20clean%2Dup%20actions
  - https://dave.cheney.net/2020/02/23/the-zen-of-go