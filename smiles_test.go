package calculator

import "testing"

func TestSmile(t *testing.T) {
	if got := (Smiles{}).Smile(); got != ":)" {
		t.Fatalf("Smile() = %q", got)
	}
}
