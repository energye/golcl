//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// energy extension
//
//govcl/vcl/api/dllimports -> golcl/lcl/api/dllimports/energy_imports.go

package dllimports

func NewEnergyImport(name string, addr ProcAddr) *ImportTable {
	r := &ImportTable{}
	r.name = name
	r.addr = addr
	return r
}

func (m *ImportTable) Name() string {
	return m.name
}

func (m *ImportTable) Addr() ProcAddr {
	return m.addr
}

// ImportDefFunc Get Import
func ImportDefFunc(lib DLL, importTable []*ImportTable, index int) ProcAddr {
	return internalGetImportFunc(lib, importTable, index)
}

func DllImportDefs() []*ImportTable {
	return dllImportDefs
}

func DllImports() []*ImportTable {
	return dllImports
}
