package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
)

func init() {
	flag.Parse()
}

// htmlを読み取って、すべてのjsファイルを統合してコンパイルするプログラム
func main() {
	arglen := len(flag.Args())
	if arglen == 0 {
		fmt.Println("引数を入力してください\n_(アンダーバー)で要素指定を省略できます")
		return
	}
	elem := flag.Arg(0)
	attr := flag.Arg(1)
	d := xml.NewDecoder(os.Stdin)
	// Decoderの設定
	d.Strict = false
	d.AutoClose = xml.HTMLAutoClose
	d.Entity = xml.HTMLEntity
	// XMLを順次解析
	for {
		token, err := d.Token()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			panic(err)
		}
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			if tok.Name.Local != elem && elem != "_" {
				continue
			}
			for _, v := range tok.Attr {
				if v.Name.Local != attr {
					continue
				}
				fmt.Println(v.Value)
			}
		case xml.EndElement:
		case xml.CharData:
		case xml.Comment:
		case xml.ProcInst:
		case xml.Directive:
		default:
			panic("unknown xml token.")
		}
	}
}
