# Introduction to Go language.





#### About Go.

--- 

*"Go is an open source programming language that makes it easy to build simple, reliable and efficient software."*
* Go was developed by **Robert Griesemer**, **Rob Pike** and **Ken Thompson** at Google. It was started in Sept 2007.
* This first official announcement of Go was done in Nov 2009.
* Go has a surface similarity to the mother language **C**, and like it is a professional tool for achieving the maximum benefits with minimum efforts.
* Go has an automatic memory management or garbage collection.
* Go is often called as "Golang".
*When you are learning a new programming language, it is easier to get biased and use the programming style of the language which you already know. Please be aware of it.*

---

**The first Go program**

```go
package main
// The first program of any language.
import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```

---

**Features of Golang**

* Go is a compiled language. The Go toolchain converts a source program and it's dependencies into an instructions in the native language (binary or machine code) of a computer.
* Due to this feature it outperform all those languages that are interpreted or have virtual runtimes.
* The go tools are accessed by a single command called ```go``` , which has a number of sub-commands.
* To run a golang program, we use ```go run <program name>```. For example, ```go run helloworld.go```.
* Go handles the Unicode, therefore it can process text in all languages.
* In order to create a build for your program, and save the complied result for the later use, then you can : ```go build <program file name>```. for example, ```go build ./helloworld.go```
* go build command will generate an executable file, which can be run any time without further processing.

---
**Anatomy of your first Go program**

```go
package main 
// The first program of any language.

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```
The package main declaration
```go
package main
```

Before getting into package main declaration. Let us first understand about packages in general.
* Go code is organised into packages, which are similar to libraries or modules in other languages.
* A package consists of one or more .go source files in a single directory that defines what package does.
* Each source file begins with a package declaration, which states which package it belongs to.

Now, coming to ```package main``` declaration.
* Package main is the package declaration for that source file, also it is special because it defines the standalone executable program and not a library or dependencies.
* We can also say that, the program execution starts from the package main source file.

---

```
// The first program of any language.
```
* This a comment in golang. Golang support single line comments using ```//``` .

---
```
import "fmt"
```
* This tells the program to import a package called ```"fmt"```, which the part of Go standard library.

---

```
func main() {
	fmt.Println("Hello, Go!")
}
```
* `func` is used to declare a function inside a golang.
* However `func main` is a special function declaration, this is because the execution on the program starts from this function.

---
 ## Notes :
* Go does not require semi-colons.
* Memory management / Garbage collection: As a program runs they write objects to memory. At some point these objects should be removed, when they are not needed anymore.This process is called Memory management or garbage collection.
* `go fmt` tools is used to format go code. This is mostly installed with your editors like VS-Code or Goland.