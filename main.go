package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/antchfx/xmlquery"
)

func repRight(s string) string {
	return strings.ReplaceAll(s, "</div>", "")
}

func repLeft(s string, k string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "amp;", ""), k+"(请开启代理后再拉取)：", "")
}

func main() {
	var url = "http://feeds.feedburner.com/mattkaydiary/pZjG"

	node, err := xmlquery.LoadURL(url)
	if err != nil {
		panic(err)
	}

	str := xmlquery.FindOne(node, "//channel/item[3]/description")

	v2ray := regexp.MustCompile(`v2ray\(请开启代理后再拉取\)：(.+?)</div>`)
	clash := regexp.MustCompile(`clash\(请开启代理后再拉取\)：(.+?)</div>`)

	v2rayHttp := repLeft(repRight(v2ray.FindAllString(str.InnerText(), -1)[0]), "v2ray")
	clashHttp := repLeft(repRight(clash.FindAllString(str.InnerText(), -1)[0]), "clash")

	read(v2rayHttp, "./v2ray.txt")
	read(clashHttp, "./clash.yml")
}

func read(s string, f string) {
	response, err := http.Get(s)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))

	ioFile(f, body)
}

func ioFile(f string, body []byte) {
	if err := os.Remove(f); err != nil {
		panic(err)
	}

	file, err := os.Create(f)
	if err != nil {
		panic(err)
	}

	if _, err := file.Write(body); err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
}
