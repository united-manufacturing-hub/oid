package managementConsole

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMgmtConsoleAsn10id(t *testing.T) {
	assert.Equal(t, []int{1, 3, 6, 1, 4, 1, 59193, 100}, GetMgmtConsoleAsn10id())
}
