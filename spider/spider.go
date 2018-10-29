package main;

import (
	"fmt";
	"net/http";
	"io/ioutil";
	"regexp";
	//"time";
	"strconv";
	"os"
)


// h.p03.space/viewthread.php?tid=253453&page=21
/// http://h.p03.space/attachments/180827171132291b5106ef17dd.jpg
//http://h.p03.space/viewthread.php?tid=299789
//http://h.p03.space/viewthread.php?tid=263385&page=

func main() {
	for i := 299789; i > 200000; i-- {
		res, _ := getHtml("http://h.p03.space/viewthread.php?tid="+strconv.Itoa(i));
		fmt.Println("http://h.p03.space/viewthread.php?tid="+strconv.Itoa(i));

		savePagedImg(res, strconv.Itoa(i));
		
	} 
}


/*
获取当前页面图片列表并保存
**/
func savePagedImg(res string, dir string) {
	//获取列表
	imgList := getList(res);
	imgLen := len(imgList);
	if imgLen != 0 {
		os.Mkdir(dir, os.ModePerm);
		os.Chdir(dir);
		// a, _ := os.Getwd();
		// fmt.Println(a);
	}
	for index, value := range imgList {
		imgList[index] = "http://h.p03.space/"+value;
		saveRes(imgList[index]); //保存图片
		//time.Sleep(time.Second/2);
		fmt.Println(imgList[index]);
	}
	if imgLen != 0 {
	    os.Chdir("..");
	}
	fmt.Println(" ");
}

/*
获取91图片列表
**/
func getList(content string) []string {
	re := regexp.MustCompile(`attachments/[a-zA-Z0-9]*(.jpg|.png)`);
	return re.FindAllString(content, -1)
}


/*
获取网络资源
@params（url string：获取的源网址）
@return（[]byte：资源内容； error：读取错误）
**/
func getResource(url string) ([]byte, error) {
	res, err := http.Get(url);
	if err != nil {
		fmt.Println("Get URl Error");
		ioutil.WriteFile("error_log.txt", []byte(url), 0x755);
	}
	defer res.Body.Close();
	data, err := ioutil.ReadAll(res.Body);
	return data, err
}


/*
获取网页内容
@params（url string：获取的源网址）
@return（string：网页内容； error：读取错误）
**/
func getHtml(url string) (string, error) {
	res, err := getResource(url);
	return string(res), err
}

/*
获取并保存网络资源
**/
func saveRes(url string) {
	data, _ := getResource(url);
	ioutil.WriteFile(getName(url), data, 0x755);
}

/*
获取文件名
@params （url string：获取的包含文件名的链接字符串）
@return （string 文件名）
**/
func getName(url string) string {
	re := regexp.MustCompile(`[a-zA-Z0-9]*(.jpg|.png)`);
	return re.FindString(url);	
}



