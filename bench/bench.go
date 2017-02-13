package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

var (
	d *bool
)

func init() {
	d = flag.Bool("d", false, "指定すると入出力します。")
	flag.Parse()
}

func main() {

	if len(flag.Args()) == 0 {
		fmt.Println("時間を計測したいコマンドを入力してください。")
		return
	}

	command := flag.Arg(0)
	s := time.Now()
	// out, err := exec.Command(command, flag.Args()[1:]...).CombinedOutput()
	cmd := exec.Command(command, flag.Args()[1:]...)
	if *d {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err := cmd.Run()

	if err != nil {
		str := "'" + command + "'は、操作可能なプログラムまたはバッチファイルとして認識されていません。\n" +
			"スペルミスもしくは環境変数PATHを確認してください。"
		fmt.Println(str)
	} else {
		fmt.Println("\n\ndone.")
		fmt.Println(time.Now().Sub(s))
	}
}
