# SOLID principles. 

##  Single Responsibility Principle (SRP)

This principle states that a class should have only one reason to change. If we violate this principle, the class will have multiple responsibilities, making it harder to maintain, test and extend. This can lead to code that is tightly coupled, difficult to reuse, and prone to errors.

## Open-Closed Principle (OCP):
This principle states that classes should be open for extension but closed for modification. If we violate this principle, we may have to modify existing code to add new functionality, which can introduce bugs and make it difficult to maintain the code. This can also result in code that is difficult to test and reuse.


## Liskov Substitution Principle
This principle states that subtypes should be substitutable for their base types. If we violate this principle, we may introduce behavior that is unexpected and inconsistent, which can lead to errors that are difficult to track down. This can also make it difficult to write code that works with a variety of different types.
Этот принцип гласит, что подтипы должны заменять их базовые типы. 

## Interface Segregation Principle (ISP): 
This principle states that clients should not be forced to depend on interfaces they do not use. 


## Dependency Inversion Principle (DIP):
This principle states that high-level modules should not depend on low-level modules. Instead, both should depend on abstractions. If we violate this principle, we may have code that is difficult to test and reuse, and that is tightly coupled. This can also result in code that is difficult to maintain and extend.