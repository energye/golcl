//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"unsafe"

	. "github.com/energye/golcl/lcl/api"
	. "github.com/energye/golcl/lcl/types"
)

func messageCallbackProc(f uintptr, msg uintptr) uintptr {
	v := PtrToElementValue(f)
	if v != nil {
		v.(TWndProcEvent)(
			(*TMessage)(unsafe.Pointer(msg)),
		)
	}
	return 0
}
