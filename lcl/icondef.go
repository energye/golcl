//----------------------------------------
//
// Copyright © sxm. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/golcl/emfs"
	//. "github.com/energye/golcl/lcl/api"
)

func (i *TIcon) LoadFromBytes(data []byte) {
	if len(data) == 0 {
		return
	}
	mem := NewMemoryStreamFromBytes(data)
	defer mem.Free()
	mem.SetPosition(0)
	i.LoadFromStream(mem)
}

// 从FS文件加载。
func (i *TIcon) LoadFromFSFile(Filename string) error {
	bytes, err := emfs.GetResources(Filename)
	if err != nil {
		return err
	}
	i.LoadFromBytes(bytes)
	return nil
}
