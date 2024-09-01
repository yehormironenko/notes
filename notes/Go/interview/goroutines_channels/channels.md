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

