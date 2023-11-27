package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("Expected %s but got %s: ", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 5)
	fmt.Println(repeated)
	// Output: bbbbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
