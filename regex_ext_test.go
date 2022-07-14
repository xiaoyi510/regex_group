package regex_group

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetRegxGroup(t *testing.T) {
	str := `XArr v1.2 家庭媒体管理中心
XArr v1.3 创造家庭媒体新时代
`
	//ret := GetRegxGroup(`XArr (?P<episode>v[\d\.]+) (\W+)`, str)
	ret := GetRegxGroupOne(`XArr (?P<episode>v[\d\.]+) (\W+)`, str)
	log.Println(ret.Get("episode").SubMatch)
	log.Println(ret.GetString("episode"))
	marshal, _ := json.MarshalIndent(ret, "", "  ")
	log.Println(string(marshal))
}
