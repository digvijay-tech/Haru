package interpreter

// runtimeError print the error message and ends the program execution
func runtimeErr(msg string) {
	panic("Runtime Error: " + msg)
}
