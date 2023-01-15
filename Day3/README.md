# Introduction to Go language.

#### Finding duplicate lines

---
Programs for file copying, printing, searching, sorting, counting and the like all have a similar structure; a loop over the input, some computation on each element, and generation of output at the end.

We will also see some useful functions using the "finding duplicate line" example.

---
`Example 1` The first version of finding duplicate line that appears more than once in the standard input, preceded by its count. We will call the programs `dup`

```go
package main
// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		counts[input.Text()]++
    }
	// Note: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Println(fmt.Sprintf("%d \t %s \n", n, line))
        }
    }
}


```

* A map holds a set of key/value pair and provides constant time operation to store, retrieve or test for an item in the set.
* The statement `counts[input.Text()]++` is equivalent to these two statements:
  * `line :=input.Text()`
  * `counts[line]=counts[line]+1`
* `bufio` package helps to make input and output efficient and convenient. One of its features is a type called `Scanner` that reads input and breaks it into 2 lines or words; it's often the easiest way to process input that comes naturally in lines.
* `input.Scan()` reads the next line and removes the newline character from the end.
* The result of `Scan()` can be retrieved by a `input.Text()`.
* The function  `fmt.Printf()` produces formatted output fromm a list of expressions.
* `Printf()` has over dozen conversion, which Go programmers calls verbs.



| Verbs | Usage                                  |
|-------|----------------------------------------|
| `%d`  | Decimal Integer                        |
| `%x`  | Hexadecimal Integer                    |
| `%o`  | Octal Integer                          |
| `%b`  | Binary Integer                         |
| `%f`  | Floating point numbers (e.g. 3.141680) |
