# Context 

## What is context 

Context this is interface from the package context:

```
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key any) any
}
```

**Purposes**
- for deadlines code executing
- to say about finish 
- understand the reason of cancel 
- Get a value by a key


## Root context 

Root context usually creating in the main. 
It's not possible to cancel it and it doesn't have a deadline.

2 functions to create a root context:
- context.Background()
- context.TODO()

They are very similar, just one TODO() - temproray, just for tests or experiments 

## Creating a new context 
All contexts creating on the base of root context. New context extends root context. 
The following functions exists in the package for creating new context:
- WithCancel
- WithDeadline
- WithTimeout
- WithValue
- WithoutCancel

### WithCancel
Returns new context and with cantcel function. Neew to call int the same place where we cretted it 
```
ctx, cancel := context.WithCancel(context.Background())
defer cancel() // cancel when we are finished
```

### WithDeadline and WithTimeout

We can limit our context time using functions above. 

```
func slowOperationWithTimeout(ctx context.Context) (Result, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	return slowOperation(ctx)
}
```

### WithValue
Adding pair of key value to the context 
```
type favContextKey string

key := favContextKey("language")
ctx := context.WithValue(context.Background(), key, "Go")
```
Get value:
`ctxValue := ctx.Value(key)`

## Error handling in the Go

If context was canceled we can know it using Done and Err
Done this is channel which closed when context was canceled 

Err can returns three diffrent values:
- **nil** - if context still active
- **context.Canceled** - if context was canceled by user 
- **context.DeadlineExceeded** - if context was canceled by timeout


## WithoutCancel

In the go 1.21 have been added the function `WithoutCancel`
returns a copy of the rune context which won't be canceled with root context.

Cases to use:
- for rollback/cleanup
- for a long term operations 

## AfterFunc 

Allow to register the functions which will called after the cancel of the context.
