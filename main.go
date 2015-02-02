package main

import (
           "strings"
           "github.com/PuerkitoBio/goquery"
           "github.com/dongjujang/slackchat-cus"
)

func chat_downlink() {
  webhook_url := "WEBHOOK URL"

  doc, _ := goquery.NewDocument("http://www.torrentbest.net/bbs/board.php?bo_table=torrent_kortv_ent")

  doc.Find("td.subject").Each(func(i int, s *goquery.Selection) {
 //   subject := s.Find("a").Text()
    attr, _ := s.Find("a").Attr("href")

    str := "http://www.torrentbest.net"
    substr := string([]byte(attr[2:]))
    url := str + substr

    doc2, _ := goquery.NewDocument(url)
    
    doc2.Find("td.view_file").Each(func(j int, s2 *goquery.Selection) {
      downlink, _ := s2.Find("a").Attr("href")
      if strings.Contains(downlink, "download") {
        data := strings.Split(downlink, "'")
        path_temp := data[1]
        path := string([]byte(path_temp[1:]))
        link := str + "/bbs" + path

//      com := "<" + link + "|" + subject + ">"
        slack_chat := slack.BuildPayload("#test-channel", "Korea-tv", "", "", link)
        slack.Post(webhook_url, slack_chat)
      }
    })
  })  
}

func main() {
  chat_downlink()
}
