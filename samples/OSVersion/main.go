package main

import (
	"fmt"

	. "github.com/energye/golcl/lcl/rtl/version"
	_ "github.com/energye/golcl/pkgs/winappres"
)

func main() {

	fmt.Println("Major: ", OSVersion.Major)
	fmt.Println("Minor: ", OSVersion.Minor)
	fmt.Println("Name: ", OSVersion.Name)
	fmt.Println("String: ", OSVersion.ToString())
	switch OSVersion.Platform {
	case PfWindows:
		if OSVersion.CheckMajor(10) {
			fmt.Println("当前为Windows10")
		} else {
			fmt.Println("当前不是Windows10")
		}
	case PfLinux:
		fmt.Println("LIBC Version:", OSVersion.LibCVersionMajor, ", ", OSVersion.LibCVersionMinor)
	case PfMacOS:
		//
	}

}
