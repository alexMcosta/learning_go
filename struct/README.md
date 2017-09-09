# STRUCTS

##Structs is Go's way of OOP

Declaring a struct looks like so:

```
    type coworker struct {
   	firstname string
   	lastame   string
   	age       int
   	developer bool
   }
   
```

Creating a a value of type coworker:

```
    
    co1 := coworker{"Holden", "Ender", 27, true}
    
```

This will create a coworker with first name Holden, last name Ender, Age 27, and
marked as true for developer.

 