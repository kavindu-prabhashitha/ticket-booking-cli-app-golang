### Best Practice ###
- Define variable as "local" as possible
- Create the variable where you need it

### Why Go ###
- Concurrency in Go is cheap and easy
  - "go ..." - starts a new goroutine
  - A goroutine is a lightweight thread managed by the Go runtime
  ### Goroutine ###
   - Go is using , what's called a "Green thread "
   - Abstraction of an actual thread
   - Managed by the go runtime, we are only interacting with these high level goroutine
   - Cheaper and lightweight
   - can run hundreds of thousands or millions goroutines without affecting the performance
   - goroutines has "Channels" which is a built-in feature for go routines to talk with one another

### Slice ###
- slice is an abstraction of an Array
- slices are more flexible and powerful: variable-length or get an sub arrya of its own
- slices are also index based and have a size, but is resized when needed
- can have only one data type

### append ###
- Adds the element(s) at the end of the slice
- Grows the slice if a greater capacity is needed and return the updated slicer value

### loops ###
- In general, languages provide various control structures to control the application flow
- A loop statement allows us to execute code multiple times, in a loop
- loops are simplified in Go, we only have the "for loop" for loops

 - #### For-Each" loop ####
  - iterating over a list

### range ###
- Range iterates over elements for different data structures ( not only for arrays and slices)

### Blank Identifier "_" ###
- To ignore a variable you don't want to use
- So with Go you need to make unused variables explicit


### Functions ###
- Encapsulate code into own container (= function). Which logically belong together!
- Like variable name, you should give a function a descriptive name
- call the function by its name , whenever you want to execute this block of code
- Every program has at least one function, which is the main() function
- Function is only executed , when "called"!
- You can call a function as many times you want
- So a function is also used to reduce code duplication
- A Go function can return multiple values

### Package Level Variable ###
- Defined at the top outside all functions
- They can be accessed inside any of the functions
- And in all files, which are in the same package


### Local Variables ###
- Defined inside a function or a block
- They can be accessed only inside that function or block of code

### Packages in Go ###
- Go programs are organized into packages
- A package is a collection of Go files

### Maps ###
- Maps unique keys to values
- We cannot mix data types
- All keys have the same data type
- All values have the same data type
- syntax  ---> map[key-data-type][value-data-type]


### Struct ###
- Stands for " Structure" 
- Can hold mixed data types 

  #### "type" statement - Custom Types ####
  - The type keyword creates a new type, with the name you specify
  - In fact, we could also create a type based on every other data type like int, string etc.

### "time" - functionality for time ###
- The sleep function stops or blocks the current "thread" (goroutine) execution for the defined duration

### Waitgroup ###
- Waits for the launched goroutine to finish
- Package "sync" provides basic synchronization functionality
- Add : sets the number of goroutines to wait for 
- Wait : Block until the WaitGroup counter is 0
- Done : Decrements the WaitGroup counter by 1 , So this is called by the goroutine to indicate that it's finished