package certificates

import (
	oid "oid/cmd"
	"oid/cmd/managementConsole"
)

const (
	caCertificateId     = 100
	techCertificateId   = 200
	deviceCertificateId = 300
)

func GetCaCertificateAsn10id() []int {
	return oid.ConcatCopyPreAllocate([][]int{managementConsole.GetMgmtConsoleAsn10id(), {caCertificateId}})
}

func GetTechCertificateAsn10id() []int {
	return oid.ConcatCopyPreAllocate([][]int{managementConsole.GetMgmtConsoleAsn10id(), {techCertificateId}})
}

func GetDeviceCertificateAsn10id() []int {
	return oid.ConcatCopyPreAllocate([][]int{managementConsole.GetMgmtConsoleAsn10id(), {deviceCertificateId}})
}
