# Generics 

## What are generics

Generics, also known as parametric polymorphism, enable you to write code that operates on multiple types without explicitly specifying the types.

Prior to the introduction of generics, you had to use interfaces and type assertions for achieving similar functionality. However, this approach had its drawbacks, such as a lack of type safety and increased boilerplate code.

## Type Parameters and Type Constraints

Type parameters and type constraints are the building blocks of generics in Go. Type parameters are placeholders for the actual types that will be used when the generic function or data structure is instantiated. Type constraints restrict the types that can be used as type arguments.

**Example**
```
package main

import "fmt"

func PrintSlice[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}

func main() {
    intSlice := []int{1, 2, 3}
    stringSlice := []string{"hello", "world"}

    PrintSlice[int](intSlice)
    PrintSlice[string](stringSlice)
}
```

In the PrintSlice function, `T` is a type parameter, and `any` is a type constraint. The any constraint is built into the language and allows any type to be used.


### Real-World Use Cases
Generics can be applied to a wide range of real-world use cases, such as:

- Writing generic algorithms, like sorting, searching, or filtering
- Creating generic data structures, like linked lists, trees, or queues
- Developing reusable utility functions for error handling, logging, or caching


### Note
-  Remove type arguments when calling the generic function
Note that this isn’t always possible. For example, if you needed to call a generic function that had no arguments, you would need to include the type arguments in the function call.
