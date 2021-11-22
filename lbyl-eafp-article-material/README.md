# Look Before You Leap vs Easier to Ask for Forgiveness than Permission

## [Ask for Forgiveness or Look Before You Leap?](https://switowski.com/blog/ask-for-permission-or-look-before-you-leap)

### Introduction
- If you “look before you leap”, you first check if everything is set correctly, then you perform an action. For example, you want to read text from a file. What could go wrong with that? Well, the file might not be in the location where you expect it to be. So, you first check if the file exists:
  ```python
  import os
  if os.path.exists("path/to/file.txt"):
      ...

  # Or from Python 3.4
  from pathlib import Path
  if Path("/path/to/file").exists():
      ...
  ```
- Even if the file exists, maybe you don’t have permission to open it? So let’s check if you can read it:
  ```python
  import os
  if os.access("path/to/file.txt", os.R_OK):
      ...
  ```
- But what if the file is corrupted? Or if you don’t have enough memory to read it? This list could go on. Finally, when you think that you checked every possible corner-case, you can open and read it:
  ```python
  with open("path/to/file.txt") as input_file:
      return input_file.read()
  ```
- Even when you think you covered everything, there is no guarantee that some unexpected problems won’t prevent you from reading this file. So, instead of doing all the checks, you can “ask for forgiveness.”
- With “ask for forgiveness,” you don’t check anything. You perform whatever action you want, but you wrap it in a try/catch block
- If an exception happens, you handle it. You don’t have to think about all the things that can go wrong, your code is much simpler (no more nested ifs), and you will usually catch more errors that way. That’s why the Python community, in general, prefers this approach, often called “EAFP” - “Easier to ask for forgiveness than permission.”
- Example:
  ```python
  try:
      with open("path/to/file.txt", "r") as input_file:
          return input_file.read()
  except IOError:
      # Handle the error or just ignore it
  ```
- Here we are catching the IOError. If you are not sure what kind of exception can be raised, you could catch all of them with the BaseException class, but in general, it’s a bad practice. 

### “Ask For Forgiveness” vs “Look Before You Leap” - speed
- Time for a simple test. Let’s say that I have a class, and I want to read an attribute from this class. But I’m using inheritance, so I’m not sure if the attribute is defined or not. I need to protect myself, by either checking if it exists (“look before you leap”) or catching the AttributeError (“ask for forgiveness”:
  ```python
  # permission_vs_forgiveness.py

  class BaseClass:
      hello = "world"

  class Foo(BaseClass):
      pass

  FOO = Foo()

  # Look before you leap
  def test_lbyl():
      if hasattr(FOO, "hello"):
          FOO.hello

  # Ask for forgiveness
  def test_aff():
      try:
          FOO.hello
      except AttributeError:
          pass
  ```
- Run benchmarking:
  ```shell
  $ python -m timeit -s "from permission_vs_forgiveness import test_lbyl" "test_lbyl()"
  2000000 loops, best of 5: 155 nsec per loop

  $ python -m timeit -s "from permission_vs_forgiveness import test_aff" "test_aff()"
  2000000 loops, best of 5: 118 nsec per loop
  ```
- “Look before you leap” is around 30% slower (155/118≈1.314).
- “Look before you leap” is now around 85% slower (326/176≈1.852). So the “ask for forgiveness” is not only much easier to read and robust but, in many cases, also faster. Yes, you read it right, “in many cases,” not “in every case!”

### The main difference between “EAFP” and “LBYL”
- Example:
  ```python
  # permission_vs_forgiveness.py

  class BaseClass:
      pass  # "hello" attribute is now removed

  class Foo(BaseClass):
      pass

  FOO = Foo()

  # Look before you leap
  def test_lbyl3():
      if hasattr(FOO, "hello"):
          FOO.hello

  # Ask for forgiveness
  def test_aff3():
      try:
          FOO.hello
      except AttributeError:
          pass
  ```
- The tables have turned. “Ask for forgiveness” is now over four times as slow as “Look before you leap” (562/135≈4.163). That’s because this time, our code throws an exception. And handling exceptions is expensive.
- If you expect your code to fail often, then “Look before you leap” might be much faster.

### Verdict
- “Ask for forgiveness” results in much cleaner code, makes it easier to catch errors, and in most cases, it’s much faster. No wonder that EAFP (“Easier to ask for forgiveness than permission”) is such a ubiquitous pattern in Python.
- “Look before you leap” often results in a longer code that is less readable (with nested if statements) and slower. And following this pattern, you will probably sometimes miss a corner-case or two.
- Just keep in mind that handling exceptions is slow. Ask yourself: “Is it more common that this code will throw an exception or not?” If the answer is “yes,” and you can fix those problems with a well-placed “if,” that’s great!
- But in many cases, you won’t be able to predict what problems you will encounter. And using “ask for forgiveness” is perfectly fine - your code should be “correct” before you start making it faster.

<br />

---

## [In Python, Don’t Look Before You Leap](https://betterprogramming.pub/in-python-dont-look-before-you-leap-cff250881930)

### Look Before You Leap
- LBYL is the traditional programming style in which we check if a piece of code is going to work before actually running it.
- In other words, if a piece of code needs some prerequisites, we place conditional statements such that the code only runs if all the prerequisites are met.
- Snippet:
  ```python
  person = {'name': 'John Doe', 'age':30, 'gender': 'male'}
  
  #LBYL
  if 'name' in person and 'age' in person and 'gender' in person:
      print("{name} is a {age} year old {gender}.".format(**person))
  else:
      print("Some keys are missing")
  ```

### Easier To Ask for Forgiveness Than Permission
- In this approach, we simply run our code, hoping for the best while being prepared to handle any errors if the code fails.
- Typically, this means enclosing our code in try-except blocks and handling any exceptions that might occur.
- Snippet:
  ```python
  person = {'name': 'John Doe', 'age':30, 'gender': 'male'}
  
  #EAFP
  try:
      print("{name} is a {age} year old {gender}.".format(**person))
  except KeyError:
      print("Some keys are missing")
  ```

### So, Which One’s Better?
- As a rule of thumb, EAFP is considered more Pythonic and should be preferred in most scenarios.
- Here are some reasons why EAFP is preferred over LBYL.
  - Explicit and more readable
    - EAFP makes the “happy path” more explicit and readable. “Explicit is better than implicit” is an important tenet of Python.
    - In the example above, we expect that the keys will be present in the dictionary in most cases, which is what the EAFP code suggests. 
    - The LBYL code, however, emphasizes the rare case where the keys are missing.
  - Better performance
    - EAFP is usually faster than LBYL — especially when lots of checks are needed.
    - Calling the dictionary thrice is time-consuming and repetitive.
    - It’s true that exceptions are costlier than if statements.
    - Exceptions are only triggered a few times, whereas an if statement is always executed.
  - Prevents race conditions
    - Most importantly, EAFP helps us avoid race conditions. Race conditions occur when multiple threads are trying to access an object.
    - Consider two threads trying to access the person dictionary in the LBYL scenario. If one of the threads deletes a key from the dictionary and the second thread is past the checks, we’ll get an exception when the print statement is executed.

### There are some scenarios where using LBYL does make more sense.
- Complex side effects
  - Suppose you begin writing to a file and get an exception when some changes have been made. Trying to revert the changes in the except block can be tricky.
  - Using LBYL to stop the operation beforehand will make your life much easier.
- Too many exceptions
  - If you expect exceptions to be thrown at multiple points in your code, using EAFP might diminish your ability to trace where the code actually failed. Ideally, you should refactor your code to have narrow try-except blocks and continue to use EAFP.