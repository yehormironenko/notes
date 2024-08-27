## Slices

Slices are more flexible than arrays. They created array and if needed expand it/

We can create slices in few ways. 

`s = make([]byte,s,c)`
s - size
c - capacity

`bar := []byte{}`

for case when we cannot set *capacity* it will be equal *size*

if *size* > *capacity* -> 