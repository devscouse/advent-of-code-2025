package core

import "strconv"

type (
	Set map[int]struct{}
)

func (s *Set) String() string {
	repr := "{"
	for k := range *s {
		repr += strconv.Itoa(k)
		repr += ", "
	}
	repr = repr[:len(repr)-2]
	repr += "}"
	return repr
}

func NewSet() *Set {
	set := make(Set)
	return &set
}

func (s *Set) Add(value int) *Set {
	(*s)[value] = struct{}{}
	return s
}

func (s *Set) UnionInPlace(other *Set) *Set {
	for k := range *other {
		(*s)[k] = struct{}{}
	}
	return s
}
