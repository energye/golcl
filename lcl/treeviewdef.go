package lcl

import "github.com/energye/golcl/emfs"

// 文件流加载。
func (s *TTreeView) LoadFromBytes(data []byte) {
	if len(data) == 0 {
		return
	}
	mem := NewMemoryStreamFromBytes(data)
	defer mem.Free()
	mem.SetPosition(0)
	s.LoadFromStream(mem)
}

// 从FS文件加载。
func (m *TTreeView) LoadFromFSFile(Filename string) error {
	bytes, err := emfs.GetResources(Filename)
	if err != nil {
		return err
	}
	m.LoadFromBytes(bytes)
	return nil
}
