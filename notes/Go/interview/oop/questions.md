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

Lower or upper cases. So, like private and public. 

### Polymorphism in Go

Polymorphism in go is implemented using interfaces. 
General idea: We can declare interfaces for out types. And we need to implemet  functionts  that satisfy these interfaces.


### Packages in Go
Mechanism of reusing code. In all files on the top  we add word package and his name. 
All functions here are global.

## Global variable 
This is variable declared outside the code and avaliable 