package moco

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCsvReader(t *testing.T) {

	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"`

	ioReader := strings.NewReader(in)

	r := NewCsvReader(ioReader, ',')
	records, err := r.ReadAll()

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 4, len(records), "Length of data should be 4")
}
