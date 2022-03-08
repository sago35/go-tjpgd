package tjpgd

/*
#include "tjpgd.h"
#include "tjpgd-go.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func Version() {
	fmt.Printf("TJpgD R0.0.3\n")
}

var callback func(left, top, right, bottom uint16, buf []uint16)

func SetCallback(fn func(left, top, right, bottom uint16, buf []uint16)) {
	callback = fn
}

type Scale int

const (
	ScaleNone = iota
	ScaleHalf
	ScaleQuarter
)

func DecodeFromBytes(b []byte, scale Scale) error {
	C.decodeFromBytes((*C.uchar)(unsafe.Pointer(&b[0])), C.int(len(b)), C.int(scale))
	return nil
}
