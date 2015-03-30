package iso8583

import (
	"fmt"
)

var _DEBUG = false

func SetDebug() {
	_DEBUG = true
}

func DumpHex(data []byte) {
	if !_DEBUG {
		return
	}
	Len := len(data)
	for i := 0; i < Len; i++ {
		if (i+1)%25 != 0 {
			fmt.Printf("%02x ", data[i])
		} else {
			fmt.Printf("%02x\n", data[i])
		}
	}
	fmt.Printf("\n")
}

func Debug(format string, a ...interface{}) {
	if !_DEBUG {
		return
	}
	fmt.Printf(format, a...)
	fmt.Printf("\n")
}
