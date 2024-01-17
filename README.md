## License native API Endpoints

All URIs are relative to *https://localhost:8088/sentinel/ldk_runtime/v1*

Class | Method | Description
------------ | ------------- |  -------------
*LicensingApi* | **HaspLogin(feature_id int, vendor_code string) (handle uintptr, err int)** | login
*LicensingApi* | **HaspLoginScope(feature_id int, scope string, vendor_code string) (handle uintptr, err int)** | login scope
*LicensingApi* | [**HaspLogout(handle uintptr) (err int)**] |  logout
*LicensingApi* | [**HaspEncrypt(handle uintptr, data *byte, size uint) (err int)**] |  encrypt
*LicensingApi* | [**HaspDecrypt(handle uintptr, data *byte, size uint) (err int)**] |  desrypt
*LicensingApi* | [**HaspGetSize(handle uintptr, file_id int, size *uint) (err int)**] |  get size
*LicensingApi* | [**HaspWrite(handle uintptr, file_id uint, offset uint, len uint, data *byte) (err int)**] |  write file
*LicensingApi* | [**HaspRead(handle uintptr, file_id uint, offset uint, len uint, data *byte)**] |  read file
*LicensingApi* | [**HaspGetRtc(handle uintptr, time *uint64) (err int)**] |  get rtc
*LicensingApi* | **HaspGetSessionInfo(handle uintptr, format string) (info string, err int)** |  get session info
*LicensingApi* | **HaspGetInfo(scope string, format string, vendor_code string) (info string, err int)** |  get info
*LicensingApi* | [**HaspUpdate(update_data string) (info string, err int)**] |  update
*LicensingApi* | [**HaspTransfer(action string, scope string, vendor_code string, recipient string) (output string, err int)**] |  tranfter
*AdminApi* | [**SntlAdminContextNewScope(scope string) (context uintptr, err int)**] |  create admin context 
*AdminApi* | [**SntlAdminSet(context uintptr, input string) (status string, err int)**] |  admin set
*AdminApi* | [**SntlAdminGet(context uintptr, scope string, format string) (info string, err int)**] |  admin get
*AdminApi* | [**SntlAdminContextDelete(context uintptr) (err int)**] |  delete context

## Dependencies
On each platform, you'll need a Go installation that supports cgo compilation. On Windows, you also need to download and install GCC compiler as the guide on https://sourceforge.net/projects/gcc-win64/.  

You need install **Sentinel LDK & LDK-EMS** on your system. If you are a new user, please go https://supportportal.thalesgroup.com to download and install the package.

On Winodws, please copy the files **apidsp_windows_x64.dll** and **hasp_windows_x64_xxxx.dll** (xxxx means vendor id) to a folder which DLL can be searched（https://learn.microsoft.com/en-us/windows/win32/dlls/dynamic-link-library-search-order）by your applications. 

On Linux, please copy **libapidsp_linux_x86_64.so** and **libhasp_linux_x86_64_xxxx.so** (xxxx means vendor id) to a specific folder, then update the go env **CGO_LDFLAGS** and set enviroment **LD_LIBRARY_PATH**
```shell
go env -w CGO_LDFLAGS="-g -O2 -L<your folder>"
export LD_LIBRARY_PATH=<your folder>
```
##usage
```go
import licensingApi "github.com/thalesgroupsm/ldk-golang-licensing-native-api"
```

## sample
```go
	L := licensingApi.NewLicenseApi()

	/* login to default feature (0)                 */
	/* this default feature is available on any key */
	/* search for local and remote HASP key         */
	handle, err := L.HaspLogin(HASP_DEFAULT_FID, vendor_code)
	if err != 0 {
		fmt.Printf("login to default feature failed, err code : %d\n", err)
		return
	} else {
		fmt.Printf("login to default feature success\n")
	}

	/*
	 * hasp_get_sessioninfo
	 *   retrieve Sentinel key attributes
	 *
	 * Please note: In case of performing an activation we recommend to use
	 *              hasp_get_info() instead of hasp_get_sessioninfo(), as
	 *              demonstrated in the activation sample. hasp_get_info()
	 *              can be called without performing a login.
	 */
	format := "<haspformat format=\"keyinfo\"/>"
	info, err := L.HaspGetSessionInfo(handle, format)
	if err != 0 {
		fmt.Printf("get session info failed, err code : %d\n", err)
	} else {
		fmt.Printf("get session info success, %s\n", info)
	}

	/*
		     * hasp_get_size
		     *   retrieve the memory size of the HASP key
			 *   you can also retrieve dynamic memory file size,
			 *   only need to pass the dynamic memory file id which
			 *	 created in token
	*/
	var size uint
	err = L.HaspGetSize(handle, HASP_FILEID_RW, &size)
	if err != 0 {
		fmt.Printf("get size failed, err code : %d\n", err)
	} else {
		fmt.Printf("sentinel memory size is %d bytes\n", size)
	}
	/*
		    * hasp_write
		    *   write to HASP memory
			*   you can also write dynamic memory file,
			*   only need to pass the dynamic memory file id which
			*	 created in token
	*/
	var data [64]byte
	size = 64
	for i := 0; i < int(size); i++ {
		data[i] = byte(i)
	}
	err = L.HaspWrite(handle,
		HASP_FILEID_RW,
		0,    /* offset */
		size, /* length */
		&data[0])
	if err != 0 {
		fmt.Printf("write memory failed, err code : %d\n", err)
	} else {
		fmt.Printf("write memory success, data = %v\n", data)
	}
	/*
		    * hasp_read
		    *   read from HASP memory
			*   you can also read dynamic memory file,
			*   only need to pass the dynamic memory file id which
			*	 created in token
	*/
	err = L.HaspRead(handle,
		HASP_FILEID_RW, /* read/write file ID */
		0,              /* offset */
		size,           /* length */
		&data[0])       /* file data */
	if err != 0 {
		fmt.Printf("read memory failed, err code : %d\n", err)
	} else {
		fmt.Printf("read memory success, data = %v\n", data)
	}
	/*
	 * hasp_encrypt
	 *   encrypts a block of data using the HASP key
	 *   (minimum buffer size is 16 bytes)
	 */
	size = 64
	err = L.HaspEncrypt(handle, &data[0], size)
	if err != 0 {
		fmt.Printf("encrypt failed, err code : %d\n", err)
	} else {
		fmt.Printf("encrypt success, encrypted data = %v\n", data)
	}
	/*
	 * hasp_decrypt
	 *   decrypts a block of data using the HASP key
	 *   (minimum buffer size is 16 bytes)
	 */
	size = 64
	err = L.HaspDecrypt(handle, &data[0], size)
	if err != 0 {
		fmt.Printf("decrypt failed, err code : %d\n", err)
	} else {
		fmt.Printf("decrypt success, decrypted data = %v\n", data)
	}
	/*
	 * hasp_get_rtc
	 *   read current time from HASP Time key
	 */
	var time uint64
	err = L.HaspGetRtc(handle, &time)
	if err != 0 {
		fmt.Printf("get rtc failed, err code : %d\n", err)
	} else {
		fmt.Printf("get rtc success, time: %d\n", time)
	}
	/*
	 * hasp_logout
	 *   closes established session and releases allocated memory
	 */
	err = L.HaspLogout(handle)
	if err != 0 {
		fmt.Printf("logout failed, err code : %d\n", err)
	} else {
		fmt.Printf("logout success\n")
	}
```
