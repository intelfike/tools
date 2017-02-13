package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
)

var levelList = []string{"WHITESPACE_ONLY", "SIMPLE", "ADVANCED"}
var (
	level int
)

func init() {
	flag.IntVar(&level, "level", 1,
		`最適化のレベルです。
	0 = 空白削除のみ
	1 = シンプルに最適化
	2 = 高度な最適化(デッドコード削除あり)
	※2の場合はonClickでの呼び出しなどもデッドコードとして判断され、削除されます。
`,
	)
	flag.Parse()
}

// htmlを読み取って、すべてのjsファイルを統合してコンパイルするプログラム
func main() {
	b1, _ := ioutil.ReadAll(os.Stdin)
	reg, _ := regexp.Compile("[\r\n]+")
	list := reg.Split(string(b1), -1)
	// コンパイルする文字列
	js := []byte{}
	//
	for _, v := range list {
		if v == "" {
			continue
		}
		b2, err := getTextURI(v)
		fmt.Print(v)
		if err != nil {
			fmt.Println("\t", err)
		} else {
			fmt.Println()
		}
		if len(b2) == 0 {
			continue
		}
		b2 = append(b2, byte('\n'))
		js = append(js, b2...)
	}

	cmd := exec.Command("java", "-jar", "S:\\compiler-latest\\closure-compiler-v20161201.jar",
		"--compilation_level", levelList[level],
	)
	cmd.Stderr = os.Stderr
	// 入出力のパイプを取得
	cmdin, _ := cmd.StdinPipe()
	defer cmdin.Close()
	cmdout, _ := cmd.StdoutPipe()
	defer cmdout.Close()
	// コンパイル開始
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
	// データ送信
	cmdin.Write([]byte(js))
	cmdin.Close()
	// データ受信
	b, _ := ioutil.ReadAll(cmdout)
	// 受信データを新しいファイルに保存
	jsfile, _ := os.Create("compiled.js")
	jsfile.Write(b)
}

// URI(ファイルパスやURL)からそのテキストを取得
// ファイルパス優先
func getTextURI(path string) ([]byte, error) {
	// ファイルパスを試す
	b, err := ioutil.ReadFile(path)
	if err == nil {
		return b, nil
	}
	// httpを試す
	res, err := http.Get(path)
	if err != nil {
		// 失敗
		return nil, errors.New("ファイルパス、URLのどちらも無効でした")
	}
	defer res.Body.Close()

	b2, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// 想定してない
		return nil, errors.New("URLから正しくデータを取得出来ませんでした")
	}
	return b2, nil
}
