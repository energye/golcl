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
	Qac                           = &queueAsyncCall{id: 0, calls: sync.Map{}}
	applicationQueueAsyncCallFunc *dylib.LazyProc
)

func ApplicationQueueAsyncCallInit() {
	applicationQueueAsyncCallFunc = api.GetLibLCL().NewProc("SetApplicationQueueAsyncCallFunc")
	applicationQueueAsyncCallFunc.Call(applicationQueueAsyncCallEvent)
}

func applicationQueueAsyncCallProc(id uintptr) uintptr {
	Qac.call(id)
	return 0
}

// 队列异步调用函数 id:事件id
type QacFn func(id int)

type QueueCall struct {
	IsSync bool
	Fn     QacFn
	Wg     *sync.WaitGroup
}

type queueAsyncCall struct {
	id    uintptr
	calls sync.Map
}

func QueueAsyncCall(fn QacFn) int {
	id := Qac.Set(&QueueCall{
		IsSync: false,
		Fn:     fn,
	})
	api.GetLazyProc("CEFApplication_QueueAsyncCall").Call(id)
	return int(id)
}

func QueueSyncCall(fn QacFn) int {
	qc := &QueueCall{
		IsSync: true,
		Fn:     fn,
		Wg:     &sync.WaitGroup{},
	}
	qc.Wg.Add(1)
	id := Qac.Set(qc)
	api.GetLazyProc("CEFApplication_QueueAsyncCall").Call(id)
	qc.Wg.Wait()
	qc.Fn = nil
	qc.Wg = nil
	qc = nil
	return int(id)
}

func (m *queueAsyncCall) call(id uintptr) {
	if call, ok := m.calls.LoadAndDelete(id); ok {
		qc := call.(*QueueCall)
		if qc.IsSync {
			qc.Fn(int(id))
			qc.Wg.Done()
		} else {
			qc.Fn(int(id))
		}
	}
}
func (m *queueAsyncCall) Set(fn *QueueCall) uintptr {
	if m.id >= math.MaxUint {
		m.id = 0
	}
	m.id++
	m.calls.Store(m.id, fn)
	return m.id
}
