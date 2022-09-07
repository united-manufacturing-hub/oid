package certificates

import (
	"github.com/united-manufacturing-hub/oid/cmd/oid"
	"github.com/united-manufacturing-hub/oid/cmd/oid/100_managementConsole"
)

const (
	id = 200
)

func GetMgmtConsoleTUIAsn10id() []int {
	return oid.ConcatCopyPreAllocate([][]int{managementConsole.GetMgmtConsoleAsn10id(), {id}})
}
