## Errors


**Error** in go this any type of data for which defined method error that returns string 

To define a custom error type, you must satisfy the predeclared error interface.


```
type error interface{
  Error() string
}
```

Example:

```
type SyntaxError struct {
	Line int
	Col  int
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("%d:%d: syntax error", e.Line, e.Col)
}
```


create a new error:

`err1 := errors.New("math: square root of negative number")`
`err2 := fmt.Errorf("math: square root of negative number %g", x)`

### How to check our error

- errors.Is(err, TYPE_OF_ERROR)

- errors.As

```
  err := NewRetryErr(3)
  var retryError retryError
  errors.As(err, &retryError)
  ```


- switch case and type assertion 
`err.(type)`

General problem cannot find wrapped errors, can detect just a hight level error

## Wrapping

Wrapping adding new context information to the returned error
it can be type, reason or name of function

