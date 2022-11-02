
----
## 基于 [govcl](https://gitee.com/ying32/govcl/) 库, 出处 [Github](https://github.com/ying32/govcl) | [Gitee](https://gitee.com/ying32/govcl)
#### 为配合 [Energy](https://github.com/energye/energy) 基于CEF的桌面客户端框架,修改和增加做出了部分源码修改
#### 增加资源文件在embed.FS读取和使用
#### 增加使用embed.FS内置liblcl库,运行时自动释放至系统用户目录：用户目录/golcl/liblcl.dll
#### 修改所有库使用初始化,为手动调用inits.Init(libs *embed.FS, resource *embed.FS) 用于导入liblcl库和资源文件
#### 修改二进制类库加载和init函数的初始化
#### 其它修改详见源码

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

### Golcl

跨平台的Golang GUI库, 核心绑定自 [Lazarus](https://www.lazarus-ide.org/) 创建的通用跨平台GUI库 [liblcl](https://github.com/energye/liblcl) 。

**Golcl是一个原生GUI库，不是基于HTML，更不是DirectUI库，一切以实用为主。**

**中文全称：`Go语言可视化组件库`；英文全称：`Go Language Visual Component Library`**

*Golcl最低要求go1.9.2。*

[截图查看](https://z-kit.cc/screenshot.html) |
[中文文档](https://gitee.com/ying32/golcl/wikis/pages) |
[更新日志](https://z-kit.cc/changelog.html) |
[视频教程(非官方)](https://video.0-w.cc/videos/1) |
[赞助govcl](https://z-kit.cc/sponsor.html)

----

### 支持的平台
**Windows** | **Linux** | **macOS**

> go get -u github.com/energye/golcl
>
> 注：也可用go module方式拉取，在go.mod中配置 【 go get -v -d ./... 或 go mod tidy 】
> 
> 参考 [samples](https://github.com/energye/golcl/tree/main/samples)
----
[Apache License 2.0](https://github.com/ying32/govcl/blob/master/LICENSE)

---