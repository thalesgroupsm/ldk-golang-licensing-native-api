package license_api

// #cgo CFLAGS: -I .
// #cgo windows LDFLAGS: -L. -lapidsp_windows_x64
// #cgo linux LDFLAGS: -L. -lapidsp_linux_x86_64
//#include <hasp_api.h>
//#include <stdlib.h>
import "C"
import (
	"unsafe"
)

type LicenseApi struct {
}

func NewLicenseApi() *LicenseApi {
	return &LicenseApi{}
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

/*
LicenseApi HaspLogin
  - @param feature_id unique identifier for a specific Feature stored in a Sentinel protection key
  - @param vendor_code vendor code string

@return handle the resulting session handle
*/
func (l *LicenseApi) HaspLogin(feature_id int, vendor_code string) (handle uintptr, err int) {
	vendor_code_p := append([]byte(vendor_code), 0)
	err = int(C.hasp_login((C.uint)(feature_id), C.hasp_vendor_code_t(&vendor_code_p[0]), (*C.hasp_handle_t)(unsafe.Pointer(&handle))))
	return
}

/*
LicenseApi HaspLoginScope
  - @param feature_id unique identifier for a specific Feature stored in a Sentinel protection key
  - @param scope definition of the search parameters for this Feature ID
  - @param vendor_code vendor code string

@return handle the resulting session handle
*/
func (l *LicenseApi) HaspLoginScope(feature_id int, scope string, vendor_code string) (handle uintptr, err int) {
	vendor_code_p := append([]byte(vendor_code), 0)

	ptr_scope := C.CString(scope)
	defer C.free(unsafe.Pointer(ptr_scope))

	err = int(C.hasp_login_scope((C.uint)(feature_id), ptr_scope, C.hasp_vendor_code_t(&vendor_code_p[0]), (*C.hasp_handle_t)(unsafe.Pointer(&handle))))
	return
}

/*
LicenseApi HaspGetSessionInfo
  - @param handle Handle for the session
  - @param format Definition for the type of output data structure, in XML format
  - @param vendor_code vendor code string

@return handle the resulting session handle
*/
func (l *LicenseApi) HaspGetSessionInfo(handle uintptr, format string) (info string, err int) {

	info_p := make([]uintptr, 1)
	ptr_format := C.CString(format)
	defer C.free(unsafe.Pointer(ptr_format))

	err = int(C.hasp_get_sessioninfo((C.hasp_handle_t)(handle), ptr_format, (**C.char)(unsafe.Pointer(&info_p[0]))))
	if err == 0 {
		info = UintPtrToString(uintptr(info_p[0]))
		C.hasp_free((*C.char)(unsafe.Pointer(info_p[0])))
	}
	return
}

/*
LicenseApi HaspGetSize
  - @param handle Handle for the session
  - @param fileid Identifier for the file that is to be queried
  - @param size hold the resulting file size

@return
*/
func (l *LicenseApi) HaspGetSize(handle uintptr, file_id int, size *uint) (err int) {

	err = int(C.hasp_get_size((C.hasp_handle_t)(handle), C.uint(file_id), (*C.uint)(unsafe.Pointer(size))))
	return
}

/*
LicenseApi HaspGetRtc
  - @param handle Handle for the session
  - @param time hold the current time

@return
*/
func (l *LicenseApi) HaspGetRtc(handle uintptr, time *uint64) (err int) {

	err = int(C.hasp_get_rtc((C.hasp_handle_t)(handle), (*C.hasp_time_t)(unsafe.Pointer(time))))
	return
}

/*
LicenseApi HaspGetInfo
  - @param scope definition of the data that is to be searched, in XML format
  - @param format definition for the type of output data structure, in XML format
  - @param vendor_code vendor code string

@return info the information that is retrieved
*/
func (l *LicenseApi) HaspGetInfo(scope string, format string, vendor_code string) (info string, err int) {

	vendor_code_p := append([]byte(vendor_code), 0)
	info_p := make([]uintptr, 1)

	ptr_scope := C.CString(scope)
	defer C.free(unsafe.Pointer(ptr_scope))

	ptr_format := C.CString(format)
	defer C.free(unsafe.Pointer(ptr_format))

	err = int(C.hasp_get_info(ptr_scope, ptr_format, C.hasp_vendor_code_t(&vendor_code_p[0]), (**C.char)(unsafe.Pointer(&info_p[0]))))

	if err == 0 {
		info = UintPtrToString(info_p[0])
		C.hasp_free((*C.char)(unsafe.Pointer(info_p[0])))
	}

	return
}

/*
LicenseApi HaspLogout
  - @param handle handle for the session

@return
*/
func (l *LicenseApi) HaspLogout(handle uintptr) (err int) {

	err = int(C.hasp_logout((C.hasp_handle_t)(handle)))
	return
}

/*
LicenseApi HaspUpdate
  - @param update_data the complete update data

@return info the information that is retrieved
*/
func (l *LicenseApi) HaspUpdate(update_data string) (info string, err int) {

	ack := make([]uintptr, 1)

	ptr_data := C.CString(update_data)
	defer C.free(unsafe.Pointer(ptr_data))

	err = int(C.hasp_update(ptr_data, (**C.char)(unsafe.Pointer(&ack[0]))))
	if err == 0 {
		info = UintPtrToString(ack[0])
		C.hasp_free((*C.char)(unsafe.Pointer(ack[0])))
	}
	return
}

/*
LicenseApi HaspEncrypt
  - @param handle handle for the session
  - @param data the data to be encrypted
  - @param size the size of data to encrypted

@return
*/
func (l *LicenseApi) HaspEncrypt(handle uintptr, data *byte, size uint) (err int) {

	err = int(C.hasp_encrypt((C.hasp_handle_t)(handle), unsafe.Pointer(data), C.uint(size)))
	return
}

/*
LicenseApi HaspDecrypt
  - @param handle handle for the session
  - @param data the data to be decrypted
  - @param size the size of data to decrypted

@return
*/
func (l *LicenseApi) HaspDecrypt(handle uintptr, data *byte, size uint) (err int) {

	err = int(C.hasp_decrypt((C.hasp_handle_t)(handle), unsafe.Pointer(data), C.uint(size)))
	return
}

/*
LicenseApi HaspWrite
  - @param handle handle for the session
  - @param fileid identifier for the file that is to be written
  - @param offset offset in the file
  - @param len number of bytes to be written to the file
  - @param data data to be written to the file

@return
*/
func (l *LicenseApi) HaspWrite(handle uintptr, file_id uint, offset uint, len uint, data *byte) (err int) {
	err = int(C.hasp_write((C.hasp_handle_t)(handle), (C.uint)(file_id), (C.uint)(offset), (C.uint)(len), unsafe.Pointer(data)))
	return
}

/*
LicenseApi HaspRead
  - @param handle handle for the session
  - @param fileid identifier for the file that is to be read
  - @param offset offset in the file
  - @param len number of bytes to be read from the file
  - @param data the read data

@return
*/
func (l *LicenseApi) HaspRead(handle uintptr, file_id uint, offset uint, len uint, data *byte) (err int) {
	err = int(C.hasp_read((C.hasp_handle_t)(handle), (C.uint)(file_id), (C.uint)(offset), (C.uint)(len), unsafe.Pointer(data)))
	return err

}

/*
LicenseApi HaspTransfer
  - @param action parameters for the operation, in XML format
  - @param scope search parameters for the conatiner-id that is to be re-hosted
  - @param vendor_code vendor code string
  - @param recipient definition in XML format of the recipient computer, on which the detached Product will be installed

@return output the output information that is retrieved, in XML format
*/
func (l *LicenseApi) HaspTransfer(action string, scope string, vendor_code string, recipient string) (output string, err int) {

	output_p := make([]uintptr, 1)

	vendor_code_p := append([]byte(vendor_code), 0)

	ptr_action := C.CString(action)
	defer C.free(unsafe.Pointer(ptr_action))

	ptr_scope := C.CString(scope)
	defer C.free(unsafe.Pointer(ptr_scope))

	ptr_recipient := C.CString(recipient)
	defer C.free(unsafe.Pointer(ptr_recipient))

	err = int(C.hasp_transfer(ptr_action, ptr_scope, C.hasp_vendor_code_t(&vendor_code_p[0]), ptr_recipient, (**C.char)(unsafe.Pointer(&output_p[0]))))
	if err == 0 {
		output = UintPtrToString(output_p[0])
		C.hasp_free((*C.char)(unsafe.Pointer(output_p[0])))
	}

	return

}

/*
LicenseApi HaspConfig
  - @param config parameters for the operation, in XML format

@return
*/
/*func (l *LicenseApi) HaspConfig(config string) (err int) {

	configPtr := C.CString(config)
	defer C.free(unsafe.Pointer(configPtr))

	err = int(C.hasp_config(configPtr))
	return err

}*/
