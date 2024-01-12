//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

// Callback 回調事件
var Callback callback

type callback uintptr

func (*callback) Event() uintptr {
	return eventCallback
}

func (*callback) Message() uintptr {
	return messageCallback
}

func (*callback) ThreadSync() uintptr {
	return threadSyncCallback
}

func (*callback) RequestCallCreateParams() uintptr {
	return requestCallCreateParamsCallback
}

func (*callback) RemoveEven() uintptr {
	return removeEventCallback
}
