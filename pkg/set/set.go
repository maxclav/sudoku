package set

// String is a hash set
// that contains unique string.
type String struct {
	items map[string]struct{}
}

// NewString returns a new String
// (hash set of strings).
func NewString() *String {
	return &String{
		items: make(map[string]struct{}, 9),
	}
}

// Add adds the string in the set and return true
// or return false if the string is already in the set.
func (s *String) Add(str string) bool {
	if s.Has(str) {
		return false
	}
	s.items[str] = struct{}{}
	return true
}

// Has returns whether the string
// is in the set (true) or not (false).
func (s *String) Has(str string) bool {
	_, ok := s.items[str]
	return ok
}
