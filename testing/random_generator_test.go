package testing

import (
	"testing"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/helpers"
	"github.com/stretchr/testify/assert"
)

func TestRandomGenerator(t *testing.T) {
	str := helpers.GenerateRandomString(15)

	assert.Equal(t, str, "l")
}
