package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainString(t *testing.T) {
	a := assert.New(t)

	arr := []string{"a", "b", "c", "d", "e"}
	a.Equal(true, ContainString(arr, "a"), "a is in arr")
	a.Equal(true, ContainString(arr, "e"), "e is in arr")
	a.Equal(false, ContainString(arr, "f"), "f is not in arr")
}
