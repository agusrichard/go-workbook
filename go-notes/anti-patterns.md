# Anti Patterns in Go

</br>

## List of Contents:
### 1. [5 Common mistakes in Go](#content-1)


</br>

---

## Contents:

## [5 Common mistakes in Go](https://medium.com/deepsource/5-common-mistakes-in-go-2d57fbd9357b) <span id="content-1"></span>


### 1. Infinite recursive call
- A function that calls itself recursively needs to have an exit condition.
- Some languages have tail-call optimization, which makes certain infinite recursive calls safe to use.
- Tail-call optimization allows you to avoid allocating a new stack frame for a function because the calling function will return the value that it gets from the called function.

### 2. Assignment to nil map
- A new, empty map value is made using the built-in function make, which takes the map type and an optional capacity hint as arguments:
  ```go
  make(map[string]int)
  make(map[string]int, 100)
  ```
- A nil map is equivalent to an empty map except that no elements may be added.
- Example:
  ```go
  // Bad pattern
  var countedData map[string][]ChartElement

  // Good pattern
  countedData := make(map[string][]ChartElement)
  ```
- The thing that you have to remember is if you want to make a slice or map, it's better to use make.


### 3. Method modifies receiver
- A method that modifies a non-pointer receiver value may have unwanted consequences.
- To propagate the change, the receiver must be a pointer.
- Example:
  ```go
  // Bad
  type data struct {
  	num   int
  	key   *string
  	items map[string]bool
  }

  func (d data) vmethod() {
  	d.num = 8
  }

  func (d data) run() {
  	d.vmethod()
  	fmt.Printf("%+v", d) // Output: {num:1 key:0xc0000961e0 items:map[1:true]}
  }

  // Good
  type data struct {
  	num   int
  	key   *string
  	items map[string]bool
  }

  func (d *data) vmethod() {
  	d.num = 8
  }

  func (d *data) run() {
  	d.vmethod()
  	fmt.Printf("%+v", d) // Output: &{num:8 key:0xc00010a040 items:map[1:true]}
  }
  ```

### 4. Possibly undesired value being used in goroutine
- In the example below, the value of index and value used in the goroutine are from the outer scope. Because the goroutines run asynchronously, the value of index and value could be (and usually are) different from the intended value.
- Bad Example:
  ```go
  mySlice := []string{"A", "B", "C"}
  for index, value := range mySlice {
  	go func() {
  		fmt.Printf("Index: %d\n", index)
  		fmt.Printf("Value: %s\n", value)
  	}()
  }
  ```
- Better Example:
  ```go
  mySlice := []string{"A", "B", "C"}
  for index, value := range mySlice {
  	index := index
  	value := value
  	go func() {
  		fmt.Printf("Index: %d\n", index)
  		fmt.Printf("Value: %s\n", value)
  	}()
  }
  ```
- Another way:
  ```go
  mySlice := []string{"A", "B", "C"}
  for index, value := range mySlice {
  	go func(index int, value string) {
  		fmt.Printf("Index: %d\n", index)
  		fmt.Printf("Value: %s\n", value)
  	}(index, value)
  }
  ```

### 5. Deferring Close before checking for a possible error
- Bad Example:
  ```go
  f, err := os.Open("/tmp/file.md")
  if err != nil {
      return err
  }
  defer f.Close()
  ```
- Better way:
  ```go
  f, err := os.Open("/tmp/file.md")
  if err != nil {
  	return err
  }

  defer func() {
  	closeErr := f.Close()
  	if closeErr != nil {
  		if err == nil {
  			err = closeErr
  		} else {
  			log.Println("Error occured while closing the file :", closeErr)
  		}
  	}
  }()
  return err
  ```


**[⬆ back to top](#list-of-contents)**

</br>

---

## References:
- https://medium.com/deepsource/5-common-mistakes-in-go-2d57fbd9357b