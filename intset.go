package sets

import (
	"strconv"
)

type IntSet map[int]struct{}

var Exists struct{}

func (is IntSet) Add(i int) {
	is[i] = Exists
}

func (is IntSet) Remove(i int) {
	delete(is, i)
}

func (is IntSet) Contains(i int) bool {
	_, ok := is[i]
	return ok
}

func (is IntSet) Copy() IntSet {
	c := IntSet{}
	for k, _ := range is {
		c.Add(k)
	}
	return c
}

func (is IntSet) Equals(os IntSet) bool {
	for k, _ := range is {
		if !os.Contains(k) {
			return false
		}
	}
	for k, _ := range os {
		if !is.Contains(k) {
			return false
		}
	}
	return true
}

func (is IntSet) Intersect(os IntSet) IntSet {
	out := IntSet{}
	for k, _ := range is {
		if os.Contains(k) {
			out.Add(k)
		}
	}
	return out
}

func (is IntSet) Union(os IntSet) IntSet {
	out := IntSet{}
	for k, _ := range is {
		out.Add(k)
	}
	for k, _ := range os {
		out.Add(k)
	}
	return out
}

func (is IntSet) Subtract(os IntSet) IntSet {
	out := IntSet{}
	for k, _ := range is {
		if !os.Contains(k) {
			out.Add(k)
		}
	}
	return out
}

func (is IntSet) String() string {
	s := "["
	first := true
	for k, _ := range is {
		if !first {
			s += " "
		} else {
			first = false
		}
		s += strconv.Itoa(k)
	}
	s += "]"
	return s
}
