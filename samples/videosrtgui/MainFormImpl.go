// Automatically generated by the simple GoVCL IDE.
// 不要更改此文件名
// 在这里写你的事件。

package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/types"
)

//::private::
type TMainFormFields struct {
	extMap  map[string]string
	running bool
}

func (f *TMainForm) OnFormCreate(sender lcl.IObject) {

	f.extMap = make(map[string]string, 0)
	ss := strings.Split(f.OpenDialog1.Filter(), "|")

	if len(ss) >= 2 {
		ss = strings.Split(strings.Replace(ss[1], "*", "", -1), ";")
		for _, ext := range ss {
			s := strings.ToLower(strings.TrimSpace(ext))
			f.extMap[s] = s
		}
	}
}

func (f *TMainForm) OnFormDropFiles(sender lcl.IObject, fileNames []string) {
	f.addFiles(fileNames)
}

// 帮助菜单的

func (f *TMainForm) OnMiQQGroupClick(sender lcl.IObject) {
	rtl.SysOpen("https://jq.qq.com/?_wv=1027&k=5Sv7Qiq")
}

func (f *TMainForm) OnMiSponsorClick(sender lcl.IObject) {
	rtl.SysOpen("https://z-kit.cc/sponsor.html")
}

func (f *TMainForm) OnMiHelpTextClick(sender lcl.IObject) {
	rtl.SysOpen("https://z-kit.cc/doc.html")
}

func (f *TMainForm) OnMiGiteeClick(sender lcl.IObject) {
	rtl.SysOpen("https://gitee.com/ying32/govcl")
}

func (f *TMainForm) OnMiGithubClick(sender lcl.IObject) {
	rtl.SysOpen("https://github.com/energye/golcl")
}

// --- 软件设置

func (f *TMainForm) OnMiOSSSaveSettingsClick(sender lcl.IObject) {
	if OSSSaveSettings.ShowModal() == types.MrOk {
		fmt.Println("OSSSaveSettings 点击了确认按钮")
		fmt.Println(OSSSaveSettings.EdtEndPoint.Text(),
			OSSSaveSettings.EdtAccessKeyId.Text(),
			OSSSaveSettings.EdtAccessKeySecret.Text(),
			OSSSaveSettings.EdtBucketName.Text(),
			OSSSaveSettings.EdtBucketDomain.Text())

	}
}

func (f *TMainForm) OnMiAppSettingsClick(sender lcl.IObject) {
	if AppSettings.ShowModal() == types.MrOk {

		fmt.Println("AppSettings 点击了确认按钮")
		fmt.Println(AppSettings.EdtTaskParallelCount.Value(),
			AppSettings.EdtOutputFileDir.Text(),
			AppSettings.ChkCloseOSS.Checked(),
			AppSettings.ChkCloseAutoUpdate.Checked())
	}
}

// 引擎

func (f *TMainForm) OnMiNewAliyunAudoEngineClick(sender lcl.IObject) {
	if NewAliyunAudioEngine.ShowModal() == types.MrOk {

		name := strings.TrimSpace(NewAliyunAudioEngine.EdtName.Text())
		if f.CbbAudioEngines.Items().IndexOf(name) == -1 {
			f.CbbAudioEngines.Items().Add(name)
		}
		f.checkAudioEngine()

		fmt.Println("NewAliyunAudioEngine 点击了确认按钮")
		fmt.Println(name,
			NewAliyunAudioEngine.EdtAppKey.Text(),
			NewAliyunAudioEngine.EdtAccessKeyId.Text(),
			NewAliyunAudioEngine.EdtAccessKeySecret.Text(),
			NewAliyunAudioEngine.CbbArea.ItemIndex(),
			NewAliyunAudioEngine.CbbArea.Text(),
		)
	}
}

func (f *TMainForm) checkAudioEngine() {
	if f.CbbAudioEngines.Items().Count() == 0 {
		return
	}
	if f.CbbAudioEngines.ItemIndex() == -1 {
		f.CbbAudioEngines.SetItemIndex(0)
	}
}

func (f *TMainForm) checkTranslateEngine() {
	if f.CbbTranslateEngines.Items().Count() == 0 {
		return
	}
	if f.CbbTranslateEngines.ItemIndex() == -1 {
		f.CbbTranslateEngines.SetItemIndex(0)
	}
}

func (f *TMainForm) addTranslateEngine(s string) {
	name := strings.TrimSpace(s)
	if f.CbbTranslateEngines.Items().IndexOf(name) == -1 {
		f.CbbTranslateEngines.Items().Add(name)
	}
}

func (f *TMainForm) OnMiNewBaiduTranslateEngineClick(sender lcl.IObject) {
	if NewBaiduTranslateEngine.ShowModal() == types.MrOk {
		f.addTranslateEngine(NewBaiduTranslateEngine.EdtName.Text())
		f.checkTranslateEngine()
	}
	fmt.Println("NewBaiduTranslateEngine 点击了确认按钮")
	fmt.Println(NewBaiduTranslateEngine.EdtName.Text(),
		NewBaiduTranslateEngine.EdtAppId.Text(),
		NewBaiduTranslateEngine.EditAppSecret.Text(),
		NewBaiduTranslateEngine.CbbType.ItemIndex(),
		NewBaiduTranslateEngine.CbbType.Text(),
	)
}

func (f *TMainForm) OnMiNewTencentTranslateEngineClick(sender lcl.IObject) {
	if NewTencentTranslateEngine.ShowModal() == types.MrOk {
		f.addTranslateEngine(NewTencentTranslateEngine.EdtName.Text())
		f.checkTranslateEngine()
	}
	fmt.Println("NewTencentTranslateEngine 点击了确认按钮")
	fmt.Println(NewTencentTranslateEngine.EdtName.Text(),
		NewTencentTranslateEngine.EdtSecretId.Text(),
		NewTencentTranslateEngine.EditSecretKey.Text())
}

func (f *TMainForm) addFiles(files []string) {
	// 按照他原来的逻辑吧
	for _, s := range files {
		if _, ok := f.extMap[strings.TrimSpace(filepath.Ext(s))]; !ok {
			lcl.MessageDlg(fmt.Sprintf("文件后缀不允许：%s\r\n", s), types.MtWarning, types.MbOK)
			return
		}
	}

	f.MmoInpuFiles.Clear()
	for _, s := range files {
		f.MmoInpuFiles.Lines().Add(s)
	}
}

// 打开

func (f *TMainForm) OnMiOpenMediaFileClick(sender lcl.IObject) {
	if f.OpenDialog1.Execute() {
		s := f.OpenDialog1.Files()
		if s.Count() > 0 {
			var ss []string
			for i := int32(0); i < s.Count(); i++ {
				ss = append(ss, s.S(i))
			}
			f.addFiles(ss)
		}
	}
}

func (f *TMainForm) OnBtnDelAudioEngineClick(sender lcl.IObject) {

}

func (f *TMainForm) OnPanel6Resize(sender lcl.IObject) {
	// 始终2个都一样大
	w := f.Panel6.Width() - f.Splitter1.Width()
	f.MmoInpuFiles.SetWidth(w / 2)
	f.MmoOutputLog.SetWidth(w / 2)
}

func (f *TMainForm) OnActGenerateSubtitlesUpdate(sender lcl.IObject) {
	lcl.AsAction(sender).SetEnabled(!f.running)
}

func (f *TMainForm) OnActGenerateSubtitlesExecute(sender lcl.IObject) {
	f.running = true
	f.MmoOutputLog.Clear()
	f.MmoOutputLog.Lines().Add("开始生成中......")
	f.ActGenerateSubtitles.SetCaption("正在生成中...")
	go func() {
		defer func() {
			f.running = false
		}()
		for i := 0; i < 6; i++ {
			lcl.ThreadSync(func() {
				f.MmoOutputLog.Lines().Add(fmt.Sprintf("%d", i))
			})
			time.Sleep(time.Second * 1)
		}

		lcl.ThreadSync(func() {
			f.MmoOutputLog.Lines().Add("完成......")
			f.ActGenerateSubtitles.SetCaption("生成字幕")
		})
	}()
}

func (f *TMainForm) OnActAudioEngineDeleteExecute(sender lcl.IObject) {
	f.CbbAudioEngines.Items().Delete(f.CbbAudioEngines.ItemIndex())
	f.checkAudioEngine()
}

func (f *TMainForm) OnActAudioEngineDeleteUpdate(sender lcl.IObject) {
	lcl.AsAction(sender).SetEnabled(f.CbbAudioEngines.Items().Count() > 1)
}

func (f *TMainForm) OnActTranslateEngineDeleteExecute(sender lcl.IObject) {
	f.CbbTranslateEngines.Items().Delete(f.CbbTranslateEngines.ItemIndex())
	f.checkTranslateEngine()
}

func (f *TMainForm) OnActTranslateEngineDeleteUpdate(sender lcl.IObject) {
	lcl.AsAction(sender).SetEnabled(f.CbbTranslateEngines.Items().Count() > 1)
}

func (f *TMainForm) OnFormCloseQuery(sender lcl.IObject, canClose *bool) {
	if f.running {
		*canClose = lcl.MessageDlg("当前正运行中，是否等待完成后再关闭？", types.MtWarning, types.MbYes, types.MbNo) == types.MrYes
	}
}

func (f *TMainForm) OnTBtnSettingsClick(sender lcl.IObject) {
	f.TBtnSettings.CheckMenuDropdown()
}

func (f *TMainForm) OnTBtnHelpClick(sender lcl.IObject) {
	f.TBtnHelp.CheckMenuDropdown()
}

func (f *TMainForm) OnTBtnOpenClick(sender lcl.IObject) {
	f.TBtnOpen.CheckMenuDropdown()
}

func (f *TMainForm) OnTBtnNewClick(sender lcl.IObject) {
	f.TBtnNew.CheckMenuDropdown()
}
