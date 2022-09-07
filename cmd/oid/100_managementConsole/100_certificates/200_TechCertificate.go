package certificates

import (
	"github.com/united-manufacturing-hub/oid/cmd/oid"
)

const (
	techCertificateId = 200
)

func GetTechCertificateAsn10id() []int {
	return oid.ConcatCopyPreAllocate([][]int{GetMgmtConsoleCertificateAsn10id(), {techCertificateId}})
}
