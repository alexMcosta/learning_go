package iteration

import (
	"fmt"
	"testing"
)

// Functionality Testing
func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s", expected, repeated)
	}
}

//Benchmark Testing
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

//Example for Docs
func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
