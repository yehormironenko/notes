## Slices

Slices are more flexible than arrays. They created array and if needed expand it/

```
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```

We can create slices in few ways. 

`s = make([]byte,s,c)`
s - size
c - capacity

`bar := []byte{}`

for case when we cannot set *capacity* it will be equal *size*

if *size* > *capacity* ->  panic


### Append

append return a new slice 
if count of elements which we added less than capacity, then append will return new slice which contains the same link to array. If count of added elements more than cap< then append will return new slice with new pointer to the array


### Copy 

Create new independent copy of slice