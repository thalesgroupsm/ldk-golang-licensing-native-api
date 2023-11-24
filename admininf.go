package license_api

// #cgo CFLAGS: -I .
// #cgo windows LDFLAGS: -L. -lapidsp_windows_x64
// #cgo linux LDFLAGS: -L. -lapidsp_linux_x86_64
//#include <sntl_adminapi.h>
//#include <stdlib.h>
import "C"
import (
	"unsafe"
)

type AdminApi struct {
}

func NewAdminApi() *AdminApi {
	return &AdminApi{}
}

/*
Admin SntlAdminContextNewScope
  - @param scope XML scope specification

@return handle the resulting context handle
*/
func (a *AdminApi) SntlAdminContextNewScope(scope string) (context uintptr, err int) {
	admin_context := make([]uintptr, 1)

	ptr_scope := C.CString(scope)
	defer C.free(unsafe.Pointer(ptr_scope))

	err = int(C.sntl_admin_context_new_scope((**C.struct_sntl_admin_context_t)(unsafe.Pointer(&admin_context[0])),
		ptr_scope))
	if err == 0 {
		context = admin_context[0]
	}
	return
}

/*
Admin SntlAdminSet
  - @param context context handle of License Manager
  - @param input definition of the actions/settings to be done, in XML format

@return
*/
func (a *AdminApi) SntlAdminSet(context uintptr, input string) (status string, err int) {
	result_p := make([]uintptr, 1)

	ptr_action := C.CString(input)
	defer C.free(unsafe.Pointer(ptr_action))

	err = int(C.sntl_admin_set((*C.struct_sntl_admin_context_t)(unsafe.Pointer(context)),
		ptr_action,
		(**C.char)(unsafe.Pointer(&result_p[0]))))
	if err == 0 {
		status = UintPtrToString(result_p[0])
		C.sntl_admin_free((*C.char)(unsafe.Pointer(result_p[0])))
	}
	return
}

/*
Admin SntlAdminGet
  - @param context context handle of License Manager
  - @param scope definition of the data that is to be searched, in XML format
  - @param format definition of the data to be retrieved, in XML format

@return info the information string that is retrieved, in XML format
*/
func (a *AdminApi) SntlAdminGet(context uintptr, scope string, format string) (info string, err int) {
	result_p := make([]uintptr, 1)

	ptr_scope := C.CString(scope)
	defer C.free(unsafe.Pointer(ptr_scope))
	ptr_format := C.CString(format)
	defer C.free(unsafe.Pointer(ptr_format))

	err = int(C.sntl_admin_get((*C.struct_sntl_admin_context_t)(unsafe.Pointer(context)),
		ptr_scope,
		ptr_format,
		(**C.char)(unsafe.Pointer(&result_p[0]))))
	if err == 0 {
		info = UintPtrToString(result_p[0])
		C.sntl_admin_free((*C.char)(unsafe.Pointer(result_p[0])))
	}
	return
}

/*
Admin SntlAdminContextDelete
  - @param context context handle of License Manager

@return
*/
func (a *AdminApi) SntlAdminContextDelete(context uintptr) (err int) {
	err = int(C.sntl_admin_context_delete((*C.struct_sntl_admin_context_t)(unsafe.Pointer(context))))
	return
}
