//----------------------------------------
//
// Copyright © sxm. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package main

import (
	"fmt"
	"strings"

	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

func showNotAllowEmpty(s *lcl.TEdit, tips string) bool {
	if strings.TrimSpace(s.Text()) == "" {
		lcl.MessageDlg(fmt.Sprintf("“%s”不能为空！", tips), types.MtInformation, types.MbOK)
		return true
	}
	return false
}
