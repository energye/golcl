package consts

import (
	"fmt"
	"github.com/energye/golcl/tools/homedir"
	"os"
	"path/filepath"
	"strings"
)

const (
	RootPath = "golcl"
)

var (
	ExePath string //执行文件目录
	//用户目录
	DIR, err = homedir.Dir()
	//目录结构分割符
	Separator = string(filepath.Separator)
	//完整的用户目录
	HomeDir = fmt.Sprintf("%v%v%v", DIR, Separator, RootPath)
)

func init() {
	ExePath = os.Args[0]
	ExePath = ExePath[:strings.LastIndex(ExePath, Separator)]
	if err != nil {
		panic(err)
	}
}
