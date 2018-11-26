package main;

import (
	"net/http";
	"encoding/json";
	"io/ioutil";
	"fmt";
)





type Person struct {
	Name string  `json:"myname"`;
	Title string  `json:"mytitle"`;
	Age int  `json:"myagg"`;
}



func testHandler(w http.ResponseWriter, r *http.Request) {
	
	result := Person{
		Name: "silver",    
		Title: "Front-End-Develp",
		Age: 18,
	}

	str, _ := json.Marshal(result);


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


func getSiteData() {
	var (
		BaseUrl string;
		Api string;
	)
	

	// var filtedData []map[string]string;
	//filtedData = make([]map[string]string);
	
	BaseUrl = "http://vis-screen-fnw-dev.tipaas.enncloud.cn";
	Api = BaseUrl + "/web/site.json";
	var stationInfos map[string]interface{};
	

	resp, err := http.Get(Api);

	if err != nil {
		// handle error
	}
	defer resp.Body.Close();
	body, _ := ioutil.ReadAll(resp.Body);
	json.Unmarshal(body, &stationInfos);

	// jsonMap := make(map[string]interface{});
	// jsonMap["obj"] = getArea(stationInfos["obj"].([]interface{}), "area", "华北");
	// jsonStr, _ := json.Marshal(jsonMap);
	// fmt.Print([]byte(jsonStr));


	res := getArea(stationInfos["obj"].([]interface{}), "area", "华北");
	for i, item := range res {
		fmt.Println(i, item["websiteName"]);
	}
	// res = getArea(([]interface{})res, "websiteName", "天津京滨工业园泛能区域网");
	// for i, item := range res {
	// 	fmt.Println(i, item["company"]);
	// }	

	// for i , stationItem := range (stationInfos["obj"]).([]interface{}) {
	// 	fmt.Println(i, stationItem.(map[string]interface{})["websiteName"]);
	// }
	//fmt.Println(stationInfos["obj"].([]interface{})[3].(map[string]interface{})["websiteName"]);


}



func getArea(data []interface{}, filterKey string, filterValue string) []map[string]interface{} {
	
	res := make([]map[string]interface{}, 200);
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
	 

	getSiteData();
	// http.HandleFunc("/test", testHandler);
	// http.ListenAndServe(":8080", nil);
}