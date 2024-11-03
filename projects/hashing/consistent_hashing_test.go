package hashing

import (
	"testing"
)

func TestDefaultHash(t *testing.T) {
	key := "test"

	hash1 := DefaultHash(key)
	hash2 := DefaultHash(key)

	if hash1 != hash2 {
		t.Errorf("Hash function is not giving the same value for hte same key. Got %d and %d for the same testing input", hash1, hash2)
	}
}