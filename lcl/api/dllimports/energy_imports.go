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

//Energy extend
var energyImportDefs []*ImportTable

func NewEnergyImport(name string, addr ProcAddr) *ImportTable {
	r := &ImportTable{}
	r.name = name
	r.addr = addr
	return r
}

//Energy Set Import
func SetEnergyImportDefs(eis []*ImportTable) {
	energyImportDefs = eis
}

//Energy Get Import Addr
func GetEnergyImportDefFunc(uiLib DLL, index int) ProcAddr {
	return internalGetImportFunc(uiLib, energyImportDefs, index)
}

//Energy Get Import
func GetEnergyImport(index int) *ImportTable {
	return energyImportDefs[index]
}

func GetEnergyImports() []*ImportTable {
	return energyImportDefs
}

func (m *ImportTable) Name() string {
	return m.name
}

func (m *ImportTable) Addr() ProcAddr {
	return m.addr
}
