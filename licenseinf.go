package license_api

// #cgo CFLAGS: -I .
// #cgo LDFLAGS: -L . -lhasp_windows_x64_demo
//#include <hasp_api_int.h>
//#include <stdlib.h>
import "C"
import (
	"fmt"
	"os"
	"unsafe"
)

const (
	HASP_FILEID_RW = 0xfff4
)

type LicenseApi struct {
	VendorId   string
	VendorCode string
}

func NewLicenseApi(vendorId string) *LicenseApi {
	var vendorCode []byte
	var err error
	if vendorId != "" {
		vendorCode, err = os.ReadFile(vendorId + ".hvc")
		if err != nil {
			return nil
		}
	}
	return &LicenseApi{
		VendorId:   vendorId,
		VendorCode: string(vendorCode),
	}
}

func strStr(haystack string, needle string) int {
	m := len(haystack)
	n := len(needle)

	if m == 0 || n == 0 {
		return -1
	}
	if m < n {
		return -1
	}
	for i := 0; i < m-n+1; i++ {
		k := i + n

		if haystack[i] != needle[0] {
			continue
		}

		if haystack[i:k] == needle {
			return i
		}

	}
	return -1
}

func UintPtrToString(r uintptr) string {
	p := (*uint8)(unsafe.Pointer(r))
	if p == nil {
		return ""
	}

	n, end, add := 0, unsafe.Pointer(p), unsafe.Sizeof(*p)
	for *(*uint8)(end) != 0 {
		end = unsafe.Add(end, add)
		n++
	}
	return string(unsafe.Slice(p, n))
}

func (l *LicenseApi) HaspLogin(feature_id int, vendor_code string) (handle uintptr) {

	vendor_code_p := append([]byte(vendor_code), 0)

	ret := C.hasp_login((C.uint)(feature_id), C.hasp_vendor_code_t(&vendor_code_p[0]), (*C.hasp_handle_t)(unsafe.Pointer(&handle)))

	if ret != 0 {
		handle = 0
	}

	return handle
}

/*
hasp_status_t HASP_CALLCONV hasp_login_scope(

						 hasp_feature_t feature_id,
	                const char *scope,
	                hasp_vendor_code_t vendor_code,
	                hasp_handle_t *handle);
*/
func (l *LicenseApi) HaspLoginScope(feature_id int, scope string, vendor_code string) (handle uintptr, r int) {

	vendor_code_p := append([]byte(vendor_code), 0)

	ptr_scope := C.CString(scope)
	defer C.free(unsafe.Pointer(ptr_scope))

	ret := C.hasp_login_scope((C.uint)(feature_id), ptr_scope, C.hasp_vendor_code_t(&vendor_code_p[0]), (*C.hasp_handle_t)(unsafe.Pointer(&handle)))

	r = int(ret)

	if ret != 0 {
		handle = 0
	}

	return

}

func (l *LicenseApi) HaspGetSessionInfo(handle uintptr, format string) (info string) {

	info_p := make([]uintptr, 1)

	ptr_format := C.CString(format)
	defer C.free(unsafe.Pointer(ptr_format))

	ret := C.hasp_get_sessioninfo((C.hasp_handle_t)(handle), ptr_format, (**C.char)(unsafe.Pointer(&info_p[0])))

	if ret != 0 {
		fmt.Printf("Call hasp_get_session_info:%d", ret)
	} else {
		info_string := UintPtrToString(uintptr(info_p[0]))
		info = info_string
	}
	return
}

func (l *LicenseApi) HaspGetInfo(scope string, format string, vendor_code string) (r int, info string) {

	vendor_code_p := append([]byte(vendor_code), 0)
	info_p := make([]uintptr, 1)

	ptr_scope := C.CString(scope)
	defer C.free(unsafe.Pointer(ptr_scope))

	ptr_format := C.CString(format)
	defer C.free(unsafe.Pointer(ptr_format))

	ret := C.hasp_get_info(ptr_scope, ptr_format, C.hasp_vendor_code_t(&vendor_code_p[0]), (**C.char)(unsafe.Pointer(&info_p[0])))

	if ret != 0 {
		fmt.Printf("Call hasp_get_info:%d", ret)
	} else {
		license_string := UintPtrToString(info_p[0])
		//hasp_free(info_p[0])
		info = license_string
	}
	r = (int)(ret)
	return
}

func (l *LicenseApi) HaspLogout(handle uintptr) (r int) {

	ret := C.hasp_logout((C.hasp_handle_t)(handle))

	r = (int)(ret)

	return
}

func (l *LicenseApi) HaspFree(address uintptr) {
	C.hasp_free((*C.char)(unsafe.Pointer(address)))
}

func (l *LicenseApi) HaspUpdate(v2c string) (r int) {

	var ack *C.char

	ptr_v2c := C.CString(v2c)
	defer C.free(unsafe.Pointer(ptr_v2c))

	ret := C.hasp_update(ptr_v2c, &ack)
	if ret != 0 {
		fmt.Printf("Call hasp_update:%d", ret)
	}
	C.hasp_free(ack)

	r = (int)(ret)

	return

}

func (l *LicenseApi) HaspEncrypt(handle uintptr, data *byte, len uint) (r int) {

	ret := C.hasp_encrypt((C.hasp_handle_t)(handle), unsafe.Pointer(data), C.uint(len))
	return int(ret)

}

func (l *LicenseApi) HaspDecrypt(handle uintptr, data *byte, len uint) (r int) {

	ret := C.hasp_decrypt((C.hasp_handle_t)(handle), unsafe.Pointer(data), C.uint(len))
	return int(ret)
}

func (l *LicenseApi) HaspWrite(handle uintptr, fileid uint, offset uint, len uint, data *byte) (r int) {
	ret := C.hasp_write((C.hasp_handle_t)(handle), (C.uint)(fileid), (C.uint)(offset), (C.uint)(len), unsafe.Pointer(data))
	return int(ret)
}

func (l *LicenseApi) HaspRead(handle uintptr, fileid uint, offset uint, len uint, data *byte) (r int) {
	ret := C.hasp_read((C.hasp_handle_t)(handle), (C.uint)(fileid), (C.uint)(offset), (C.uint)(len), unsafe.Pointer(data))
	return int(ret)

}

func (l *LicenseApi) HaspCleanup() {

	C.hasp_cleanup()
}

func (l *LicenseApi) HaspTransfer(action string, scope string, vendor_code string, recipient string) (r int, output string) {

	output_p := make([]uintptr, 1)

	vendor_code_p := append([]byte(vendor_code), 0)

	ptr_action := C.CString(action)
	defer C.free(unsafe.Pointer(ptr_action))

	ptr_scope := C.CString(scope)
	defer C.free(unsafe.Pointer(ptr_scope))

	ptr_recipient := C.CString(recipient)
	defer C.free(unsafe.Pointer(ptr_recipient))

	ret := C.hasp_transfer(ptr_action, ptr_scope, C.hasp_vendor_code_t(&vendor_code_p[0]), ptr_recipient, (**C.char)(unsafe.Pointer(&output_p[0])))
	if ret != 0 {
		fmt.Printf("Call hasp_transfer:%d", ret)
	} else {
		output = UintPtrToString(output_p[0])
		l.HaspFree(output_p[0])
	}

	r = (int)(ret)

	return

}

func (l *LicenseApi) HaspConfig(config string) (r int) {

	configPtr := C.CString(config)
	defer C.free(unsafe.Pointer(configPtr))

	ret := C.hasp_config(configPtr)
	if ret != 0 {
		fmt.Printf("Call hasp_config:%d", ret)
	}

	r = (int)(ret)

	return r

}

/*
func sntl_admin_context_new(host string, port int, password string) (context uintptr) {

	admin_context := make([]uintptr, 1)

	//C.sntl_admin_context_new_scope()

	ptr_host := C.CString(host)
	defer C.free(unsafe.Pointer(ptr_host))

	ptr_password := C.CString(password)
	defer C.free(unsafe.Pointer(ptr_password))

	ret := C.sntl_admin_context_new((**C.struct_sntl_admin_context_t)(unsafe.Pointer(&admin_context[0])),
		ptr_host,
		(C.ushort)(port),
		ptr_password)

	if ret != 0 {
		fmt.Printf("Call sntl_admin_context_new:%d", ret)
	}
	context = admin_context[0]
	return
}

func sntl_admin_context_new_scope() (context uintptr) {

	admin_context := make([]uintptr, 1)

	var scope string = "<haspscope>" +
		"<host>sntl_integrated_lm</host>" +
		"<vendor_code>" + string(VendorCode) + "</vendor_code>" +
		"</haspscope>"

	//C.sntl_admin_context_new_scope()
	ptr_scope := C.CString(scope)
	defer C.free(unsafe.Pointer(ptr_scope))

	ret := C.sntl_admin_context_new_scope((**C.struct_sntl_admin_context_t)(unsafe.Pointer(&admin_context[0])),
		ptr_scope)

	if ret != 0 {
		fmt.Printf("Call sntl_admin_context_new:%d", ret)
	}
	context = admin_context[0]
	return
}

func sntl_admin_set(context uintptr, action string) (r int) {

	result_p := make([]uintptr, 1)

	ptr_action := C.CString(action)
	defer C.free(unsafe.Pointer(ptr_action))

	ret := C.sntl_admin_set((*C.struct_sntl_admin_context_t)(unsafe.Pointer(context)),
		ptr_action,
		(**C.char)(unsafe.Pointer(&result_p[0])))

	if ret != 0 {
		fmt.Printf("Call sntl_admin_set:%d, %s", ret, action)
	}

	return int(ret)

}

func sntl_admin_get(context uintptr, scope string, format string) (r int, info string) {

	result_p := make([]uintptr, 1)

	ptr_scope := C.CString(scope)
	defer C.free(unsafe.Pointer(ptr_scope))

	ptr_format := C.CString(format)
	defer C.free(unsafe.Pointer(ptr_format))

	ret := C.sntl_admin_get((*C.struct_sntl_admin_context_t)(unsafe.Pointer(context)),
		ptr_scope,
		ptr_format,
		(**C.char)(unsafe.Pointer(&result_p[0])))

	if ret != 0 {
		fmt.Printf("Call sntl_admin_set:%d", ret)
	}
	r = int(ret)

	info = UintPtrToString(result_p[0])

	return
}

func sntl_admin_context_delete(context uintptr) {

	ret := C.sntl_admin_context_delete((*C.struct_sntl_admin_context_t)(unsafe.Pointer(context)))

	if ret != 0 {
		fmt.Printf("Call sntl_admin_context_delete:%d", ret)
	}

}
*/
