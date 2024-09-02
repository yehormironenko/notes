## Const в Golang

### What is constants in Go
Constants - this si unchanged variables, not possible to change constant

### What is the iota
`iota` - an identifier that allows to create consecutive untyped constants.

`
const (
	C0 = iota
	C1
	C2
)
fmt.Println(C0, C1, C2) // "0 1 2"
``

### Good practive for enums 
 - create a new integer type
 - list its values using iota
 - give the type a String function
```
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}
```
In use:
```
var d Direction = North
fmt.Print(d)
switch d {
case North:
	fmt.Println(" goes up.")
case South:
	fmt.Println(" goes down.")
default:
	fmt.Println(" stays put.")
}
// Output: North goes up.
```