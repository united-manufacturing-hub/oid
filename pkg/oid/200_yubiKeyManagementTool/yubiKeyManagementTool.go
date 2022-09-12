package managementConsole

import "github.com/united-manufacturing-hub/oid/pkg/oid"

const (
	id = 200
)

func GetYubikeyManagementToolAsn10id() []int {
	return oid.MergeSlices([][]int{oid.GetIanaRoot(), oid.GetUMHPen(), {id}})
}
