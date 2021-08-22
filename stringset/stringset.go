package stringset

import "fmt"

type Stringset map[string]bool

func NewStringSet() Stringset {
	set := make(map[string]bool)
	return set
}

// Insert adds a string value to the set.
func (s Stringset) Insert(value string) error {
	if s == nil {
		return fmt.Errorf("the reciever Stringset is nil")
	}

	s[value] = true
	return nil
}

// Remove a value from the string set
func (s Stringset) Remove(value string) error {
	if s == nil {
		return fmt.Errorf("the reciever Stringset is nil")
	}

	delete(s, value)
	return nil
}

// Contains return true if the given value exists in the set
func (s Stringset) Contains(value string) (bool, error) {
	if s == nil {
		return false, fmt.Errorf("the reciever Stringset is nil")
	}

	_, found := s[value]
	return found, nil
}
