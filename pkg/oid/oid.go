package oid

// ianaRoot is the root of the IANA OID tree (iso.org.dod.internet.private.enterprise (1.3.6.1.4.1))
var ianaRoot = [6]int{1, 3, 6, 1, 4, 1}
var umhPen = [1]int{59193}

func GetIanaRoot() []int {
	return ianaRoot[:]
}

func GetUMHPen() []int {
	return umhPen[:]
}
