package emfs

import (
	"embed"
	"io/ioutil"
	"os"
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

func IsExist(file string) bool {
	if GetResourcesFS() == nil {
		return false
	}
	if fs, err := GetResourcesFS().Open(file); err != nil {
		return false
	} else {
		_, err := fs.Stat()
		if os.IsExist(err) {
			return true
		} else if os.IsNotExist(err) {
			return false
		}
		return true
	}
}

func GetResources(file string) ([]byte, error) {
	if GetResourcesFS() != nil {
		return GetResourcesFS().ReadFile(file)
	} else {
		return ioutil.ReadFile(file)
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
