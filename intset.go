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
