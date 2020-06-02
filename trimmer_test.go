package moco

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhitespaceSplit(t *testing.T) {
	s := "some text with spaces"
	v := WhitespaceSplit(s)

	assert.Equal(t, 4, len(v), "Length of text slice is wrong")
}

func TestBlacklistTrim(t *testing.T) {
	s := []string{"some", "text", "with", "blacklist", "words"}
	b := []string{"blacklist"}
	v := BlacklistTrim(s, b)

	assert.Equal(t, 4, len(v), "Length of text slice is wrong")
}

func TestNumericTrim(t *testing.T) {
	s := "100,000.00"
	v, err := NumericTrim(s)

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, "100000", v, "Numeric value is wrong")
}

func TestNumericHundreds(t *testing.T) {
	s := "100,300.00"
	v, err := NumericTrim(s)

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, "100300", v, "Numeric value is wrong")
}

func TestNumericDozens(t *testing.T) {
	s := "100,370.00"
	v, err := NumericTrim(s)

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, "100370", v, "Numeric value is wrong")
}
