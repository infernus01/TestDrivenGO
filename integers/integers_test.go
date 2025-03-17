package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum) // %d is a placeholder for the actual value(integers) and %q is for strings
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

//the above ExampleAdd has the "// Output: 6" in it's syntax, so when we run the test, it will compare the output of Sum with 6.