package main

import (
	"testing"

	"fmt"

	"github.com/energye/golcl/lcl"
)

func TestPSDReader(t *testing.T) {
	bmp := lcl.NewBitmap()
	defer bmp.Free()
	if err := PsdToBitmap("J:\\Delphi\\PsSources\\spotify.psd", bmp); err == nil {
		bmp.SaveToFile("test.bmp")
	} else {
		fmt.Println("err:", err)
	}
}
