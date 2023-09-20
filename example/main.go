package main

import (
	"fmt"
	licenseApi "github.com/thalesgroupsm/ldk-golang-licensing-native-api"
)

/*vendor code for DemoMA*/
var vendor_code string = "AzIceaqfA1hX5wS+M8cGnYh5ceevUnOZIzJBbXFD6dgf3tBkb9cvUF/Tkd/iKu2fsg9wAysY" +
	"Kw7RMAsVvIp4KcXle/v1RaXrLVnNBJ2H2DmrbUMOZbQUFXe698qmJsqNpLXRA367xpZ54i8k" +
	"C5DTXwDhfxWTOZrBrh5sRKHcoVLumztIQjgWh37AzmSd1bLOfUGI0xjAL9zJWO3fRaeB0NS2" +
	"KlmoKaVT5Y04zZEc06waU2r6AU2Dc4uipJqJmObqKM+tfNKAS0rZr5IudRiC7pUwnmtaHRe5" +
	"fgSI8M7yvypvm+13Wm4Gwd4VnYiZvSxf8ImN3ZOG9wEzfyMIlH2+rKPUVHI+igsqla0Wd9m7" +
	"ZUR9vFotj1uYV0OzG7hX0+huN2E/IdgLDjbiapj1e2fKHrMmGFaIvI6xzzJIQJF9GiRZ7+0j" +
	"NFLKSyzX/K3JAyFrIPObfwM+y+zAgE1sWcZ1YnuBhICyRHBhaJDKIZL8MywrEfB2yF+R3k9w" +
	"FG1oN48gSLyfrfEKuB/qgNp+BeTruWUk0AwRE9XVMUuRbjpxa4YA67SKunFEgFGgUfHBeHJT" +
	"ivvUl0u4Dki1UKAT973P+nXy2O0u239If/kRpNUVhMg8kpk7s8i6Arp7l/705/bLCx4kN5hH" +
	"HSXIqkiG9tHdeNV8VYo5+72hgaCx3/uVoVLmtvxbOIvo120uTJbuLVTvT8KtsOlb3DxwUrwL" +
	"zaEMoAQAFk6Q9bNipHxfkRQER4kR7IYTMzSoW5mxh3H9O8Ge5BqVeYMEW36q9wnOYfxOLNw6" +
	"yQMf8f9sJN4KhZty02xm707S7VEfJJ1KNq7b5pP/3RjE0IKtB2gE6vAPRvRLzEohu0m7q1aU" +
	"p8wAvSiqjZy7FLaTtLEApXYvLvz6PEJdj4TegCZugj7c8bIOEqLXmloZ6EgVnjQ7/ttys7VF" +
	"ITB3mazzFiyQuKf4J6+b/a/Y"

const (
	HASP_DEFAULT_FID = 0
	HASP_FILEID_RW   = 0xfff4
)

func hasp_demo() {

	L := licenseApi.NewLicenseApi()

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

}

func hasp_update() {

	L := licenseApi.NewLicenseApi()

	/*
	 * hasp_get_info
	 *   get key info
	 */
	scope := "<haspscope>" +
		"  <license_manager hostname=\"localhost\" />" +
		"</haspscope>"
	format := "<haspformat root=\"location\">" +
		"  <license_manager>" +
		"    <attribute name=\"id\" />" +
		"    <attribute name=\"time\" />" +
		"    <element name=\"hostname\" />" +
		"    <element name=\"version\" />" +
		"    <element name=\"host_fingerprint\" />" +
		"  </license_manager>" +
		"</haspformat>"

	recipient, err := L.HaspGetInfo(scope, format, vendor_code)
	if err != 0 {
		fmt.Printf("get info failed, err code : %d\n", err)
	} else {
		fmt.Printf("get info success, info = %s\n", recipient)
	}
	/* detach license for local recipient (duration 120 seconds.)
	 * Note: Please change the product ID to which you want to
	 *       operate.
	 */
	action := "<?xml version=\"1.0\" encoding=\"UTF-8\" ?>" +
		"<detach><product id=\"123\"><duration>120</duration></product></detach>"
	scope = "<haspscope><license_manager hostname=\"localhost\" /></haspscope>"
	output, err := L.HaspTransfer(action, scope, vendor_code, recipient)
	if err != 0 {
		fmt.Printf("tranfter failed, err code : %d\n", err)
	} else {
		fmt.Printf("tranfter success, output = %s\n", output)
	}
	/*
	 * hasp_update
	 *   update the V2C
	 */
	v2c := output
	output, err = L.HaspUpdate(v2c)
	if err != 0 {
		fmt.Printf("update failed, err code : %d  ", err)
	} else {
		fmt.Printf("update success")
	}

}

func main() {
	hasp_demo()
	hasp_update()
}
