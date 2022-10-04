// 在这里写你的事件

package main

import (
	"fmt"

	"github.com/energye/golcl/lcl"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnFormCreate(sender lcl.IObject) {
	for i := 0; i < 20; i++ {
		f.ValueListEditor1.Strings().Add(fmt.Sprintf("aa%d=%d", i, i))
	}

}
