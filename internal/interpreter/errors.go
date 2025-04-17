package interpreter

import (
	"fmt"
	"os"
)

// runtimeError print the error message and ends the program execution
func runtimeErr(msg string) {
	fmt.Println("Runtime Error:", msg)
	os.Exit(1)
}
