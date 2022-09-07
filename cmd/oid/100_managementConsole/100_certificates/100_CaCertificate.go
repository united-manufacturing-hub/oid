package certificates

import (
	"github.com/united-manufacturing-hub/oid/cmd/oid"
)

const (
	caCertificateId = 100
)

func GetCaCertificateAsn10id() []int {
	return oid.ConcatCopyPreAllocate([][]int{GetMgmtConsoleCertificateAsn10id(), {caCertificateId}})
}
