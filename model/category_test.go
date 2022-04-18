package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestCategoryList(t *testing.T) {
	categorys, err := CategoryList(1)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(categorys))
}