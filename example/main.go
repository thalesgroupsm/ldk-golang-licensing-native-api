package main

import (
	"fmt"
	"time"

	licenseApi "github.com/thalesgroupsm/ldk-golang-licensing-native-api"
)

func LoginByIdentity(client *licenseApi.LicenseApi, keyId string) {

	var scope string = "<?xml version=\"1.0\" encoding=\"UTF-8\" ?>	<haspscope>		 <hasp id=\"" + keyId + "\" />	</haspscope>"

	runtime_handle, ret := client.HaspLoginScope(0, scope, client.VendorCode)
	if runtime_handle != 0 || ret == 0 {

		var format string = "<haspformat format=\"keyinfo\"/>"
		sessionInfo := client.HaspGetSessionInfo(runtime_handle, format)
		fmt.Println(sessionInfo)
		fmt.Println("Login success,", keyId)
		time.Sleep(time.Duration(3) * time.Second)
		client.HaspLogout(runtime_handle)
	} else {
		fmt.Println("hasp_login_scope failed")
	}
}
func main() {
	fmt.Println("licensing native api sample")
	L := licenseApi.NewLicenseApi("37515")
	if L == nil {
		fmt.Println("invalid license api client")
	}
	LoginByIdentity(L, "792409087108542559")
}
