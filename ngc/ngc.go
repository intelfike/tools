package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	fmt.Println("\a")
	kateToURL := map[string]string{}
	count := 0
	docRank, _ := goquery.NewDocument("http://www.nicovideo.jp/ranking")
	docRank.Find(".stack").Each(func(_ int, s *goquery.Selection) {
		first := true
		s.Find("a").Each(func(_ int, s *goquery.Selection) {
			if first {
				fmt.Println(s.Text())
				first = false
				return
			}
			count++
			fmt.Println("\t", count, s.Text())
			urlstr, _ := s.Attr("href")
			kateToURL[strconv.Itoa(count)] = urlstr
		})
	})
	fmt.Print("\n\nカテゴリを選んでください:")
	cate := scan()
	cateURL, ok := kateToURL[cate]
	if !ok {
		fmt.Println("cate:", cate)
		fmt.Println("入力されたカテゴリが存在しません。")
		return
	}

	rankToURL := map[string]string{}
	doc, _ := goquery.NewDocument(cateURL)
	doc.Find(".videoRanking").Each(func(_ int, s *goquery.Selection) {
		if s.Find(".new").Size() == 0 {
			return
		}
		rank := s.Find(".rankingNum").Text()
		title := s.Find(".itemTitle a")
		fmt.Println(rank, title.Text())
		link, _ := title.Attr("href")
		rankToURL[rank] = "http://www.nicovideo.jp/" + link
	})
	for {
		fmt.Print("\nランキングの番号を入力すると再生します:")
		s := scan()
		urlstr, ok := rankToURL[s]
		if !ok {
			fmt.Println("存在しない番号です。終了します。")
			return
		}
		fmt.Println(urlstr)
		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("cmd", "/c", "start", urlstr)
		} else if runtime.GOOS == "linux" {
			cmd = exec.Command("bash", "-c", "xdg-open", urlstr)
		}
		cmd.Start()
	}

}

func scan() string {
	s := ""
	fmt.Scan(&s)
	return s
}
