package managementConsole

import "github.com/united-manufacturing-hub/oid/cmd/oid"

const (
	id = 200
)

func GetYubikeyManagementToolAsn10id() []int {
	return oid.ConcatCopyPreAllocate([][]int{oid.GetIanaRoot(), oid.GetUMHPen(), {id}})
}
