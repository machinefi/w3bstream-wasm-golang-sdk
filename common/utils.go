package common

import "unsafe"

func StringToPointer(v string) (addr, size uint32) {
	return BytesToPointer([]byte(v))
}

func BytesToPointer(v []byte) (addr, size uint32) {
	ptr := &v[0]
	pptr := uintptr(unsafe.Pointer(ptr))
	return uint32(pptr), uint32(len(v))
}
