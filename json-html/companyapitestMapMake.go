package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

/**
利用反射包，获取类型
*/
func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

/**
路由处理函数
*/
func testHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 获取请求参数， 运行这一方法后， r.Form自动存入参数键值对，为map类型，可以通过range遍历
	fmt.Println(r.Form)

	str := getSiteData(r.Form)

	// // Stop here if its Preflighted OPTIONS request
	// if origin := r.Header.Get("Origin"); origin != "" {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers",
	// "Action, Module")   //有使用自定义头 需要这个,Action, Module是例子
	// }
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method == "OPTIONS" {
		return
	}
	if r.Method == "POST" {
		w.Write([]byte(str))
	}
	if r.Method == "GET" {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
	}

}

func getSiteData(requestParams url.Values) []byte {
	var (
		BaseUrl string
		Api     string
	)
	BaseUrl = "http://vis-screen-fnw-dev.tipaas.enncloud.cn"
	Api = BaseUrl + "/web/site.json" // 可以用本目录下的api.json代替

	var stationInfos map[string]interface{} // json解析后的变量

	resp, err := http.Get(Api) // get请求接口，获取json数据
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()              // 最终记得关闭请求体
	body, _ := ioutil.ReadAll(resp.Body) // 读取响应对 body部分，为二进制数据
	json.Unmarshal(body, &stationInfos)  // 解析json

	res := stationInfos["obj"].([]interface{}) // 截取json中需要的部分 obj键对应的值

	for rkey, rvalue := range requestParams { // 遍历request请求过来的参数，进行几轮数据筛选过程
		res = getArea(res, rkey, rvalue[0])
	}
	fmt.Println("---------------")
	fmt.Println(typeof(stationInfos["obj"]))
	fmt.Println(typeof(res))
	fmt.Println("---------------")
	// for i, item := range res {
	// 	fmt.Println(i, item.(map[string]interface{})["websiteName"]);
	// }
	jsonMap := make(map[string]interface{}) // 上述res筛选结果为数组，开辟一个 map，容纳这个筛选出来对结果数组
	jsonMap["obj"] = res
	jsonStr, _ := json.Marshal(jsonMap) // map转化生成json
	//fmt.Printf("%s\n", jsonStr)
	return jsonStr
}

/**
筛选数据方法
输入输出类型一致，适合链式调用
*/
func getArea(data []interface{}, filterKey string, filterValue string) []interface{} {
	var res []interface{} //定义切片
	//res := make([]interface{}, 100, 500) // 开辟数组切片，大小100元素， 容纳能力500
	//var sum int = 0                      // 最终切片内容大小计数，以便将多余空间切除
	for _, dataItem := range data {
		//fmt.Println(i, dataItem.(map[string]interface{})["websiteName"])
		if dataItem.(map[string]interface{})[filterKey] == filterValue {
			//res[sum] = dataItem.(map[string]interface{})
			res = append(res, dataItem.(map[string]interface{}))
			//sum++
		}
	}
	//res = res[:sum] // 切除多余空间
	return res
}

func main() {
	http.HandleFunc("/test", testHandler)
	http.ListenAndServe(":8080", nil)
}
