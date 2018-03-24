package module

import (
	"testing"
)

func TestXYZ(t *testing.T) {
	value := Message()
	if value != "module" {
		t.Error("Expected returned string to be 'module'")
	}
}
