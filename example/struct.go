package main

import (
	"fmt"
	"github.com/jinzaigo/xutil"
)

type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

//注意：必须实现这个方法，才能正确调用工具包转换方法
func (p Person) GetStructData() interface{} {
	return p
}

func main() {
	m := make(map[string]interface{})
	person := Person{
		Name:    "zhangsan",
		Address: "北京海淀",
	}
	m = xutil.StructToMap(person)
	fmt.Println(m)
}
