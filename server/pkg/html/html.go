package html

import (
	"fmt"
	"strings"

	"bbs-go/pkg/text"

	"github.com/PuerkitoBio/goquery"
	"github.com/mlogclub/simple/common/strs"
	"github.com/sirupsen/logrus"
)

func GetSummary(htmlStr string, summaryLen int) string {
	if summaryLen <= 0 || strs.IsEmpty(htmlStr) {
		return ""
	}
	//return text.GetSummary(htmlStr, summaryLen)
	return text.GetSummary(GetHtmlText(htmlStr), summaryLen) //这是将简介里的图片摘出来的
}

// GetHtmlText 获取html文本
func GetHtmlText(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		logrus.Error(err)
		return ""
	}
	fmt.Println("---------------------")
	fmt.Println(html)
	fmt.Println(doc.Text())
	return doc.Text()
}
