package main

import (
	"fmt"
	"strings"

	"github.com/intelfike/lib/my"
)

func main() {
	cmd := []string{"reg", "query", "HKEY_CLASSES_ROOT\\Directory\\Background\\shell\\cmd", "/v", "Extended"}
	str, _ := my.AutoEnc(my.Execute(cmd))

	if !strings.Contains(str, "エラー") {
		fmt.Println("右クリックを拡張します。")
		cmd = []string{"reg", "delete", "HKEY_CLASSES_ROOT\\Directory\\Background\\shell\\cmd", "/v", "Extended", "/f"}
	} else {
		fmt.Println("右クリックの拡張を終了します。")
		cmd = []string{"reg", "add", "HKEY_CLASSES_ROOT\\Directory\\Background\\shell\\cmd", "/v", "Extended", "/f"}
	}

	str, _ = my.AutoEnc(my.Execute(cmd))
	fmt.Println(str)

	fmt.Print("Enterで終了:")
	var s string
	fmt.Scanln(&s)
}
