package api

import "github.com/energye/golcl/dylib"

func IsNilProcApi(procPtr *dylib.LazyProc) bool {
	return procPtr == nil
}
