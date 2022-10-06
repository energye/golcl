//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package lcl

//应用主线程异步回调
import (
	"github.com/energye/golcl/dylib"
	"github.com/energye/golcl/lcl/api"
	"math"
	"sync"
)

var (
	qac                           = &queueAsyncCall{id: 0, calls: sync.Map{}}
	applicationQueueAsyncCallFunc *dylib.LazyProc
)

func ApplicationQueueAsyncCallInit() {
	applicationQueueAsyncCallFunc = api.GetLibLCL().NewProc("SetApplicationQueueAsyncCallFunc")
	applicationQueueAsyncCallFunc.Call(applicationQueueAsyncCallEvent)
}

func applicationQueueAsyncCallProc(id uintptr) uintptr {
	qac.call(id)
	return 0
}

// 队列异步调用函数 id:事件id
type QacFn func(id int)

type queueCall struct {
	isSync bool
	fn     QacFn
	wg     *sync.WaitGroup
}

type queueAsyncCall struct {
	id    uintptr
	calls sync.Map
}

func QueueAsyncCall(fn QacFn) int {
	id := qac.set(&queueCall{
		isSync: false,
		fn:     fn,
	})
	api.GetLazyProc("CEFApplication_QueueAsyncCall").Call(id)
	return int(id)
}

func QueueSyncCall(fn QacFn) int {
	qc := &queueCall{
		isSync: true,
		fn:     fn,
		wg:     &sync.WaitGroup{},
	}
	qc.wg.Add(1)
	id := qac.set(qc)
	api.GetLazyProc("CEFApplication_QueueAsyncCall").Call(id)
	qc.wg.Wait()
	qc.fn = nil
	qc.wg = nil
	qc = nil
	return int(id)
}

func (m *queueAsyncCall) call(id uintptr) {
	if call, ok := m.calls.LoadAndDelete(id); ok {
		qc := call.(*queueCall)
		if qc.isSync {
			qc.fn(int(id))
			qc.wg.Done()
		} else {
			qc.fn(int(id))
		}
	}
}
func (m *queueAsyncCall) set(fn *queueCall) uintptr {
	if m.id >= math.MaxUint {
		m.id = 0
	}
	m.id++
	m.calls.Store(m.id, fn)
	return m.id
}
