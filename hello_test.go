package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("James")
	if got != "Hello, James" {
		t.Fatalf("Expected: \"Hello, James\", got %s", got)
	}
}
