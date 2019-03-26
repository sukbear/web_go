package main

import (
	"fmt"
	"strings"
	"os"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"io"
	"time"
	"log"
	"runtime"
)

//var url = "http://www.umei.cc/"
var url = "http://www.sccnn.com/"
var c chan int

func match(image string) {
	fmt.Println(image)
}

func getData(url string) (eader io.Reader, err error) {
	req := buildRequest(url)
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	return io.Reader(resp.Body), err
}

func imageRule(doc *goquery.Document, f func(image string)) (urls []string) {
	str := make([]string, 0)
	//直接找以img 开头的标签 过滤掉不符合规则的url 即可
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		url, result := s.Attr("src")
		if result {
			if strings.HasSuffix(url, ".jpg") {
				str = append(str, url)
			}
		}
	})
	return str
}

//根据url 创建http 请求的 request
//网站有反爬虫策略 wireshark 不解释
func buildRequest(url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Referer", url)
	return req
}

// 判读文件夹是否存在
func isExist(dir string) bool {
	_, err := os.Stat(dir)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

// 通过url 得到图片名字
func getNameFromUrl(url string) string {
	arr := strings.Split(url, "/")
	return arr[len(arr)-1]
}

// 得到一个网页中所有 ImageUrl
func parseImg(reader io.Reader) (res []string, err error) {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}
	fmt.Println(doc.Url)
	imageRule(doc, func(image string) {
		res = append(res, image)
	})
	return res, nil
}

//download
func downloadImage(url string) {
	fileName := getNameFromUrl(url)
	req := buildRequest(url)
	http.DefaultClient.Timeout = 10 * time.Second;
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("failed download ")
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("failed download " + url)
		return
	}
	defer func() {
		resp.Body.Close()
		if r := recover(); r != nil {
			fmt.Println(r)
		}
		c <- 0
	}()
	fmt.Println("begin download " + fileName)
	os.MkdirAll("./images/", 0777)
	localFile, _ := os.OpenFile("./images/"+fileName, os.O_CREATE|os.O_RDWR, 0777)

	if _, err := io.Copy(localFile, resp.Body); err != nil {
		panic("failed save " + fileName)
	}

	fmt.Println("success download " + fileName)
}

//url->Document->所有图片url->开启多线程进行下载->保存到本地
func spider() {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	urls := imageRule(doc, match)
	fmt.Println("共解析到", len(urls), "图片地址")
	c = make(chan int)
	for _, s := range urls {
		fmt.Println(s)
		go downloadImage(s)
	}
	//可以等待一会儿，留时间给子goroutine 执行
	//但是这种方式不怎么靠谱 //直接采用chan 的方式
	//time.Sleep(1e9*10)
	for i := 0; i < len(urls); i++ {
		<-c
	}
}
func main() {
	runtime.GOMAXPROCS(4)
	spider()
}
