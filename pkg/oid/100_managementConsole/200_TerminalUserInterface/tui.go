package tui

import (
	"github.com/united-manufacturing-hub/oid/pkg/oid"
	"github.com/united-manufacturing-hub/oid/pkg/oid/100_managementConsole"
)

const (
	id = 200
)

func GetMgmtConsoleTUIAsn10id() []int {
	return oid.MergeSlices([][]int{managementConsole.GetMgmtConsoleAsn10id(), {id}})
}
