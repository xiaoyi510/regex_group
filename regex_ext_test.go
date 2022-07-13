package regex_group

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetRexGroup(t *testing.T) {
	str := `XArr v1.2 家庭媒体管理中心
XArr v1.3 创造家庭媒体新时代
`
	ret := GetRexGroup(`XArr (?P<episode>v[\d\.]+) (\W+)`, str)
	marshal, _ := json.MarshalIndent(ret, "", "  ")
	log.Println(string(marshal))
}
