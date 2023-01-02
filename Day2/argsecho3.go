package main

// args-echo2 prints the command line arguments received.(Using strings.Join )
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
