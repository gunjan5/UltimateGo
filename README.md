# UltimateGo
Hardcore Go training (Ardan Labs)

##Tips:

- Every decision has a cost 
- 10 lines of code = 1 bug
- Every function is a data transformation problem 
- Mechanical sympathy ??
- Tighter scope is better
- smaller scope = smaller var name, large scope = longer/meaningful var names
- Program flow is not single entry single exit, it's single entry multiple exits (i.e. if err { return})
- Few lines of copy/paste is better than having a dependancy
- Stay on stack if possible to reduce the stress on GC
- Every goroutine you create is like making a baby. Think about it before you make one
- Don't use `init()` when order of execution matters
- Always think like this: "Every goroutine that is in runnable state is RUNNING at the same time"

## Golang Stuff:

- To see if some var is escaping to heap use: `go build -gcflags -m`
- To see assembly for your code `go build -gcflags -S`
- Inlining is when the compiler makes the function code inline, to avoid function call overhead
- Constants have 256 bits precision 
- type aliasing is for 1. methods attached to them or 2. protect the type //type duration int64

- Good:
```go
    var u user
	err = json.Unmarshal([]byte(r), &u)
	return &u, err
```
- Bad (not readable)
```go
    var u *user
	err = json.Unmarshal([]byte(r), u)
	return u, err
```

- Prefer this:
```go
if u, err := retrieveUser("sally"); err != nil {
		fmt.Println(err)
		return
	}
```
- Over this:
```go
u, err := retrieveUser("sally")
	if err != nil {
		fmt.Println(err)
		return
	}
```

- Every string has a backing array for the data and header with a pointer to the data and length
- Whenever you pass a string, it's actually just copying the header (need to confirm this)

- string chars have 3 things 1. char 2. codepoint 3. bytes
- English (ASCII) chars have 1/1/1, but other chars such as Chinese chars 世 have 1/1/3 


-------------
- When you take a slice out of a slice `slice2 := slice1[2:4]` and don't want `slice2 = append(slice2, "CHANGED")` to affect `slice1`, then use `slice2 := slice1[2:4:4]` which is start from index 2, upto 4 (not included), len=2, cap=2, so when you append, it creates a new backing array! 


-------------
- Think of every folder in your go code as a mini-library/component, files inside the folder to organize the code
- Package names should be short and simple (no "-", no "_", no camelCase), make it like strings or fmt
- file.go in the package bulk of the public API, the place where the main logic/understanding is


-------------
- Don't share the primitive types (no need to use pointers for methods, so it can stay in the stack)
- When you're not sure use pointer receiver


-------------
- For custom errors, use Error variables only if you have to compare and see if it's this error or that error. Use ErrXxxxXxxx format for the error variable 
- Custome error types. type XxxXxxError struct { } format
- When playing with errors, always use error interface, don't play with the concrete 


-------------
- `go` keyword, ch ops, garbage collection, syscall (such as fmt.Println) (only times when scheduler can say to to goroutines you stop and you start)
- start, stop, shutdown, init (for every goroutine)


-------------
- Buffered channels = faster speed but no guarantee 
- Everything needs to have a timeout 
- Only take the amount of work you can handle 
- your code should be able to handle a shutdown when it's at full throttle 


-------------
- when you stop a panic, dump the stack trace to the logs anyway 