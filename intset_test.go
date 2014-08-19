package sets

import (
	"testing"
)

func TestIntSet(t *testing.T) {
	x := IntSet{}
	x.Add(1)
	x.Add(2)
	x.Add(1)
	if !x.Contains(1) {
		t.Errorf("Should have had 1")
	}
	if x.Contains(3) {
		t.Errorf("Should not have had 3")
	}
	x.Remove(2)
	if x.Contains(2) {
		t.Errorf("Should not have had 2")
	}
	y := x.Copy()
	if len(y) != 1 {
		t.Errorf("Should have had 1 element")
	}
	if !y.Contains(1) {
		t.Errorf("Should have had 1")
	}
	yString := y.String()
	if "[1]" != yString {
		t.Errorf("output was wrong: %v", yString)
	}
	y.Add(5)
	yString = y.String()
	if "[1 5]" != yString && "[5 1]" != yString {
		t.Errorf("output was wrong: %v", yString)
	}
	z := IntSet{}
	zString := z.String()
	if "[]" != zString {
		t.Errorf("output was wrong: %v", zString)
	}
}
