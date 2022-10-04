package lcl

import "github.com/energye/golcl/emfs"

func (p *TPngImage) LoadFromBytes(data []byte) {
	if len(data) == 0 {
		return
	}
	mem := NewMemoryStreamFromBytes(data)
	defer mem.Free()
	mem.SetPosition(0)
	p.LoadFromStream(mem)
}

// 从FS文件加载。
func (m *TPngImage) LoadFromFSFile(Filename string) error {
	bytes, err := emfs.GetResources(Filename)
	if err != nil {
		return err
	}
	m.LoadFromBytes(bytes)
	return nil
}
