package stringset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSet(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		s := NewStringSet()
		assert.Empty(t, s)
	})

	t.Run("insert element", func(t *testing.T) {
		s := NewStringSet()
		assert.NoError(t, s.Insert("a"))
		assert.NoError(t, s.Insert("a")) // duplicate
		assert.NoError(t, s.Insert("b"))

		assert.Equal(t, 2, len(s))
	})

	t.Run("remove element", func(t *testing.T) {
		s := NewStringSet()
		assert.NoError(t, s.Insert("a"))
		assert.NoError(t, s.Insert("a")) // duplicate
		assert.NoError(t, s.Insert("b"))

		assert.NoError(t, s.Remove("c"))
		assert.Equal(t, 2, len(s))
		assert.True(t, assertContain(t, s, "a"))
		assert.True(t, assertContain(t, s, "b"))

		assert.NoError(t, s.Remove("a"))
		assert.Equal(t, 1, len(s))
		assert.False(t, assertContain(t, s, "a"))
		assert.True(t, assertContain(t, s, "b"))

		assert.NoError(t, s.Remove("a"))
		assert.Equal(t, 1, len(s))
		assert.False(t, assertContain(t, s, "a"))
		assert.True(t, assertContain(t, s, "b"))

		assert.NoError(t, s.Remove("b"))
		assert.Empty(t, s)
		assert.False(t, assertContain(t, s, "a"))
		assert.False(t, assertContain(t, s, "b"))
	})
}

func assertContain(t *testing.T, s Stringset, key string) bool {
	v, err := s.Contains(key)
	assert.NoError(t, err)
	return v
}
