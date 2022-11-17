//go:build darwin
// +build darwin

package macapp

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/energye/golcl/consts"
	"github.com/energye/golcl/pkgs/libname"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	mac_cef_helper          = "Helper"
	mac_cef_helper_Gpu      = "Helper (GPU)"
	mac_cef_helper_Plugin   = "Helper (Plugin)"
	mac_cef_helper_Renderer = "Helper (Renderer)"
	cef                     = "Chromium Embedded Framework.framework"
)

var helpers = []string{
	mac_cef_helper,
	mac_cef_helper_Gpu,
	mac_cef_helper_Plugin,
	mac_cef_helper_Renderer,
}

func init() {
	MacApp.execName = os.Args[0][strings.LastIndex(os.Args[0], "/")+1:]
	MacApp.macContentsDir = os.Args[0] + ".app/Contents"
	MacApp.macOSDir = MacApp.macContentsDir + "/MacOS"
	MacApp.macResources = MacApp.macContentsDir + "/Resources"
	MacApp.execFile = MacApp.macOSDir + "/" + MacApp.execName
	MacApp.lclLibFileName = MacApp.macOSDir + "/liblcl.dylib"
	MacApp.plistFileName = MacApp.macContentsDir + "/Info.plist"
	MacApp.pkgInfoFileName = MacApp.macContentsDir + "/PkgInfo"
	MacApp.macAppFrameworksDir = MacApp.macContentsDir + "/Frameworks"
	MacApp.lsUIElement = "false"
}

//func copyFile(src, dest string) error {
//	fileDest, err := os.Create(dest)
//	if err != nil {
//		return err
//	}
//	defer fileDest.Close()
//	fileSrc, err := os.Open(src)
//	if err != nil {
//		return err
//	}
//	defer fileSrc.Close()
//	_, err = io.Copy(fileDest, fileSrc)
//	return err
//}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

/*
 * 拷贝文件夹,同时拷贝文件夹中的文件
 * srcPath 需要拷贝的文件夹路径: /test
 * destPath 拷贝到的位置: /backup/
 */
func copyDir(srcPath string, destPath string) error {
	if srcInfo, err := os.Stat(srcPath); err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		if !srcInfo.IsDir() {
			e := errors.New("srcPath不是一个正确的目录！")
			return e
		}
	}
	if destInfo, err := os.Stat(destPath); err != nil {
		return err
	} else {
		if !destInfo.IsDir() {
			e := errors.New("destInfo不是一个正确的目录！")
			return e
		}
	}
	err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if !f.IsDir() {
			sPath := strings.Replace(path, "\\", "/", -1)
			destNewPath := strings.Replace(sPath, srcPath, destPath, -1)
			//fmt.Println("\n\t复制文件:  ", sPath, "\n\t到  ", destNewPath)
			copyFile(path, destNewPath)
		}
		return nil
	})
	return err
}

//生成目录并拷贝文件
func copyFile(src, dest string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()
	destSplitPathDirs := strings.Split(dest, "/")
	destSplitPath := ""
	for index, dir := range destSplitPathDirs {
		if index < len(destSplitPathDirs)-1 {
			destSplitPath = destSplitPath + dir + "/"
			b := fileExists(destSplitPath)
			if b == false {
				//fmt.Println("创建目录:" + destSplitPath)
				if err := os.Mkdir(destSplitPath, 0755); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
	dstFile, err := os.Create(dest)
	if err != nil {
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}

func (m *macApp) Init() {
	if strings.Contains(os.Args[0], ".app/Contents/MacOS") {
		return
	}
	m.isMain = true
	if m.createMacOSApp(m) {
		m.copyDylib()
		m.cefHelper()
		m.runMacOSApp()
	}
}

func (m *macApp) SetBaseCefFrameworksDir(s string) {
	m.cefFrameworksDir = s
}

func (m *macApp) SetBrowseSubprocessPath(s string) {
	m.browseSubprocessPath = s
}

func (m *macApp) IsCEF(s bool) {
	m.isCEF = s
}

func (m *macApp) cefHelper() {
	if m.isCEF {
		if m.cefFrameworksDir == "" {
			m.cefFrameworksDir = os.Getenv("ENERGY_HOME")
		}
		if !fileExists(m.cefFrameworksDir) {
			panic("cef frameworks 不存在: " + m.cefFrameworksDir)
		}
		if !fileExists(m.macAppFrameworksDir) {
			os.Mkdir(m.macAppFrameworksDir, 0755)
		}
		if m.browseSubprocessPath != "" {
			if !fileExists(m.browseSubprocessPath) {
				panic("子进程执行文件不存在: " + m.browseSubprocessPath)
			}
		}
		b := fileExists(m.macAppFrameworksDir + consts.Separator + cef)
		if !b {
			copyDir(m.cefFrameworksDir, m.macAppFrameworksDir)
		}
		for _, app := range helpers {
			var execName = m.execName
			hPath := fmt.Sprintf("%s/%s %s.app", m.macAppFrameworksDir, execName, app)
			if !fileExists(hPath) {
				if err := os.MkdirAll(hPath, 0755); err != nil {
					panic("创建cef helper失败: " + err.Error())
				}
			}
			helper := &macApp{}
			helper.browseSubprocessPath = m.browseSubprocessPath
			helper.execName = fmt.Sprintf("%s %s", m.execName, app)
			helper.macContentsDir = hPath + "/Contents"
			helper.macOSDir = helper.macContentsDir + "/MacOS"
			helper.macResources = helper.macContentsDir + "/Resources"
			helper.execFile = helper.macOSDir + "/" + helper.execName
			helper.plistFileName = helper.macContentsDir + "/Info.plist"
			helper.pkgInfoFileName = helper.macContentsDir + "/PkgInfo"
			helper.macAppFrameworksDir = helper.macContentsDir + "/Frameworks"
			helper.lsUIElement = "true"
			m.createMacOSApp(helper)
		}
	}
}

func (m *macApp) runMacOSApp() {
	cmd := exec.Command(m.execFile, os.Args...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func (m *macApp) copyDylib() {
	// 文件不存在，复制
	if !fileExists(m.lclLibFileName) {
		if fileExists(libname.LibName) {
			copyFile(libname.LibName, m.lclLibFileName)
		}
	} else {
		// 文件存在，对比后更新
		if fileExists(libname.LibName) {
			f1, _ := os.Stat(libname.LibName)
			f2, _ := os.Stat(m.lclLibFileName)
			if f1.Size() != f2.Size() {
				copyFile(libname.LibName, m.lclLibFileName)
			}
		}
	}
}

// 以一个Mac下的app形式运行
// 调试下使用，正式发布的时候虽然可 以不用去掉，但也不咋好
func (*macApp) createMacOSApp(m *macApp) bool {
	if !fileExists(m.macOSDir) {
		if err := os.MkdirAll(m.macOSDir, 0755); err != nil {
			return false
		}
	}
	if !fileExists(m.macResources) {
		os.MkdirAll(m.macResources, 0755)
	}
	resName := fmt.Sprintf("%s/%s.icns", m.macResources, m.execName)
	if !fileExists(resName) {
		ioutil.WriteFile(resName, macOSAppIcon, 0666)
	}
	if !fileExists(m.plistFileName) {
		datas := map[string]string{
			"execName":    m.execName,
			"devRegion":   "China", // China English
			"locale":      "zh_CN", //os.Getenv("LANG"),
			"copyright":   "copyright xxxx",
			"LSUIElement": m.lsUIElement,
		}
		buff := bytes.NewBuffer([]byte{})
		tmp := template.New("file")
		tmp.Parse(infoplist)
		tmp.Execute(buff, datas)
		ioutil.WriteFile(m.plistFileName, buff.Bytes(), 0666)
	}
	if !fileExists(m.pkgInfoFileName) {
		ioutil.WriteFile(m.pkgInfoFileName, pkgInfo, 0666)
	}
	if m.browseSubprocessPath != "" && !m.isMain {
		if _, err := copyFile(m.browseSubprocessPath, m.execFile); err == nil {
			os.Chmod(m.execFile, 0755)
			return true
		}
	} else {
		if _, err := copyFile(os.Args[0], m.execFile); err == nil {
			os.Chmod(m.execFile, 0755)
			return true
		}
	}
	return false
}
