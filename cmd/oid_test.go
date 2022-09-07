package oid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetIanaRoot(t *testing.T) {

	assert.Equal(t, []int{1, 3, 6, 1, 4, 1}, GetIanaRoot())
}

func TestGetUMHPen(t *testing.T) {
	assert.Equal(t, []int{59193}, GetUMHPen())
}
