package lcl

import (
	"github.com/energye/golcl/emfs"
	. "github.com/energye/golcl/lcl/api"
)

// 新建Delphi/Lazarus内存流来自Go字节数组。
//
// New Delphi/Lazarus memory stream from Go byte array.
func NewMemoryStreamFromBytes(data []byte) *TMemoryStream {
	m := NewMemoryStream()
	m.Write(data)
	return m
}

// 读数据
//
// Read Data.
func (m *TMemoryStream) Read(count int32) (int32, []byte) {
	return MemoryStream_Read(m.instance, count)
}

// 写数据
//
// Write Data.
func (m *TMemoryStream) Write(buffer []byte) int32 {
	return MemoryStream_Write(m.instance, buffer)
}

func (j *TMemoryStream) LoadFromBytes(data []byte) {
	if len(data) == 0 {
		return
	}
	mem := NewMemoryStreamFromBytes(data)
	defer mem.Free()
	mem.SetPosition(0)
	j.LoadFromStream(mem)
}

// 从FS文件加载。
func (m *TMemoryStream) LoadFromFSFile(Filename string) error {
	bytes, err := emfs.GetResources(Filename)
	if err != nil {
		return err
	}
	m.LoadFromBytes(bytes)
	return nil
}
