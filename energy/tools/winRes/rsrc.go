package winRes

import (
	"errors"
	"fmt"
	"github.com/energye/golcl/tools/command"
	"runtime"
)

// RSRC => github.com/akavel/rsrc
type RSRC struct {
	CMDDir       string //命令行执行并生成文件的根目录
	AppName      string
	IconPath     string
	ManifestPath string
	RSRCPath     string // rsrc.exe 目录
	arch         string
}

//runtime.GOARCH

//默认采用 syso 生成
//syso 方式生成失败，可尝试rsrc方式生成
//rsrc需要生成执行文件 > github.com/akavel/rsrc
func NewRSRC() *RSRC {
	m := &RSRC{
		arch:         runtime.GOARCH,
		AppName:      "app",
		IconPath:     "resources/app.ico",
		ManifestPath: "resources/manifest.xml",
	}
	return m
}

func (m *RSRC) Gen() error {
	var err error
	if !isExist(m.IconPath) {
		return errors.New(m.IconPath + " 文件未找到")
	}

	if !isExist(m.ManifestPath) {
		return errors.New(m.ManifestPath + " 文件未找到")
	}
	if m.RSRCPath != "" {
		if !isExist(m.RSRCPath) {
			return errors.New(m.RSRCPath + " 文件未找到")
		}
	} else {
		if runtime.GOOS == "windows" {
			m.RSRCPath = "rsrc.exe"
		} else if runtime.GOOS == "linux" {
			m.RSRCPath = "rsrc"
		}
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
		var out = m.AppName + "_" + runtime.GOOS + "_" + runtime.GOARCH + ".syso"
		// rsrc.exe -manifest app.manifest -ico resources/app.ico -o out.syso -arch amd64
		var args = []string{"-manifest", m.ManifestPath, "-ico", m.IconPath, "-o", out, "-arch", runtime.GOARCH}
		cmd.Command(m.RSRCPath, args...)
		cmd.Close()
	}
	return err
}
