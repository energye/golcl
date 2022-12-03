package emfs

import (
	"embed"
	"io/ioutil"
)

var (
	resourcesFS *embed.FS
	libsFS      *embed.FS
)

func SetLibsFS(lib *embed.FS) {
	libsFS = lib
}

func SetResourcesFS(resource *embed.FS) {
	resourcesFS = resource
}

func GetLibsFS() *embed.FS {
	return libsFS
}

func GetResourcesFS() *embed.FS {
	return resourcesFS
}

func GetResources(fileName string) ([]byte, error) {
	if GetResourcesFS() != nil {
		return GetResourcesFS().ReadFile(fileName)
	} else {
		return ioutil.ReadFile(fileName)
	}
}

func GetLibs(fileName string) ([]byte, error) {
	if GetLibsFS() != nil {
		return GetLibsFS().ReadFile(fileName)
	} else {
		return ioutil.ReadFile(fileName)
	}
}

func SetEMFS(libs *embed.FS, resources *embed.FS) {
	SetLibsFS(libs)
	SetResourcesFS(resources)
}
