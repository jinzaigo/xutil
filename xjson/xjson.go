package xjson

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
	extra.RegisterFuzzyDecoders() //开启PHP兼容模式
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func MarshalToString(v interface{}) (string, error) {
	return json.MarshalToString(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func UnmarshalFromString(str string, v interface{}) error {
	return json.UnmarshalFromString(str, v)
}
