//go:build !baremetal
//+build !baremetal

package tjpgd

/*
#include "tjpgd.h"
#include "tjpgd-go.h"
*/
import "C"

import (
	"unsafe"
)

var tpjbdBuf [256]uint16

//export callbackFromTjpgd
func callbackFromTjpgd(left, top, right, bottom uint16, buf *uint8, length uint16) {
	b := tpjbdBuf[:length/2]
	C.memcpy(unsafe.Pointer(&b[0]), unsafe.Pointer(buf), C.ulonglong(length))
	callback(left, top, right, bottom, b)
}
