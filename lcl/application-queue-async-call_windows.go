//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package lcl

import "syscall"

var (
	applicationQueueAsyncCallEvent = syscall.NewCallback(applicationQueueAsyncCallProc)
)
