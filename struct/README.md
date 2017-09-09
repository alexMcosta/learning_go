# STRUCTS

##Structs is Go's way of OOP

```
"A struct is an aggregate data type the groups together zero or more
named values of arbitrary types as a single identity" 
- The Go Programming Language
```

Declaring a struct looks like so:

```
    type Coworker struct {
   	Firstname string // These
   	Lastame   string // Are
   	Age       int // Known as
   	Developer bool // Fields
   }
   
```

Creating a a value of type coworker:

```
    
    co1 := Coworker{"Holden", "Ender", 27, true}
    
```

This will create a coworker with first name Holden, last name Ender, Age 27, and
marked as true for developer.

You can also use dot notation with an instance of a struct type like so:

```
var p2 Coworker
p2.Firstname = "Tessa"
```

### Encapsulation 

Encapsulation is defined by 

### Reusibility

### Polymorphism

### Overriding 

 