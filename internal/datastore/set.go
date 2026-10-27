package datastore

type Set struct {
	items map[string]struct{}
}

func NewSet() *Set {
	return &Set{
		items: make(map[string]struct{}),
	}
}

func (s *Set) Add(value string) bool {

	if _, ok := s.items[value]; ok {
		return false
	}

	s.items[value] = struct{}{}

	return true
}

func (s *Set) Remove(value string) bool {

	if _, ok := s.items[value]; !ok {
		return false
	}

	delete(s.items, value)

	return true
}

func (s *Set) Contains(value string) bool {

	_, ok := s.items[value]

	return ok
}

func (s *Set) Members() []string {

	result := make([]string, 0, len(s.items))

	for value := range s.items {
		result = append(result, value)
	}

	return result
}

func (s *Set) Len() int {
	return len(s.items)
}
