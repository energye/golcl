// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
	"strings"

	"github.com/energye/golcl/lcl"
)

//::private::
type TNewConnectionFields struct {
}

func (f *TNewConnection) OnFormCreate(sender lcl.IObject) {
	f.ScreenCenter()
}

func (f *TNewConnection) OnBtnTestConnectionClick(sender lcl.IObject) {
	name := strings.TrimSpace(f.EdtConnName.Text())
	host := strings.TrimSpace(f.EdtHost.Text())
	port := strings.TrimSpace(f.EdtPort.Text())
	password := strings.TrimSpace(f.EdtPassword.Text())

	if name == "" {
		lcl.ShowMessage("连接名得要设置个吧。")
		return
	}
	if _, ok := config[name]; ok {
		lcl.ShowMessage("配置名称已经存在，换个吧。")
		return
	}

	if host == "" {
		lcl.ShowMessage("要连接的主机名或者IP地址没设置啊。")
		return
	}
	if port == "" {
		lcl.ShowMessage("端口呢？")
		return
	}

	conn, err := newRedisConn(host, port, password)
	if err != nil {
		lcl.ShowMessage(err.Error())
		return
	}
	defer conn.Close()

	lcl.ShowMessage("测试连接成功。")
}
