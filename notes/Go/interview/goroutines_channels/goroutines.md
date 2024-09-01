# Goroutine

*Do not communicate by sharing memory; instead, share memory by communicating*

**Goroutine** - this is the function is executed in parallel with other goroutines

`go list.Sort()  // run list.Sort concurrently; don't wait for it.`


```
func Announce(message string, delay time.Duration) {
    go func() {
        time.Sleep(delay)
        fmt.Println(message)
    }()  // Note the parentheses - must call the function.
}
```

