# regex_group

Golang 简单使用正则分组命名

### 使用方式

```shell
go get -u github.com/xiaoyi510/regex_group
```

### 调用示例

```go
str := `XArr v1.2 家庭媒体管理中心
XArr v1.3 创造家庭媒体新时代
`
ret := GetRexGroup(`XArr (?P<episode>v[\d\.]+) 家庭媒体管理中心`, str)
marshal, _ := json.MarshalIndent(ret, "", "  ")
log.Println(string(marshal))
```

### 打印结果

```json
[
  {
    "empty_0": {
      "sub_match": "XArr v1.2 家庭媒体管理中心\n",
      "offset": 0,
      "size": 35
    },
    "empty_2": {
      "sub_match": "家庭媒体管理中心\n",
      "offset": 10,
      "size": 35
    },
    "episode": {
      "sub_match": "v1.2",
      "offset": 5,
      "size": 9
    }
  },
  {
    "empty_0": {
      "sub_match": "XArr v1.3 创造家庭媒体新时代\n",
      "offset": 0,
      "size": 35
    },
    "empty_2": {
      "sub_match": "创造家庭媒体新时代\n",
      "offset": 10,
      "size": 35
    },
    "episode": {
      "sub_match": "v1.3",
      "offset": 5,
      "size": 9
    }
  }
]
```