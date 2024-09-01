# Channels 

[good description](https://habr.com/ru/companies/oleg-bunin/articles/522742/)

Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
*Shortly*: Structure with buffer, 2 queues and one lock option 


```
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```

Like maps and slices, channels must be created before use:
`ch := make(chan int)`

By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.


All channels have:
- size (if you create not buffered channel size is 0)
- passing values between  goroutine this is general purpose of channels
- channels are flow-safe
- channels works as First In — First Out
- gouroutines waiting for a channels 
- Channels can block/unblock channels

4 public actions:
- `newChan := make(chan int)` - create a new channel
- `newChan <— 1` - write to the channel 
- `<-newChan` - read from the channel
- `close(newChan)` - close a channel 

## What does the channel consist of?

### Structs
`type hchan struct — struct which contains main data`
`type sudog struct — super set for work with diffrent fields like select e.g.`
`type waitq struct — this struct this is linked list which persist two pointers firs and last`

### Functions
`func makechan(t *chantype, size int) *hchan — create a new channel.`
`func chansend(c *hchan, ep unsafe.Pointer, block bool, callerpc uintptr) bool — func to write into the channel`
`func chanrecv(c *hchan, ep unsafe.Pointer, block bool) (selected, received bool) — func to read from the channel.`
`func closechan(c *hchan) — func to close the channel`



## Package sync 

This type allow to define a group of goroutines that should be executed together as a group. 
And we can set a lock that will pause execution of the function until the entire group of gouroutines has finished executing. 

```
func main() { 
    var wg sync.WaitGroup 
    wg.Add(2)       // 2 gouroutines in the group
    counter := 5
    doubleCounter := func() { 
        defer wg.Done() 
        counter = counter * 2
   } 
  
   // calls gouroutines
   go doubleCounter() 
   go doubleCounter() 
  
   wg.Wait()        // waiting 2 gouriutines
   fmt.Println("Counter:", counter) 
}
```

**Wait group** can be passed as a parameter to the function but only as **pointer**

### Mutex

Mutex allows to block critical section of the code, section here - part of code which should not be executed in the several goroutines

## RW Mutex
an arbitrary number of readers can lock section or just one writer


## Sync Map

It is safe for simultaneous use.

Sync.Map is optimized for two cases, when we write a one time and will read a lot of times - as cache.
Or if few goroutines read, where and re-write for different keys

we cannot copy sync.map becaouse contas Mutex. 

## notes

go run -race 
flag for finding race conditions 

## interesting pattern in lesson 
with channel and blocker with special boolean channel 