package main;

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func main() {
	data, _ := ioutil.ReadFile("config.json");
	var dbgInfos map[string]string

	json.Unmarshal([]byte(data), &dbgInfos)


	fmt.Println(dbgInfos);
	fmt.Println(dbgInfos["url"]);
}