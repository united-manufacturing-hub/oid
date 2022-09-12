package oid

import "crypto/x509/pkix"

func MergeSlices(slices [][]int) []int {
	var totalLen int
	for _, s := range slices {
		totalLen += len(s)
	}
	tmp := make([]int, totalLen)
	var i int
	for _, s := range slices {
		i += copy(tmp[i:], s)
	}
	return tmp
}

func ToPkixExtension(oid []int, critical bool, value []byte) pkix.Extension {
	return pkix.Extension{
		Id:       oid,
		Critical: critical,
		Value:    value,
	}
}
