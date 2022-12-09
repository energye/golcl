## Energy 扩展说明

### 对支持Energy扩展增加和修改
```go
//修改后每个示例运行增加
func main() {
    inits.Init(nil, nil)
	...
}
```

#### 修改版本 v1.0.0

> 1. 修改 govcl to golcl
>> github.com/ying32/govcl
>>
>> to
>>
>> github.com/energye/golcl

```go
包-Package

新增 包
    1. golcl/energy
    2. golcl/tools/command
修改 包
    1. golcl/tools/winRes
移除 包

文件-File

新增 xxx.go
    1. golcl/lcl/api/energy_api.go
    2. golcl/lcl/energy_emfs_load.go
    3. golcl/lcl/api/dllimports
修改 xxx.go
    1. golcl/pkgs/macapp/app.go
    2. golcl/pkgs/macapp/deficon.go
    3. golcl/lcl/api/dylib.go
    4. golcl/lcl/rtl/version/init.go
    5. golcl/lcl/win/init.go
    6. golcl/lcl/init.go
    7. golcl/lcl/rtl/init_posix.go
    8. golcl/lcl/rtl/init_windows.go
    8. golcl/pkgs/libname/lib.go
    9. golcl/lcl/i18n/mlang.go
    10. golcl/pkgs/macapp/other.go
    11. golcl/pkgs/macapp/plist.go
    12. golcl/lcl/api/dllimports/dll_posix.go
移除 xxx.go
```

