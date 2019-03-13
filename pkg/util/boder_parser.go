package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func RequestBodyAsMap(req *http.Request) *map[string]interface{} {
	body, _ := ioutil.ReadAll(req.Body)
	var reqBody map[string]interface{}
	err := json.Unmarshal(body, &reqBody)
	if err == nil {
		fmt.Println(reqBody)
		return &reqBody
	} else {
		fmt.Println("TESTTTINGG")
		emptyMap := &map[string]interface{}{}
		return emptyMap
	}
}
