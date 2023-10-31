package sampgo

import "C"
import (
	"fmt"
	"unsafe"
)

func findNative(name string) (C.AMX_NATIVE, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	native := C.sampgdk_FindNative(cName)
	if native == nil {
		return nil, fmt.Errorf("native could not be found")
	}

	return native, nil
}

func invokeNative(native C.AMX_NATIVE, types string, args ...interface{}) int {
	return C.sampgdk_InvokeNative(native, types, args)
}
