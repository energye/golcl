package winRes

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/energye/golcl/tools/command"
	"io/fs"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"text/template"
)

type SYSO struct {
	FileVersion      string `json:"file_version"`
	FileRCVersion    string `json:"file_rc_version"`
	ProductVersion   string `json:"product_version"`
	ProductRCVersion string `json:"product_rc_version"`
	Manifest         string `json:"manifest"`
	MainIcon         string `json:"main_icon"`
	AppName          string `json:"app_name"`
	CompanyName      string `json:"company_name"`
	FileDescription  string `json:"file_description"`
	InternalName     string `json:"internal_name"`
	LegalCopyright   string `json:"legal_copyright"`
	OriginalFilename string `json:"original_filename"`
	ProductName      string `json:"product_name"`
	CMDDir           string //命令行执行并生成文件的根目录
	IconName         string //指定的xxxx.ico名称，为空采用AppName名称，在[CMDDir]/resources目录查找该文件
}

//默认生成图标的方式 syso ,需要安装 windres
func NewSYSO() *SYSO {
	syso := &SYSO{
		FileVersion:      "1.0.0.0",
		ProductVersion:   "1.0.0.0",
		Manifest:         "app.Manifest",
		MainIcon:         "app.ico",
		AppName:          "app",
		CompanyName:      "<TODO:Company Name>",
		FileDescription:  "",
		InternalName:     "",
		LegalCopyright:   "<TODO:Copyright>",
		OriginalFilename: "",
		ProductName:      "",
	}
	return syso
}

var inputMsg = []string{
	"应用名称(app name) default【app】:",
	"应用版本号(app version) default【1.0.0.0】:",
	"产品名称(product name) default【】:",
	"产品版本号(product version) default【1.0.0.0】:",
	"文件描述(file description) default【】:",
	"公司名称(company name) default【<TODO:Company Name>】:",
	"内部名称(internal name) default【】:",
	"版权信息(copyright) default【<TODO:Copyright>】:",
	"原始文件名称(original filename) default【】:",
}

func (m *SYSO) RC() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("生成EXE-icon程序，回车继续,N退出")
	var i = 0
	var b = false
	for {
		if b {
			fmt.Print(inputMsg[i])
		}
		line, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		var text = strings.Trim(string(line), " ")
		if !b {
			if "N" == strings.ToUpper(text) {
				return nil
			}
		}
		if b {
			switch i {
			case 0:
				if text == "" {
					text = "app"
				}
				m.AppName = text
				m.Manifest = text + ".Manifest"
				if m.IconName != "" {
					m.MainIcon = "resources/" + m.IconName
				} else {
					m.MainIcon = "resources/" + text + ".ico"
				}
			case 1:
				if text == "" {
					text = "1.0.0.0"
				}
				m.FileVersion = text
				m.FileRCVersion = strings.ReplaceAll(text, ".", ",")
			case 2:
				m.ProductName = text
			case 3:
				if text == "" {
					text = "1.0.0.0"
				}
				m.ProductVersion = text
				m.ProductRCVersion = strings.ReplaceAll(text, ".", ",")
			case 4:
				m.FileDescription = text
			case 5:
				if text == "" {
					text = "<TODO:Company Name>"
				}
				m.CompanyName = text
			case 6:
				m.InternalName = text
			case 7:
				if text == "" {
					text = "<TODO:Company Name>"
				}
				m.LegalCopyright = text
			case 8:
				m.OriginalFilename = text
			}
			i++
			if i >= len(inputMsg) {
				break
			}
		}
		b = true
	}
	return m.Gen()
}

func (m *SYSO) Gen() error {
	var (
		err     error
		icoFile string
	)
	if m.CMDDir == "" {
		exeRoot := os.Args[0]
		if runtime.GOOS == "windows" {
			exeRoot = exeRoot[:strings.LastIndex(exeRoot, "\\")]
		} else {
			exeRoot = exeRoot[:strings.LastIndex(exeRoot, "/")]
		}
		m.CMDDir = exeRoot
	}
	if m.IconName != "" {
		icoFile = m.CMDDir + "/resources/" + m.IconName + ".ico"
	} else {
		icoFile = m.CMDDir + "/resources/" + m.AppName + ".ico"
	}
	var rcFile = m.CMDDir + "/" + m.AppName + ".rc"
	var maifesFile = m.CMDDir + "/" + m.AppName + ".manifest"
	os.Remove(rcFile)
	os.Remove(maifesFile)
	fmt.Println("================================================")
	fmt.Println("================================================")
	jsn, _ := json.MarshalIndent(m, "", "\t")
	fmt.Println("================RC信息:\n", string(jsn))
	fmt.Println("================生成RC文件================")
	temp, _ := template.New("rc").Parse(rcTemplate)
	text, _ := textTemplate(temp, m)
	fmt.Println("================RC文件信息:\n", text)
	fmt.Println("================写入RC文件信息================")
	ioutil.WriteFile(rcFile, []byte(text), fs.ModePerm)
	fmt.Println("================写入RC文件信息完成================")
	fmt.Println("================================================")
	fmt.Println("================================================")
	fmt.Println("================生成manifest文件================")
	temp, _ = template.New("manifest").Parse(manifestTemplate)
	text, _ = textTemplate(temp, m)
	fmt.Println("================manifest文件信息:\n", text)
	fmt.Println("================写入manifes文件信息================")
	ioutil.WriteFile(maifesFile, []byte(text), fs.ModePerm)
	fmt.Println("================写入manifes文件信息完成================")
	fmt.Println("================================================")
	fmt.Println("================================================")
	fmt.Println("================生成syso文件================")
	if m.CMDDir != "" {
		if m.IconName != "" {
			icoFile = m.CMDDir + "/resources/" + m.IconName
		} else {
			icoFile = m.CMDDir + "/resources/" + m.AppName + ".ico"
		}
	}
	if !isExist(icoFile) {
		os.Remove(rcFile)
		os.Remove(maifesFile)
		fmt.Println(icoFile + " 文件未找到")
		return errors.New(icoFile + " 文件未找到")
	}
	cmd := command.NewCMD()
	cmd.MessageCallback = func(s []byte, e error) {
		fmt.Println("GEN - CMD", s, " error", e)
		err = e
	}
	if m.CMDDir != "" {
		cmd.Dir = m.CMDDir
	}
	if err == nil {
		var i386SYSOFile = m.AppName + "_windows_386.syso"
		var argsi386 = []string{"/c", "windres.exe", "-i", rcFile, "-o", i386SYSOFile, "-F", "pe-i386"}
		cmd.Command("cmd.exe", argsi386...)
		cmd.Close()
		fmt.Println("\tsyso:", i386SYSOFile)
	}
	if err == nil {
		var x8664SYSOFile = m.AppName + "_windows_amd64.syso"
		var argsx8664 = []string{"/c", "windres.exe", "-i", rcFile, "-o", x8664SYSOFile, "-F", "pe-x86-64"}
		cmd.Command("cmd.exe", argsx8664...)
		cmd.Close()
		fmt.Println("\tsyso:", x8664SYSOFile)
	}
	fmt.Println("================生成syso文件完成================")
	os.Remove(rcFile)
	os.Remove(maifesFile)
	return err
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}

func textTemplate(template *template.Template, data interface{}) (string, error) {
	buf := &bytes.Buffer{}
	if err := template.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

var rcTemplate = `1 MANIFEST "{{.Manifest}}"
MAINICON ICON "{{.MainIcon}}"

1 VERSIONINFO
FILEVERSION {{.FileRCVersion}}
PRODUCTVERSION {{.ProductRCVersion}}
FILEOS 0x4
FILETYPE 0x1

BEGIN
    BLOCK "StringFileInfo"
    BEGIN
        BLOCK "040904B0"
        BEGIN
			VALUE "CompanyName", "{{.CompanyName}}"
			VALUE "FileDescription", "{{.FileDescription}}"
			VALUE "FileVersion", "{{.FileVersion}}"
			VALUE "InternalName", "{{.InternalName}}"
			VALUE "LegalCopyright", "{{.LegalCopyright}}"
			VALUE "OriginalFilename", "{{.OriginalFilename}}"
			VALUE "ProductName", "{{.ProductName}}"
			VALUE "ProductVersion", "{{.ProductVersion}}"
        END
    END
    BLOCK "VarFileInfo"
    BEGIN
            VALUE "Translation", 0x0409, 0x04B0
    END
END
`

var manifestTemplate = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0" xmlns:asmv3="urn:schemas-microsoft-com:asm.v3">
  <asmv3:application>
    <asmv3:windowsSettings xmlns="http://schemas.microsoft.com/SMI/2005/WindowsSettings">
      <!--<dpiAware>True/PM</dpiAware>-->
    </asmv3:windowsSettings>
  </asmv3:application>
  <dependency>
    <dependentAssembly>
      <assemblyIdentity
        type="win32"
        name="Microsoft.Windows.Common-Controls"
        version="6.0.0.0"
        publicKeyToken="6595b64144ccf1df"
        language="*"
        processorArchitecture="*"/>
    </dependentAssembly>
  </dependency>
  <trustInfo xmlns="urn:schemas-microsoft-com:asm.v3">
    <security>
      <requestedPrivileges>
        <requestedExecutionLevel
          level="requireAdministrator"
          uiAccess="false"
        />
        </requestedPrivileges>
    </security>
  </trustInfo>
<compatibility xmlns="urn:schemas-microsoft-com:compatibility.v1"> 
	<application> 
		<!--The ID below indicates app support for Windows Vista -->
		<supportedOS Id="{e2011457-1546-43c5-a5fe-008deee3d3f0}"/> 
		<!--The ID below indicates app support for Windows 7 -->
		<supportedOS Id="{35138b9a-5d96-4fbd-8e2d-a2440225f93a}"/>
		<!--The ID below indicates app support for Windows 8 -->
		<supportedOS Id="{4a2f28e3-53b9-4441-ba9c-d69d4a4a6e38}"/>
		<!--The ID below indicates app support for Windows 8.1 -->
		<supportedOS Id="{1f676c76-80e1-4239-95bb-83d0f6d0da78}"/>
		<!--The ID below indicates app support for Windows 10 -->
		<supportedOS Id="{8e0f7a12-bfb3-4fe8-b9a5-48fd50a15a9a}"/>			
	</application> 
</compatibility>
</assembly>
`
