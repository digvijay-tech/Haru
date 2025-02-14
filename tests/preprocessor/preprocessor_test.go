package tests

import (
	"fmt"
	"testing"

	"github.com/digvijay-tech/Haru/internal/preprocessor"
)

func TestPreprocessor(t *testing.T) {
	source := `
--- Comment: 1

func SayHello(name: string): string { --- Comment: 2
	let msg: string = "Hello " + name; --- Comment: 3

	--- Comment: 4
	return msg;

	--- Comment: 5
}



--- Comment 6


`
	expected := `func SayHello(name: string): string {
let msg: string = "Hello " + name;
return msg;
}`

	result := preprocessor.Proprocessor(source)

	if result != expected {
		t.Errorf("Preprocess() output incorrect:\nExpected:\n%s\nGot:\n%s", expected, result)
	}

	// for verbose option
	fmt.Println("RESULT START")
	fmt.Println(result)
	fmt.Println("RESULT END")
}
