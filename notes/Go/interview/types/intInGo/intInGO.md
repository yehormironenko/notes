## Int в Golang

### What numerical types are there in go
- int/int8/int16/int32/int64;
- uint/uint8/uint16/uint32/uint64;
- float32/float64; 
- complex64/complex128;
- rune(int32).

### Diffrence between int and uint

`int` - this is range between negative and positive values   
`uint` - only negative values 

### Simple int 
Architecture of machine will translate or int to `int32` or `int64`

### How to translate int to string 
We should use functions from the package `strconv`
for `int` - `strconv.Atoi` and `strconv.Itoa` 
for `int64` - `strconv.ParseInt` and `strconv.FormatInt`

### Memory for int32 and int64
for int32  - 4 bytes
for  int64 - 8 bytes

### Whar result we get if we devide by 0 int and float
- for `int` compilation error
- for `float` infinity