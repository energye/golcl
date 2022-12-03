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

type EnergyImportTable struct {
	importTable
}

func New(name string, addr ProcAddr) EnergyImportTable {
	r := EnergyImportTable{}
	r.name = name
	r.addr = addr
	return r
}

func (m *EnergyImportTable) SetName(name string) {
	m.name = name
}

func (m *EnergyImportTable) SetAddr(addr ProcAddr) {
	m.addr = addr
}

func ProcSize() int {
	return len(dllImports)
}

func RegisterEnergy(eis []EnergyImportTable) {
	for _, e := range eis {
		dllImports = append(dllImports, e.importTable)
	}
}
