package student

// StringSet represents a set of strings
type StringSet map[string]struct{}

// NewStringSet creates a new set
func NewStringSet() StringSet {
	return make(StringSet)
}

// Add adds a new element to the set
func (s StringSet) Add(value string) {
	s[value] = struct{}{}
}

// Remove removes an element from the set
func (s StringSet) Remove(value string) {
	delete(s, value)
}

// Contains checks if the set contains a specific element
func (s StringSet) Contains(value string) bool {
	_, exists := s[value]
	return exists
}

// Size returns the number of elements in the set
func (s StringSet) Size() int {
	return len(s)
}

// Elements returns all elements in the set as a slice
func (s StringSet) Elements() []string {
	elements := make([]string, 0, len(s))
	for key := range s {
		elements = append(elements, key)
	}
	return elements
}


func (s StringSet) IsEmpty() bool {
	return len(s) == 0
}