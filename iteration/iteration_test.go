package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	received := Repeat("a", 5)
	expected := "aaaaa"

	if(received != expected) {
		t.Errorf("expected %q but got %q", expected, received)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat () {
	repeatedString := Repeat("a", 5)
	fmt.Println(repeatedString)
	// Output: aaaaa

}