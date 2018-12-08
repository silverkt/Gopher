package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"compress/gzip"
	"io"
//	"time"
)

// h.p03.space/viewthread.php?tid=253453&page=21
/// http://h.p03.space/attachments/180827171132291b5106ef17dd.jpg
//http://h.p03.space/viewthread.php?tid=299789
//http://h.p03.space/viewthread.php?tid=263385&page=21
// 299789
// 307362

func main() {
	// 获取外部json里面变量
	var ch map[string]string
	data, _ := ioutil.ReadFile("config.json")
	json.Unmarshal([]byte(data), &ch)

	mainurl := ch["baseURL"] + ch["attach"]
	from, _ := strconv.Atoi(ch["from"])
	to, _ := strconv.Atoi(ch["to"])
	
	t := 999999999999
	t = 0
	for i := to; i > from; i-- {
		//t := strconv.Itoa(int(time.Now().Unix()))
		t = t + 6000
		fmt.Println(mainurl + strconv.Itoa(i)+"&_="+strconv.Itoa(t))
		// res, _ := getHtml(mainurl + strconv.Itoa(i)+"&_="+strconv.Itoa(t))
		// savePagedImg(res, strconv.Itoa(i), ch["baseURL"])
	}
}

/*
获取当前页面图片列表并保存
**/
func savePagedImg(res string, dir string, baseURL string) {
	//获取列表
	imgList := getList(res)
	imgLen := len(imgList)
	if imgLen != 0 {
		os.Mkdir(dir, os.ModePerm)
		os.Chdir(dir)
	}
	for index, value := range imgList {
		imgList[index] = baseURL + "/" + value
		saveRes(imgList[index]) //保存图片
		fmt.Println(imgList[index])
	}
	if imgLen != 0 {
		os.Chdir("..")
	}
	fmt.Println(" ")
}

/*
获取91图片列表
**/
func getList(content string) []string {
	re := regexp.MustCompile(`attachments/[a-zA-Z0-9]*(.jpg|.png|.gif)`)
	return re.FindAllString(content, -1)
}

/*
获取网络资源
@params（url string：获取的源网址）
@return（[]byte：资源内容； error：读取错误）
**/
func getResource(url string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Connection", `keep-alive`)
	req.Header.Set("Upgrade-Insecure-Requests", `1`)
	req.Header.Set("User-Agent", `Mozilla/5.0 (Linux; U; Android 7.1.2; zh-cn; Redmi 5 Plus Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/61.0.3163.128 Mobile Safari/537.36 XiaoMi/MiuiBrowser/9.7.2`)
	req.Header.Set("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8`)
	req.Header.Set("Accept-Encoding", `gzip, deflate`)
	req.Header.Set("Accept-Language", `zh-CN,en-US;q=0.8`)
	req.Header.Set("Cache-Control", `max-age=0`)
	//req.Header.Set("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36`)
	
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Get URl Error")
		ioutil.WriteFile("error_log.txt", []byte(url), 0x755)
	}
	defer res.Body.Close()

	//网页gzip之后直接读取为乱码
	var reader io.ReadCloser
	if res.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(res.Body)
		if err != nil {
			 //
		}
	} else {
		reader = res.Body
	}

	data, err := ioutil.ReadAll(reader)
	return data, err
}

/*
获取网页内容
@params（url string：获取的源网址）
@return（string：网页内容； error：读取错误）
**/
func getHtml(url string) (string, error) {
	res, err := getResource(url)
	return string(res), err
}

/*
获取并保存网络资源
**/
func saveRes(url string) {
	data, _ := getResource(url)
	ioutil.WriteFile(getName(url), data, 0x755)
}

/*
获取文件名
@params （url string：获取的包含文件名的链接字符串）
@return （string 文件名）
**/
func getName(url string) string {
	re := regexp.MustCompile(`[a-zA-Z0-9]*(.jpg|.png)`)
	return re.FindString(url)
}
