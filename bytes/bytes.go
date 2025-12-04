package go_format_bytes

import "unsafe"

func ToStringUnsafe(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
