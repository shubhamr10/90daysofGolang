package main

// args-echo2 prints the command line arguments received.(Using for range loop)
import (
	"fmt"
	"os"
)

func main() {
	args, separate := "", ""
	for _, arg := range os.Args[1:] {
		args += separate + arg
		separate = " "
	}
	fmt.Println(args)
}
