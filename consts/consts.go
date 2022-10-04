package consts

import (
	"fmt"
	"github.com/energye/golcl/tools/homedir"
	"path/filepath"
)

const (
	RootPath = "golcl"
)

var (
	//用户目录
	DIR, err = homedir.Dir()
	//目录结构分割符
	Separator = string(filepath.Separator)
	//完整的用户目录
	HomeDir = fmt.Sprintf("%v%v%v", DIR, Separator, RootPath)
)

func init() {
	if err != nil {
		panic(err)
	}
}
