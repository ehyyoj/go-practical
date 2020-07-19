package calculate_test

import (
	"testing"

	"github.com/ehyyoj/go-practical/testing-package/calculate"
)

func TestAdd(t *testing.T) {
	if res := calculate.Add(1, 2); res != 3 {
		t.Errorf("expected: 3, actual: %d", res)
	}
	if res := calculate.Add(1, -1); res != 0 {
		t.Errorf("expected: 0, actual: %d", res)
	}
}

func TestMulti(t *testing.T) {
	if res := calculate.Multi(2, 3); res != 6 {
		t.Errorf("expected: 6, actual: %d", res)
	}
}
