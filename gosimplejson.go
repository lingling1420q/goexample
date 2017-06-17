package main

import (
	"fmt"
	simplejson "github.com/bitly/go-simplejson"
)

func main() {
	sjson := `{"cost":0.010000000707805157,"data":{"preview":393531,"key":"7956ed04c52fce2e261fbbfa","fileid":"0300020600592E5BD23661065745D8841D2841-9293-6A46-6E13-A1DA3","desc":""}}`
	bjson := []byte(sjson)
	json, err := simplejson.NewJson(bjson)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(json.Get("data").Get("key").String())

}
