## Arrays

Collection elements of the same type. 
Length cannot be changed.

`arr := [4]int{3,2,5,4}`

If we create two arrays with different size it will be arrays with different type.
Length of the array in the Go part of his type.

```
a := [3]int{}
b := [2]int{}
```
These are arrays with different type.

`a := [...]int{1, 2, 3}` - compiler can count length of array


### Passing by Value
That is why an array in Go is a primitive data type, it can be copied when passed to another variable. By default, in Go, all values are copied rather than passed by reference. This means that if we pass our array to a function, Go will copy this array and the function will contain a completely different array (or rather, an exact copy of the original array).

```
package main

import "fmt"

func main() {
	var initArray = [...]int{1, 2, 3}
	var copyArray = initArray

	fmt.Printf("Address of initArray: %p\n", &initArray)
	fmt.Printf("Address of copyArray: %p\n", &copyArray)
}
```