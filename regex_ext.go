package regex_group

import (
	"regexp"
	"strconv"
)

type RexGroup map[string]RexGroupItem
type RexGroupItem struct {
	SubMatch string `json:"sub_match"`
	Offset   int    `json:"offset"`
	Size     int    `json:"size"`
}

// GetRexGroup 获取正则表达式分组内容
func GetRexGroup(regStr string, str string) []RexGroup {
	// 生成注册内容
	reg := regexp.MustCompile(regStr)

	return GetRexGroupByReg(reg, str)
}

// 通过已有正则表达式进行匹配
func GetRexGroupByReg(reg *regexp.Regexp, str string) []RexGroup {

	// 获取分组名称列表
	subexpNames := reg.SubexpNames()
	if len(subexpNames) == 0 {
		return nil
	}

	// 将空分分组名增加一个名称
	for k, v := range subexpNames {
		if v == "" {
			subexpNames[k] = "empty_" + strconv.Itoa(k)
		}
	}
	// subexpNames
	// 下标 0 为全部匹配内容
	// 其他为分组名称下标

	// 获取匹配内容
	regArr := reg.FindAllStringSubmatch(str, -1)
	if len(regArr) == 0 {
		return nil
	}
	// 获取分组匹配位置
	regIndexArr := reg.FindAllStringSubmatchIndex(str, -1)

	ret := []RexGroup{}
	for _, value := range regArr {
		temp := RexGroup{}
		for k, v := range value {
			// 循环匹配的子属性
			temp[subexpNames[k]] = RexGroupItem{
				SubMatch: v,
				Offset:   regIndexArr[0][(k)*2],
				Size:     regIndexArr[0][(k)*2+1],
			}
		}
		ret = append(ret, temp)
	}

	return ret
}
