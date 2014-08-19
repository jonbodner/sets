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
	if !y.Equals(x) {
		t.Errorf("Should be equal")
	}
	if !x.Equals(y) {
		t.Errorf("Should be equal")
	}

	a := IntSet{}
	if a.Equals(y) {
		t.Errorf("Should not be equal")
	}

	if y.Equals(a) {
		t.Errorf("Should not be equal")
	}

	a.Add(5)
	if a.Equals(y) {
		t.Errorf("Should not be equal")
	}

	if y.Equals(a) {
		t.Errorf("Should not be equal")
	}

	a.Add(1)
	if a.Equals(y) {
		t.Errorf("Should not be equal")
	}

	if y.Equals(a) {
		t.Errorf("Should not be equal")
	}

	//Intersection
	out1 := y.Intersect(a)
	out2 := a.Intersect(y)
	if !out1.Equals(out2) {
		t.Errorf("Intersections not transitive")
	}
	if len(out1) != 1 {
		t.Errorf("should only have 1")
	}
	if !out1.Contains(1) {
		t.Errorf("should have 1 for only value")
	}

	//Union
	out3 := y.Union(a)
	out4 := a.Union(y)
	if !out3.Equals(out4) {
		t.Errorf("Unions not transitive")
	}
	if len(out3) != 2 {
		t.Errorf("should have 2")
	}
	if !out3.Contains(1) {
		t.Errorf("should have 1 for value")
	}
	if !out3.Contains(5) {
		t.Errorf("should have 5 for value")
	}

	//Subtraction
	out5 := y.Subtract(a)
	out6 := a.Subtract(y)
	if out5.Equals(out6) {
		t.Errorf("Subtraction is transitive")
	}
	if len(out5) != 0 {
		t.Errorf("out5 should be empty")
	}
	if len(out6) != 1 {
		t.Errorf("out6 should only have 1")
	}
	if !out6.Contains(5) {
		t.Errorf("out6 should contain 5")
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
