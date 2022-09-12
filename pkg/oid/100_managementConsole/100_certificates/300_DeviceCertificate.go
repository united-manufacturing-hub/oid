package certificates

import (
	"github.com/united-manufacturing-hub/oid/pkg/oid"
)

const (
	deviceCertificateId = 300
)

func GetDeviceCertificateAsn10id() []int {
	return oid.MergeSlices(
		[][]int{
			GetMgmtConsoleCertificateAsn10id(),
			{deviceCertificateId}})
}
