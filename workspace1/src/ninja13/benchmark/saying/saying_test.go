package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("mata")
	if s != "Welcome mata" {
		t.Error("Wow")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("vova"))
	//Output:
	//Welcome vova
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
