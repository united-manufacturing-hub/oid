package certificates

import (
	"github.com/united-manufacturing-hub/oid/pkg/oid"
)

const (
	caCertificateId = 100
)

func GetCaCertificateAsn10id() []int {
	return oid.ConcatCopyPreAllocate([][]int{GetMgmtConsoleCertificateAsn10id(), {caCertificateId}})
}
