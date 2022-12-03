// 在这里写你的事件

package main

import (
	"fmt"
	"math"
	"os"
	"runtime"

	"github.com/energye/golcl/lcl/rtl"

	"github.com/energye/golcl/pkgs/libvlc"

	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types/colors"
)

//::private::
type TMainFormFields struct {
	player *libvlc.TVLCMediaPlayer
}

func (f *TMainForm) OnFormCreate(sender lcl.IObject) {
	f.ScreenCenter()
	os.Setenv("VLC_PLUGIN_PATH", rtl.ExtractFilePath(lcl.Application.ExeName())+"/plugins/")
	f.player = libvlc.NewVLCMediaPlayer()
	if f.player == nil {
		lcl.ShowMessage(fmt.Sprint("创建MediaPlayer失败:", libvlc.ErrMsg()))
		return
	}
	if runtime.GOOS != "linux" {
		f.player.SethWnd(f.PnlVideo.Handle())
	}
	//f.player.LoadFromFile("I:\\30分钟教你学会超火的日语歌曲- PLANET.mp4")
	f.player.LoadFromFile("test.mp4")
	f.PnlVideo.SetColor(colors.ClBlack)

}

func (f *TMainForm) OnFormDestroy(sender lcl.IObject) {
	if f.player != nil {
		f.player.Free()
	}
}

func (f *TMainForm) OnActPlayExecute(sender lcl.IObject) {
	if runtime.GOOS == "linux" {
		// 不太好使。。。
		//f.player.SethWnd(types.HWND(f.PlatformWindow().XID()))
	}
	f.Timer1.SetEnabled(true)
	f.player.Play()
}

func (f *TMainForm) OnActPlayUpdate(sender lcl.IObject) {
	lcl.AsAction(sender).SetEnabled(f.player != nil && !f.player.Playing())
}

func (f *TMainForm) OnActStopExecute(sender lcl.IObject) {
	f.Timer1.SetEnabled(false)
	f.player.Stop()
	f.LblCurTime.SetCaption("00:00:00")
	f.LblTotalTime.SetCaption("00:00:00")
	f.TrackBar1.SetPosition(0)
}

func (f *TMainForm) OnActStopUpdate(sender lcl.IObject) {
	lcl.AsAction(sender).SetEnabled(f.player != nil && f.player.Playing())
}

func (f *TMainForm) OnActPauseExecute(sender lcl.IObject) {
	f.Timer1.SetEnabled(false)
	f.player.Pause()
}

func (f *TMainForm) OnActPauseUpdate(sender lcl.IObject) {
	lcl.AsAction(sender).SetEnabled(f.player != nil && f.player.Playing())
}

func (f *TMainForm) OnTimer1Timer(sender lcl.IObject) {
	if f.player != nil && f.player.MediaPlayerValid() {
		f.LblCurTime.SetCaption(f.player.MediaTimeString())
		f.LblTotalTime.SetCaption(f.player.MediaLengthString())
		f.TrackBar1.SetPosition(int32(math.Ceil(float64(f.player.Position() * 100.0))))
		if f.player.MediaTime() >= f.player.MediaLength() {
			//	f.ActStop.Execute()
		}
	}
}
