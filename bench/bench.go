package main

import (
	"flag"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("時間を計測したいコマンドを入力してください。")
		return
	}

	command := flag.Arg(0)
	s := time.Now()
	out, err := exec.Command(command, flag.Args()[1:]...).CombinedOutput()
	fmt.Println("done.")

	if len(out) == 0 && err != nil {
		str := "'" + command + "'は、操作可能なプログラムまたはバッチファイルとして認識されていません。\n" +
			"スペルミスもしくは環境変数Pathを確認してください。"
		fmt.Println(str)
	} else {
		fmt.Print(time.Now().Sub(s))
	}
}
