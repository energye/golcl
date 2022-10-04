//go:build darwin
// +build darwin

package command

import "syscall"

func HideWindow(bool bool) *syscall.SysProcAttr {
	return nil
}
