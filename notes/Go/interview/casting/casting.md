# Casting

## How to use type casting in Golang
To use type casting in Golang, you need to specify the type you want to convert to using the cast operator which is denoted by the type name in parentheses before the variable or expression you want to convert.

```
var f float32 = 3.14
var i int = int(f)
```

## Golang type conversions
To use type casting in Golang, you need to specify the type you want to convert to using the cast operator which is denoted by the type name in parentheses before the variable or expression you want to convert. For example, to convert a float number to an integer, you can use the following expression:

```
var i int = 5
var f float32 = 3.14
var result = f + i
```

## Go type assertion
Golang also supports type assertion, which is a mechanism for extracting the concrete type from an interface value. Type assertion is used when you have an interface value and need to access its underlying type. Here is an example:

```
var val interface{} = "hello"
s := val.(string)
```

In this example, the interface value val is assigned the string value “hello.” Type assertion is used to extract the concrete type of val and assign it to the variable s.