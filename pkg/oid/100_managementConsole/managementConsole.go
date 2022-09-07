package managementConsole

import "github.com/united-manufacturing-hub/oid/pkg/oid"

const (
	id = 100
)

func GetMgmtConsoleAsn10id() []int {
	return oid.ConcatCopyPreAllocate([][]int{oid.GetIanaRoot(), oid.GetUMHPen(), {id}})
}
