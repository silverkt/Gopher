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
	"time"
	"sync"
	"math/rand"
)

// h.p03.space/viewthread.php?tid=253453&page=21
/// http://h.p03.space/attachments/180827171132291b5106ef17dd.jpg
//http://h.p03.space/viewthread.php?tid=299789
//http://h.p03.space/viewthread.php?tid=263385&page=21
// 
// 315083
var wg sync.WaitGroup //添加协程锁 计数器

func main() {
	// 获取外部json里面变量
	var ch map[string]string
	data, _ := ioutil.ReadFile("config.json")
	json.Unmarshal([]byte(data), &ch)

	mainurl := ch["baseURL"] + ch["attach"]
	from, _ := strconv.Atoi(ch["from"])
	to, _ := strconv.Atoi(ch["to"])

	//最大翻页间隔 单位为秒
	pageinterval, _ := strconv.Atoi(ch["pageinterval"])
	//最大定期睡眠时长 单位为分
	sleeptime, _ := strconv.Atoi(ch["sleeptime"])
	//最大多长间隔定时睡眠 单位为分
	biginterval, _ := strconv.Atoi(ch["biginterval"])

	//计时开始
	t := time.Now().Unix()
	//计数开始， 设定为20次翻页即检查下时间间隔是否达到预定随眠时间
	count := 0
	for i := to; i > from; i-- {
		count ++ //翻页计数
		fmt.Println(mainurl + strconv.Itoa(i))
		// res, _ := getHtml(mainurl + strconv.Itoa(i)+"&_="+strconv.Itoa(t))	
		res, _ := getHtml(mainurl + strconv.Itoa(i))
		savePagedImg(res, strconv.Itoa(i), ch["baseURL"], ch["regexp"])

		pageinterval1 := rand.Intn(pageinterval) //设置随机翻页间隔时间
		fmt.Println("page interval seconds:", pageinterval1)
		time.Sleep(time.Duration(pageinterval1) * time.Second) //翻页间隔
		//如果翻页20次并且没有设置不睡眠
		if count == 20 && sleeptime !=0 {
			count = 0 //计数重置
			//如果时间间隔大于设定的 大段间隔睡眠时间 即开始睡眠，此举是为了避免服务端统计特定时间段内流量并且封ip
			biginterval1 := rand.Intn(biginterval) //设置随机睡眠间隔
			fmt.Println("big interval minutes:", biginterval1)
			if int(time.Now().Unix() - t) > biginterval1 * 60 {
				t = time.Now().Unix() //计时重置
				sleeptime1 := rand.Intn(sleeptime) //设置随机睡眠时间
				fmt.Println("sleep time minutes:", sleeptime1)
				time.Sleep(time.Duration(sleeptime1) * time.Minute)
			}
		}
	}
}

/*
获取当前页面图片列表并保存
**/
func savePagedImg(res string, dir string, baseURL string, exp string) {
	//获取列表
	imgList := getList(res, exp)
	imgLen := len(imgList)
	if imgLen != 0 {
		os.Mkdir(dir, os.ModePerm)
	}
	for index, value := range imgList {
		imgList[index] = baseURL + "/" + value
		wg.Add(1) //协程计数加一
		go saveRes(imgList[index], dir) //保存图片
		fmt.Println(imgList[index])
	}
	wg.Wait() // 阻塞main协程， 不然最后一个 1的 协程没来得及结束，main就结束了 
	fmt.Println(" ")
}

/*
获取91图片列表
**/
func getList(content string, exp string) []string {
	re := regexp.MustCompile(exp)
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
	req.Header.Set("User-Agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36`)
	//req.Header.Set("User-Agent", `Mozilla/5.0 (Linux; U; Android 7.1.2; zh-cn; Redmi 5 Plus Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/61.0.3163.128 Mobile Safari/537.36 XiaoMi/MiuiBrowser/9.7.2`)
	req.Header.Set("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8`)
	req.Header.Set("Accept-Encoding", `gzip, deflate`)
	req.Header.Set("Accept-Language", `zh-CN,en-US;q=0.9`)
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
func saveRes(url string, path string) {
	data, _ := getResource(url)
	ioutil.WriteFile((path+"/"+getName(url)), data, 0x755)
	wg.Done() //协程计数减一
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("goRoutine Panic happend")
		}
	}()
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
