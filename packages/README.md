# Packages

Packages are the same as libraries in other programing languages and help you reuse code.

### Seperate Name Space
Each package has a seperate namespace from its declarations. For example, if a
function identifier called "meow" lives in a package called "cat/catlanguage" and in our package
called "ourpackage" we would have to explicitly identify the catlanguages meow from within our package
If we had to call upon it:
```
package ourpackage
import "cat/language"
catlanguage.meow()
func meow()

```

