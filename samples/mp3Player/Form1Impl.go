// 在这里写你的事件

package main

import (
	"fmt"
	"github.com/energye/golcl/lcl"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/energye/golcl/lcl/types"

	"github.com/energye/golcl/pkgs/bass"
)

//::private::
type TForm1Fields struct {
	bassPlayer *bass.TBass
	playCtl    *TPlayControl
	progress   *TImageTrackBar
	volbar     *TImageTrackBar
}

func (f *TForm1) OnFormCreate(sender lcl.IObject) {
	f.SetDoubleBuffered(true)
	f.EnabledMaximize(false)
	//f.SetColor(0x39302c)

	f.playCtl = NewPlayControl(f)
	f.playCtl.SetParent(f.Panel2)
	f.playCtl.SetAlign(types.AlClient)
	f.playCtl.OnSelect = f.OnPlayListSelect
	f.playCtl.SingerPic = lcl.AsBitmap(f.ImgSinger.Picture().Graphic())

	f.progress = NewImageTrackBar(f)
	f.progress.SetParent(f)
	f.progress.SetImages(f.ImgThumb, f.ImgBk, f.ImgFore)
	f.progress.SetLeft(58)
	f.progress.SetTop(288)
	f.progress.SetWidth(200)
	f.progress.SetPosition(0)

	f.volbar = NewImageTrackBar(f)
	f.volbar.SetParent(f)
	f.volbar.SetImages(f.ImgThumb, f.ImgBk, f.ImgFore)
	f.volbar.SetLeft(82)
	f.volbar.SetTop(318)
	f.volbar.SetWidth(80)
	f.volbar.SetPosition(60)
	f.volbar.OnTrackChange = f.OnVolChange

	f.bassPlayer = bass.NewBass()
	// 我的测试
	switch runtime.GOOS {
	case "windows":
		f.addFoler("F:\\KuGou")
	case "darwin":
		usrHome := os.Getenv("HOME")
		f.addFoler(usrHome + "/Music/网易云音乐")
		f.addFoler(usrHome + "/Music/iTunes/iTunes Media/Music")
	case "linux":
		f.addFoler("/home/ying32/音乐")
	}
}

func (f *TForm1) OnFormDestroy(sender lcl.IObject) {

	f.bassPlayer.Close()
	bass.BassFree()
}

func (f *TForm1) OnMIAddFileClick(sender lcl.IObject) {
	if f.OpenDialog1.Execute() {
		f.addFile(f.OpenDialog1.FileName())
	}
}

func (f *TForm1) OnMIAddFolderClick(sender lcl.IObject) {
	if ok, str := lcl.SelectDirectory2("选择目录", "", false); ok {
		f.addFoler(str)
	}
}

func (f *TForm1) addFile(fileName string) {
	name := filepath.Base(fileName)
	ext := strings.ToLower(filepath.Ext(name))
	if ext != ".mp3" && ext != ".m4a" {
		return
	}
	name = name[:len(name)-len(ext)]
	nameArr := strings.Split(name, "-")
	caption := ""
	singer := ""
	if len(nameArr) >= 2 {
		caption = strings.TrimSpace(nameArr[1])
		singer = strings.TrimSpace(nameArr[0])
	} else {
		caption = strings.TrimSpace(name)
		singer = caption
	}

	lenVal := int32(bass.GetFileLength(fileName))
	f.playCtl.Add(TPlayListItem{caption, singer, lenVal, "", fileName})

}

func (f *TForm1) addFoler(rootPath string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("addFoler: ", err)
		}
	}()
	_, err := os.Stat(rootPath)
	if os.IsNotExist(err) {
		return
	}
	filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		ext := filepath.Ext(info.Name())
		if ext == ".mp3" || ext == ".m4a" {
			f.addFile(path)
		}
		return nil
	})
}

func (f *TForm1) OnPlayListSelect(sender lcl.IObject, item TPlayListItem) {
	f.stopPlay()
	f.bassPlayer.Close()
	f.bassPlayer.OpenFile(item.FileName)
	f.bassPlayer.SetVolume(f.volbar.Position())
	f.play()
	str := item.Caption + " - " + item.Singer
	f.SetCaption(str)

}

func (f *TForm1) OnBtnPlayClick(sender lcl.IObject) {
	f.play()
}

func (f *TForm1) OnTimer1Timer(sender lcl.IObject) {
	if f.bassPlayer.IsValid() && f.bassPlayer.State == bass.PsPlaying {

		f.LblTime.SetCaption(f.bassPlayer.TimeStrLabel())
		pos, _ := f.bassPlayer.GetPosition()
		mLen, _ := f.bassPlayer.GetLength()

		if pos >= mLen {
			if f.playCtl.CanNext() {
				f.playCtl.Next()
			} else {
				f.stopPlay()
				f.playCtl.Stop()
			}
		}

		f.progress.SetPosition(int(float32(pos) / float32(mLen) * 100))

		caption := []rune(f.Caption())
		if len(caption) > 0 {
			temp := caption
			c := temp[:1]
			temp = temp[1:]
			f.SetCaption(string(temp) + string(c))
		}
	}
}

func (f *TForm1) OnVolChange(sender lcl.IObject) {
	f.bassPlayer.SetVolume(f.volbar.Position())
}

func (f *TForm1) OnBtnPauseClick(sender lcl.IObject) {
	f.pause()
}

func (f *TForm1) OnBtnMinClick(sender lcl.IObject) {
	lcl.Application.Minimize()
}

func (f *TForm1) OnBtnCloseClick(sender lcl.IObject) {
	lcl.Application.Terminate()
}

func (f *TForm1) stopPlay() {
	f.Timer1.SetEnabled(false)

	f.bassPlayer.Stop()

	f.bassPlayer.SetPosition(0)
	f.progress.SetPosition(0)

	f.BtnPause.Hide()
	f.BtnPlay.Show()

	f.SetCaption("Mp3Player")
}

func (f *TForm1) play() {
	f.bassPlayer.Play(false)

	f.BtnPlay.Hide()
	f.BtnPause.Show()
	f.Timer1.SetEnabled(true)

}

func (f *TForm1) pause() {
	f.bassPlayer.Pause()
	f.BtnPause.Hide()
	f.BtnPlay.Show()
	f.Timer1.SetEnabled(false)

}

func (f *TForm1) OnPanel1MouseDown(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	if button == types.MbLeft {
		//f.Perform(messages.WM_SYSCOMMAND, message)
	}
}

func (f *TForm1) OnBtnPrevClick(sender lcl.IObject) {
	f.playCtl.Prev()
}

func (f *TForm1) OnBtnNextClick(sender lcl.IObject) {
	f.playCtl.Next()
}
