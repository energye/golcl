// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"

	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

type TConn struct {
	Sheet *lcl.TTabSheet
	Conn  redis.Conn
	Node  *lcl.TTreeNode
}

//::private::
type TMainFormFields struct {
	conns map[string]TConn
}

func (f *TMainForm) OnFormCreate(sender lcl.IObject) {
	f.ScreenCenter()
	f.conns = make(map[string]TConn, 0)

	f.TvConns.Items().Clear()
	loadConfig()
	f.updateTree()
}

func (f *TMainForm) OnFormDestroy(sender lcl.IObject) {
	f.closeAllConn()
}

func (f *TMainForm) updateTree() {
	for key, _ := range config {
		f.TvConns.Items().Add(nil, key)
	}
}

func (f *TMainForm) closeAllConn() {
	for key, _ := range f.conns {
		f.closeConn(key, false)
	}
}

func (f *TMainForm) closeConn(name string, free bool) {
	if v, ok := f.conns[name]; ok {
		v.Conn.Close()
		for i := v.Node.Count() - 1; i >= 0; i-- {
			v.Node.Item(i).Delete()
		}
		v.Node.SetExpanded(false)
		if free {
			v.Sheet.Free() // 删除并释放
		}
		delete(f.conns, name)
	}
}

func (f *TMainForm) OnBtnNewConnClick(sender lcl.IObject) {
	if NewConnection.ShowModal() == types.MrOk {
		item := TConnConfig{}
		item.Name = strings.TrimSpace(NewConnection.EdtConnName.Text())
		item.Host = strings.TrimSpace(NewConnection.EdtHost.Text())
		port, err := strconv.Atoi(NewConnection.EdtPort.Text())
		if err != nil {
			port = 6379
		}
		item.Port = port
		item.Password = strings.TrimSpace(NewConnection.EdtPassword.Text())
		config[item.Name] = item
		// 添加到列表
		f.TvConns.Items().Add(nil, item.Name)
		// 保存到文件
		saveConfig()
	}
}

func (f *TMainForm) OnTvConnsDblClick(sender lcl.IObject) {
	node := f.TvConns.Selected()
	if node != nil {
		text := node.Text()
		if val, ok := f.conns[text]; ok {
			// 已经连接，退出
			f.PgMain.SetActivePageIndex(val.Sheet.PageIndex())
			return
		}
		if val, ok := config[text]; ok {
			conn, err := newRedisConn(val.Host, fmt.Sprintf("%d", val.Port), val.Password)
			if err != nil {
				lcl.ShowMessage(err.Error())
				return
			}

			sheet := lcl.NewTabSheet(f)
			sheet.SetParent(f.PgMain)
			sheet.SetCaption(text)

			f.conns[text] = TConn{sheet, conn, node}

			// 异步列出数据库
			go f.syncGetDataBases(conn, node)

		}
	}
}

func (f *TMainForm) syncGetDataBases(conn redis.Conn, node *lcl.TTreeNode) {
	//CONFIG GET databases
	rep, err := redis.StringMap(conn.Do("CONFIG", "GET", "databases"))
	if err == nil {
		fmt.Println(rep["databases"])
		if val, ok := rep["databases"]; ok {
			if count, err := strconv.Atoi(val); err == nil {
				lcl.ThreadSync(func() {
					for i := 0; i < count; i++ {
						f.TvConns.Items().AddChild(node, fmt.Sprintf("db%d", i))
					}
					node.SetExpanded(true)
				})
			}
		}
	} else {
		fmt.Println(err)
	}
}

func (f *TMainForm) OnMICloseClick(sender lcl.IObject) {
	f.closeConn(f.PgMain.Pages(f.PgMain.ActivePageIndex()).Caption(), true)
}
