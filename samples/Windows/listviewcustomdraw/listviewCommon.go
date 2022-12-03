package main

import (
	"strings"

	"fmt"

	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
)

var (
	fSortOrder bool
)

const (
	Color1 = 0x02E0F0F7
	Color2 = 0x02F0EEF7
)

// 这里是两个ListView实例共用代码，不然写得太多了

func fullListViewDataAndSetEvent(lv *lcl.TListView, trainData *TTrainSearchResultData) {
	if trainData != nil {
		// 填充数据到ListView

		//FormListViewDraw.ImageList1 这个imagelist内部没有项目的，用来撑大ListItem，设置了W=1, H=40
		//FormListViewDraw.LVTrain.SetGroupView(false)
		//FormListViewDraw.LVTrain.SetGridLines(false)
		lv.Items().BeginUpdate()
		for _, data := range trainData.Data.Datas {

			item := lv.Items().Add()

			item.SetCaption(data.StationTrainCode)
			//item.SetData(unsafe.Pointer(uintptr(i)))

			subItem := item.SubItems()
			subItem.Add(fmt.Sprintf("%s\r\n(%s)(%s)", data.FromStationName, data.StartTime, retStartorEnd(data.StartStationTelecode, data.FromStationTelecode, "始")))
			subItem.Add(fmt.Sprintf("%s\r\n(%s)(%s)", data.ToStationName, data.ArriveTime, retStartorEnd(data.EndStationTelecode, data.ToStationTelecode, "终")))
			subItem.Add(fmt.Sprintf("%s\r\n(%s)", getTimeStr(data.LiShiValue), dayDifference(data.DayDifference)))
			subItem.Add(data.SWZNum)
			subItem.Add(data.TZNnum)
			subItem.Add(data.ZYNum)
			subItem.Add(data.ZENum)
			subItem.Add(data.GRNum)
			subItem.Add(data.RWNum)
			subItem.Add(data.YWNum)
			subItem.Add(data.RZNum)
			subItem.Add(data.YZNum)
			subItem.Add(data.WZNum)
			subItem.Add(data.QTNum)
			subItem.Add(data.Note)
		}

		lv.Items().EndUpdate()
	}
	// 相关事件
	lv.SetOnAdvancedCustomDrawSubItem(lvTraiAdvancedCustomDrawSubItem)
	lv.SetOnAdvancedCustomDrawItem(lvTraiAdvancedCustomDrawItem)
	lv.SetOnColumnClick(lvTraiColumnClick)
	lv.SetOnCompare(lvTraiCompare)
	lv.SetOnClick(func(sender lcl.IObject) {
		sel := lcl.AsListView(sender).Selected()
		if sel.IsValid() {
			fmt.Println("select, index:", sel.Index(), ", caption:", sel.Caption(), ", data:", sel.Data())
		}
	})
}

// 子项目绘制函数
func lvTraiAdvancedCustomDrawSubItem(sender *lcl.TListView, item *lcl.TListItem, subItem int32,
	state types.TCustomDrawState, stage types.TCustomDrawStage, defaultDraw *bool) {
	canvas := sender.Canvas()
	font := canvas.Font()
	// 10 行后开始绘制，前面用于其它演示
	// 演示数据的使用，，，，， 如果使用了GroupView的话，因为分组排序问题会造成不对的，哈哈哈。。。
	i := item.Index() //   int(item.Data()) //item.Index()
	if i > 10 {
		if i%2 == 0 {
			canvas.Brush().SetColor(Color1)
		} else {
			canvas.Brush().SetColor(Color2)
		}
	}
	switch {
	case subItem >= 4 && subItem <= 14:
		s := item.SubItems().Strings(subItem - 1)
		if strings.Contains(s, "*") || strings.Contains(s, "--") || strings.Contains(s, "无") {
			font.SetColor(colors.ClSilver)
		} else {
			font.SetColor(colors.ClGreen)
		}
	default:
		font.SetColor(colors.ClBlack)
	}
}

// 演示隔行换色，如果使用lvTraiAdvancedCustomDrawSubItem来处理子项目，加上相关代码
func lvTraiAdvancedCustomDrawItem(sender *lcl.TListView, item *lcl.TListItem, state types.TCustomDrawState, Stage types.TCustomDrawStage, defaultDraw *bool) {
	//*defaultDraw = false

	// 演示数据的使用，，，，， 如果使用了GroupView的话，因为分组排序问题会造成不对的，哈哈哈。。。
	//i := int(item.Data()) //item.Index()
	// 10 行后开始绘制，前面用于其它演示
	if item.Index() > 10 {
		if item.Index()%2 == 0 {
			sender.Canvas().Brush().SetColor(Color1)
		} else {
			sender.Canvas().Brush().SetColor(Color2)
		}
		//sender.Canvas().FillRect(item.DisplayRect(types.DrBounds))
	}
}

// 柱头单击
func lvTraiColumnClick(sender lcl.IObject, column *lcl.TListColumn) {
	fSortOrder = !fSortOrder
	//lcl.AsListView(sender).AlphaSort()
	lcl.AsListView(sender).CustomSort(0, int(column.Index()))
}

// 排序
func lvTraiCompare(sender lcl.IObject, item1, item2 *lcl.TListItem, data int32, compare *int32) {
	var s1, s2 string
	if data != 0 {
		s1 = item1.SubItems().Strings(data - 1)
		s2 = item2.SubItems().Strings(data - 1)
	} else {
		s1 = item1.Caption()
		s2 = item2.Caption()
	}
	if fSortOrder {
		*compare = int32(strings.Compare(s1, s2))
	} else {
		*compare = -int32(strings.Compare(s1, s2))
	}
}
