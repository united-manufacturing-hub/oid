package certificates

import (
	"github.com/united-manufacturing-hub/oid/pkg/oid"
	"github.com/united-manufacturing-hub/oid/pkg/oid/100_managementConsole"
)

const (
	id = 100
)

func GetMgmtConsoleCertificateAsn10id() []int {
	return oid.ConcatCopyPreAllocate([][]int{managementConsole.GetMgmtConsoleAsn10id(), {id}})
}
