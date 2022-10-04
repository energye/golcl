//go:build windows
// +build windows

package win

func ToUInt64(r1, r2 uintptr) uint64 {
	ret := uint64(r2)
	ret = uint64(ret<<32) + uint64(r1)
	return ret
}

func UInt64To(val uint64) (uintptr, uintptr) {
	return uintptr(uint32(val)), uintptr(uint32(val >> 32))
}
