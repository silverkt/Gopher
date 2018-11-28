package main;

import (
	"net/http";
	"net/url";
	"encoding/json";
	"io/ioutil";
	"fmt";
	"reflect";
)


func typeof(v interface{}) string {
    return reflect.TypeOf(v).String()
}


type Person struct {
	Name string  `json:"myname"`;
	Title string  `json:"mytitle"`;
	Age int  `json:"myagg"`;
}



func testHandler(w http.ResponseWriter, r *http.Request) {
	
	// result := Person{
	// 	Name: "silver",    
	// 	Title: "Front-End-Develp",
	// 	Age: 18,
	// }

	// str, _ := json.Marshal(result);

	r.ParseForm();
	fmt.Println(r.Form);

	str := getSiteData(r.Form);

	// // Stop here if its Preflighted OPTIONS request
    // if origin := r.Header.Get("Origin"); origin != "" {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers",
	// "Action, Module")   //有使用自定义头 需要这个,Action, Module是例子
	// }
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type","application/json");
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method == "OPTIONS" {
		return
	};
	if r.Method == "POST" {
		w.Write([]byte(str));
	};
	if r.Method == "GET" {
		w.WriteHeader(405);
		w.Write([]byte("Method Not Allowed"));
	}
	
}


func getSiteData(requestParams url.Values) []byte {
	var (
		BaseUrl string;
		Api string;
	)
	

	// var filtedData []map[string]string;
	//filtedData = make([]map[string]string);
	
	BaseUrl = "http://vis-screen-fnw-dev.tipaas.enncloud.cn";
	Api = BaseUrl + "/web/site.json";   //可以用本目录下的api.json代替
	var stationInfos map[string]interface{};
	

	resp, err := http.Get(Api);

	if err != nil {
		// handle error
	}
	defer resp.Body.Close();
	body, _ := ioutil.ReadAll(resp.Body);
	json.Unmarshal(body, &stationInfos);

	jsonMap := make(map[string]interface{});
	// jsonMap["obj"] = getArea(stationInfos["obj"].([]interface{}), "area", "华北");
	// jsonStr, _ := json.Marshal(jsonMap);
	// fmt.Print([]byte(jsonStr));
	
	res := stationInfos["obj"].([]interface{});
	for rkey, rvalue := range requestParams {
		res = getArea(res, rkey, rvalue[0]);
	}

	// res := getArea(stationInfos["obj"].([]interface{}), "area", "华北");
	// res = getArea(res, "websiteName", "廊坊市新朝阳泛能微网");
	fmt.Println("---------------");
	fmt.Println(typeof(stationInfos["obj"]));
	fmt.Println(typeof(res));
	fmt.Println("---------------");
	// for i, item := range res {
	// 	fmt.Println(i, item.(map[string]interface{})["websiteName"]);
	// }
	jsonMap["obj"] = res;
	jsonStr, _ := json.Marshal(jsonMap);
	//fmt.Printf("%s\n", jsonStr)

	return jsonStr;

}



func getArea(data []interface{}, filterKey string, filterValue string) []interface{} {	
	res := make([]interface{}, 200);
	var sum int = 0;
	for _, dataItem := range data {
		//fmt.Println(i, dataItem.(map[string]interface{})["websiteName"])
		if dataItem.(map[string]interface{})[filterKey] == filterValue {
			res[sum] = dataItem.(map[string]interface{});
			sum ++;
		}		
	};
	res = res[:sum];
	return res;
}






func main() {
	http.HandleFunc("/test", testHandler);
	http.ListenAndServe(":8080", nil);
}