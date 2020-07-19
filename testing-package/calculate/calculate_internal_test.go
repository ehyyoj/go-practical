package calculate

import "testing"

func TestString2Int(t *testing.T) {
	if n := string2int("1"); n != 1 {
		t.Errorf("expected: 1, actual: %d", n)
	}
}