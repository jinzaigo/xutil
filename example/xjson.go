package main

import (
	"fmt"
	"github.com/jinzaigo/xutil/xjson"
)

type ProductInfo struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func main() {
	str := "{\"name\":\"AppleWatchS8\",\"price\":\"3199\"}"

	data := ProductInfo{}
	if err := xjson.Unmarshal([]byte(str), &data); err != nil {
		fmt.Println("error: " + err.Error())
	} else {
		fmt.Println(data)
	}
}
