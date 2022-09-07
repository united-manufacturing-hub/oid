package tui

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCaCertificateAsn10id(t *testing.T) {
	assert.Equal(t, []int{1, 3, 6, 1, 4, 1, 59193, 100, 200}, GetMgmtConsoleTUIAsn10id())
}
