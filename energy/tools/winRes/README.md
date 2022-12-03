applogo.ico生成工具

2种生成方式

1.syso 需要安装 windres
```go
func syso() error {
	syso := winRes.NewSYSO()
	//命令行运行的根目录
	//在该目录下生成xxx.syso文件
	//在该目录下resources/icon.ico文件
	syso.CMDDir = "E:\\SWT\\gopath\\src\\gitee.com\\snxamdf\\golcl" 
	//指针的icon图标名称
	syso.IconName = "icon.ico"
	//生成文件,命令行提示输入信息
	return syso.RC()
}
```

2.rsrc 使用exe目录下解压缩 rsrc.exe.rar ,或自行编译到 github.com/akavel/rsrc
```go
func rsrc() error {
	rsrc := winRes.NewRSRC()
	//指定rsrc.exe根目录
	rsrc.RSRCPath = "E:\\SWT\\gopath\\src\\gitee.com\\snxamdf\\golcl\\exe\\rsrc.exe"
    //命令行运行的根目录
    //在该目录下生成xxx.syso文件
    //在该目录下resources/icon.ico文件
	rsrc.CMDDir = "E:\\SWT\\gopath\\src\\gitee.com\\snxamdf\\golcl"
	//在[CMDDir]/resources目录查找icon和manifest.xml，或自行指定目录
	//rsrc.IconPath = "icon.ico"
	//rsrc.ManifestPath = "manifestPath.xml"
	return rsrc.Gen()
}
```