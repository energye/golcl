package rtl

import . "github.com/energye/golcl/lcl/types"

func MakeWord(A, B uint8) uint16 {
	return uint16(A | B<<8)
}

func MakeLong(A, B uint16) int32 {
	return int32(A | B<<16)
}

func HiWord(L uint32) uint16 {
	return uint16(L >> 16)
}

func HiByte(W uint8) uint8 {
	return W >> 8
}

func MakeWParam(l, h uint16) WPARAM {
	return WPARAM(MakeLong(l, h))
}

func MakeLParam(l, h uint16) LPARAM {
	return LPARAM(MakeLong(l, h))
}

func MakeLResult(l, h uint16) LRESULT {
	return LRESULT(MakeLong(l, h))
}

func PointToLParam(P TPoint) LPARAM {
	return LPARAM((P.X & 0x0000FFFF) | (P.Y << 16))
}
