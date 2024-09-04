# Interfaces
 

## List of things to keep in mind before

- **Interface segregation principle** 
  A client should never be forced to implement an interface that it doesn’t use, or clients shouldn’t be forced to depend on methods they do not use.
- **Polymorphism**
  A piece of code changes its behavior based on the concrete data it receives.
- **Liskov Substitution Principle**
  If your code depends upon an abstraction, one implementation can be replaced by another without changing your code.


An **interface** is a programming construct that describes the behavior of an object without specifying the underlying implementation details.

```
type  iface  struct {
    tab *itab  //type
    data unsafe.Pointer //data
}
```

## Why do we need interfaces 

- reducing duplicated code 
- easier to create a modul tests
- less dependent code

### Reducing duplicated code 

#### Example

If we have a structure customer and want to write data to the *bytes.Buffer* and also to the *io.File*. But in the both cases at first we should serialize it to the JSON

Standard Go interface type of io.Writer

```
type Writer interface {
        Write(p []byte) (n int, err error)
}
```



Both these types satisfy this interface because they contains the following functions bytes.*Buffer.Write()* and *os.File.Write()*.

**Impementation:**

```
package main

import (
    "encoding/json"
    "io"
    "log"
    "os"
)

// Create type Customer
type Customer struct {
    Name string
    Age  int
}

// WriteJSON, which accepts io.Writer as a param.
// it send structure Сustomer in the JSON, and if it handled with success
// then the corresponding Write() method is called from io.Writer.
func (c *Customer) WriteJSON(w io.Writer) error {
    js, err := json.Marshal(c)
    if err != nil {
        return err
    }

    _, err = w.Write(js)
    return err
}

func main() {
    //Inicialization of Customer
    c := &Customer{Name: "Alice", Age: 21}

    //Then with Buffer мwe can call func WriteJSON
    var buf bytes.Buffer
    err := c.WriteJSON(buf)
    if err != nil {
        log.Fatal(err)
    }

    // or use the file
    f, err := os.Create("/tmp/customer")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()


    err = c.WriteJSON(f)
    if err != nil {
        log.Fatal(err)
    }
}
```

### Bad pracrice

```
package tcp
type Server interface {
    Start()
}
type server struct { ... }
func (s *server) Start() { ... }
func NewServer() Server { return &server{ ... } }
```
```
package consumer
import “tcp”
func StartServer(s tcp.Server) { s.Start() }
```

Here the interface is defined in the implementer package i.e. tcp. As mentioned in Go Code Review Comments referenced earlier, this is not the best practice because the interface is defined just so one type can satisfy it. It is not defined for the purpose of abstraction. Also, it unnecessarily couples the consumer package to the implementer package.

### Good practice 

```
package tcp
type Server struct { ... }
func (s *server) Start() { ... }
func NewServer() Server { return &Server{ ... } }
```

```
package consumer
type Server interface {
    Start() 
}
func StartServer(s Server) { s.Start() }
```

Here we define the interface in the consumer package to fulfil usage requirements. This removes the unnecessary dependency on the implementer package. And it minimizes the presumptions about the way an implementation will be consumed.

It also addresses the testability concern because producer packages don’t have to provide interfaces for the user to write their own stubs. During testing, the consumer can simply create a mock implementation of the interface defined in its own scope.


### Interface Segregation Principle

***“Clients should not be forced to depend on interfaces that they do not use.”***

What this signifies is that a consumer should accept only those types whose methods are strictly required by the consumer. As an example, suppose we want to write a Document structure to a file. The function signature to accomplish this can be any one of these:

```
// os.File contains many unrelated methods 
func Save(f *os.File, doc *Document) error 
// io.ReadWriteCloser contains unrelated Read() and Close() methods 
func Save(rwc io.ReadWriteCloser, doc *Document) error 
// io.Writer contains only one method Write() that is required 
func Save(w io.Writer, doc *Document) error
```

## Links

- [Common erros for interfaces in Go](https://medium.com/@andreiboar/7-common-interface-mistakes-in-go-1d3f8e58be60) 