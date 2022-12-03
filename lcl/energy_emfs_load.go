//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// energy extension
//
//govcl/vcl -> golcl/lcl/energy_emfs_load.go
package lcl

import "github.com/energye/golcl/energy/emfs"

// 从FS文件加载。
func (i *TIcon) LoadFromFSFile(Filename string) error {
	bytes, err := emfs.GetResources(Filename)
	if err != nil {
		return err
	}
	i.LoadFromBytes(bytes)
	return nil
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

// 从FS文件加载。
func (m *TBitmap) LoadFromFSFile(Filename string) error {
	bytes, err := emfs.GetResources(Filename)
	if err != nil {
		return err
	}
	m.LoadFromBytes(bytes)
	return nil
}

// 从FS文件加载。
func (m *TGIFImage) LoadFromFSFile(Filename string) error {
	bytes, err := emfs.GetResources(Filename)
	if err != nil {
		return err
	}
	m.LoadFromBytes(bytes)
	return nil
}

func (m *TGraphic) LoadFromBytes(data []byte) {
	if len(data) == 0 {
		return
	}
	mem := NewMemoryStreamFromBytes(data)
	defer mem.Free()
	mem.SetPosition(0)
	m.LoadFromStream(mem)
}

// 从FS文件加载。
func (m *TJPEGImage) LoadFromFSFile(Filename string) error {
	bytes, err := emfs.GetResources(Filename)
	if err != nil {
		return err
	}
	m.LoadFromBytes(bytes)
	return nil
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

// 从FS文件加载。
func (m *TPicture) LoadFromFSFile(Filename string) error {
	bytes, err := emfs.GetResources(Filename)
	if err != nil {
		return err
	}
	m.LoadFromBytes(bytes)
	return nil
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

// 从FS文件加载。
func (m *TStringList) LoadFromFSFile(Filename string) error {
	bytes, err := emfs.GetResources(Filename)
	if err != nil {
		return err
	}
	m.LoadFromBytes(bytes)
	return nil
}

// 文件流加载。
func (m *TTreeView) LoadFromBytes(data []byte) {
	if len(data) == 0 {
		return
	}
	mem := NewMemoryStreamFromBytes(data)
	defer mem.Free()
	mem.SetPosition(0)
	m.LoadFromStream(mem)
}

// 文件流加载。
func (m *TMemoryStream) LoadFromBytes(data []byte) {
	if len(data) == 0 {
		return
	}
	mem := NewMemoryStreamFromBytes(data)
	defer mem.Free()
	mem.SetPosition(0)
	m.LoadFromStream(mem)
}
