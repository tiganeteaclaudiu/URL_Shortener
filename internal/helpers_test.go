package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestRandSeq tests successful generating a random requence
func TestRandSeq(t *testing.T) {

	received := randSeq(5)
	assert.Len(t, received, 5)
}
