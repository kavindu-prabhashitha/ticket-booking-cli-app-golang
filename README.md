### Best Practice ###
- Define variable as "local" as possible
- Create the variable where you need it

### Slice ###
- slice is an abstraction of an Array
- slices are more flexible and powerful: variable-length or get an sub arrya of its own
- slices are also index based and have a size, but is resized when needed


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
- All keys have the same data type
- All values have the same data type