//----------------------------------------
//
// Copyright © sxm. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"unsafe"
)

// MemoryStream_Read 还需要待测试
func MemoryStream_Read(obj uintptr, count int32) (int32, []byte) {
	if count <= 0 {
		return 0, nil
	}
	bs := make([]byte, count)
	if IsNilProcApi(memoryStream_Read) {
		memoryStream_Read = liblcl.NewProc("MemoryStream_Read")
	}
	r, _, _ := memoryStream_Read.Call(obj, uintptr(unsafe.Pointer(&bs[0])), uintptr(count))
	return int32(r), bs
}

func MemoryStream_Write(obj uintptr, buffer []byte) int32 {
	if IsNilProcApi(memoryStream_Write) {
		memoryStream_Write = liblcl.NewProc("MemoryStream_Write")
	}
	r, _, _ := memoryStream_Write.Call(obj, uintptr(unsafe.Pointer(&buffer[0])), uintptr(len(buffer)))
	return int32(r)
}
