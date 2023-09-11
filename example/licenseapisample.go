package main

import (
	"fmt"
	"time"
)

func (l *LicenseApi) LoginByIdentity(keyId string) {

	var scope string = "<?xml version=\"1.0\" encoding=\"UTF-8\" ?>	<haspscope>		 <hasp id=\"" + keyId + "\" />	</haspscope>"

	runtime_handle, ret := l.HaspLoginScope(0, scope, l.VendorCode)
	if runtime_handle != 0 || ret == 0 {

		var format string = "<haspformat format=\"keyinfo\"/>"
		sessionInfo := l.HaspGetSessionInfo(runtime_handle, format)
		fmt.Println(sessionInfo)
		fmt.Println("Login success,", keyId)
		time.Sleep(time.Duration(3) * time.Second)
		l.HaspLogout(runtime_handle)
	} else {
		fmt.Println("hasp_login_scope failed")
	}
}
