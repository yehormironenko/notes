# Functions with Closure

Closure - nested function, which have access to variables of an external function even after the last one is completed.

## Example 
```
func incrementer() func() int {
    i := 0

    return func() int {
        i++
        return i
    }
}
```