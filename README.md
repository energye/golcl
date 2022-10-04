
----
## 基于 [govcl](https://gitee.com/ying32/govcl/) 库 出处 [Github](https://github.com/ying32/govcl) | [Gitee](https://gitee.com/ying32/govcl)
#### *.为配合 [energy](https://github.com/energye/energy) 基于CEF的GUI开源框架顾修改和增加了以下内容
#### *.增加资源文件在embed.FS读取和使用
#### *.增加使用embed.FS内置liblcl库,运行时自动释放至系统用户目录：用户目录/golcl/liblcl.dll
#### *.修改所有库使用初始化,为手动调用inits.Init(libs *embed.FS, resource *embed.FS) 用于导入liblcl库和资源文件
#### *.修改二进制类库加载和init函数的初始化

```golang
//示例：inits.Init(emfs libs,emfs source)支持embed.FS
package main

import (
	"embed"
	"github.com/energye/golcl/inits"
	"github.com/energye/golcl/lcl"
)

//go:embed resource
var resource embed.FS

//go:embed libs
var libs embed.FS

func main() {
	inits.Init(&libs, &resource)
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&ClientForm, true)
	lcl.Application.Run()
}
```

----

### golcl

跨平台的Golang GUI库, 核心绑定自 [Lazarus](https://www.lazarus-ide.org/) 创建的通用跨平台GUI库 [liblcl](https://github.com/energye/liblcl) 。

**golcl是一个原生GUI库，不是基于HTML，更不是DirectUI库，一切以实用为主。**

**中文全称：`Go语言可视化组件库`；英文全称：`Go Language Visual Component Library`**

*golcl最低要求go1.9.2。*

[截图查看](https://z-kit.cc/screenshot.html) |
[中文文档](https://gitee.com/ying32/golcl/wikis/pages) |
[更新日志](https://z-kit.cc/changelog.html) |
[视频教程(非官方)](https://video.0-w.cc/videos/1) |
[赞助govcl](https://z-kit.cc/sponsor.html)

----

### 支持的平台
**Windows** | **Linux** | **macOS**

> 如果你想要支持linux arm及linux 32bit则需要自己编译对应的liblcl二进制。

----

### 预编译GUI库二进制下载（[源代码](https://github.com/ying32/liblcl)）
[![liblcl](https://img.shields.io/github/downloads/ying32/govcl/latest/liblcl-2.2.0.zip.svg)](https://github.com/ying32/govcl/releases/download/v2.2.0/liblcl-2.2.0.zip)


> go get -u github.com/energye/golcl
>
> 注：也可用go module方式拉取，在go.mod中配置 【 go get -v -d ./... 或 go mod tidy 】
> 
> 参考 [samples](https://github.com/energye/golcl/tree/main/samples)
----

#### 复制对应的二进制

* Windows: 根据编译的二进制是32还是64位的，复制对应的`liblcl.dll`到当前可执行文件目录或embed.FS映射目录。
  * Go环境变量： `GOARCH = amd64 386` `GOOS = windows` `CGO_ENABLED=0`

* Linux: 复制`liblcl.so`当前可执行文件目录下或embed.FS映射目录。
  * Go环境变量： `GOARCH = amd64` `GOOS = linux` `CGO_ENABLED=1`

* MacOS: 复制`liblcl.dylib`当前可执行文件目录下（MacOS下注意：需要自行创建info.plist文件），或者参考：[MacOS上应用打包](https://github.com/energye/golcl/blob/main/pkgs/macapp/README.md)
  * Go环境变量： `GOARCH = amd64` `GOOS = darwin` `CGO_ENABLED=1`


注：这里的`当前可执行文件目录`指的是你当前编译的项目生成的可执行文件位置。

----

### 注意:

**特别注意：所有UI组件都是非线程/协程安全的，当在goroutine中使用时，请使用[lcl.ThreadSync](https://gitee.com/ying32/golcl/wikis/pages?sort_id=976890&doc_id=102420)来同步更新到UI上。**

**特别注意2：如果你使用go>=1.15编译Windows可执行文件，则必须则必须使用`-buildmode=exe`编译选项，不然会有错误。**

### API文档

* [Lazarus LCL组件WIKI](http://wiki.freepascal.org/LCL_Components)
* [Windows API文档](https://msdn.microsoft.com/zh-cn/library/ms123401.aspx)

----

![jetbrains](https://z-kit.cc/assets/images/jetbrains.png)   
[鸣谢jetbrains](https://www.jetbrains.com/?from=golcl)  