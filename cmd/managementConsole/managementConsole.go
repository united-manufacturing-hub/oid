package managementConsole

import oid "oid/cmd"

const (
	id = 100
)

func GetMgmtConsoleAsn10id() []int {
	return oid.ConcatCopyPreAllocate([][]int{oid.GetIanaRoot(), oid.GetUMHPen(), {id}})
}
