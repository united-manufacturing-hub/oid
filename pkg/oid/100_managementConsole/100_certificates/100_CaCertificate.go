package certificates

import (
	"github.com/united-manufacturing-hub/oid/pkg/oid"
)

const (
	caCertificateId = 100
)

func GetCaCertificateAsn10id() []int {
	return oid.MergeSlices([][]int{GetMgmtConsoleCertificateAsn10id(), {caCertificateId}})
}
