package api

import "github.com/energye/dylib"

func IsNilProcApi(procPtr *dylib.LazyProc) bool {
	return procPtr == nil
}
