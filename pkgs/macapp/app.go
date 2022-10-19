package macapp

var (
	//1. macos调式时临时创建一个符合macapp的程序包
	//
	//2. 如果基于cef，需要指定cef frameworks 根目录【/homt/xxx/cef_binary_xxxxxxx_macosx64/Release】
	MacApp = &macApp{}
)

type macApp struct {
	execName             string
	execFile             string
	macContentsDir       string
	macOSDir             string
	macResources         string
	lclLibFileName       string
	plistFileName        string
	pkgInfoFileName      string
	macAppFrameworksDir  string
	isCEF                bool
	isMain               bool
	cefFrameworksDir     string
	browseSubprocessPath string
	lsUIElement          string
}
