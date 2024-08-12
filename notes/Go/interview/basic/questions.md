# Questions 

## OOP in Golang

Not classic model of OOP, Go is not OOP language. But Go contains very similar models

### Inheritance in golang
Golang doesn't support inheritance. But we have structures and thess structures can contain another structures. 
And with it fucntions of child structure will have access to the parent functions. 

```
type Parent struct{}

func (c *Parent) Print() {
	fmt.Println("parent")
}

type Child struct {
	Parent
}

func (p *Child) Print() {
	fmt.Println("child")
}

func main() {
	var x Child

	x.Print()
}
```
**Output: child**


### Incapsulation in Go

Lower or upper case. So, like privite and public. 

### Polymorphism in Go

Polymorphism in go is implemented using interfaces. 
General idea: We can declare interfaces for out types. And we need to implemet functionts  that satisfy these interfaces.


### Packages in Go
Mechanism of reusing code. In all files on the top  we add word package and his name. 
All functions here are global.

## Global variable 
This is variable declared outside the code and avaliable 

## Operators

## It is possible to call several cases in the switch case construction
Yes, using keyword `fallthrogh`

```
func main() {
	animals := []string{"bear", "bear", "rabbit", "wolf"}

	for _, animal := range animals {
		switch animal {
      
		case "rabbit":
			fmt.Println(animal, "is so weak!")
			fallthrough
      
		case "bear", "wolf":
			fmt.Println(animal, "is so strong!")
		}
	}
}
```
**output:**
bear is so strong!
bear is so strong!
rabbit is so weak!
rabbit is so strong!
wolf is so strong!

## Strings в Golang

### What is the strings in Go

String in go this is array of bytes
```
type _string struct {
	elements *byte // underlying bytes
	len      int   // number of bytes
}
```

### How do I determine the number of characters for a string?" or "What are the nuances when iterating over a string?

String this is array of bytes. 
And in standard encoding we can use: `len(string)`
But somitemes one symbol this is more than one byte and in this case we need cast it to runes:
`len([]rune(string))`
So, in the biggest part of cases we can use stadard cod

utf8.RuneCountInString(str)