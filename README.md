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


## Stuff:

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
-Over this:
```go
u, err := retrieveUser("sally")
	if err != nil {
		fmt.Println(err)
		return
	}
```