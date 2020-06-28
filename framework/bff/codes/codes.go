package codes

import "fmt"

// service inner error code list
const (
	CodeSuccess           = 0
	CodeVersion           = 1
	CodeNotFound          = 1000
	CodeRequestUrlParam   = 1001
	CodeRequestQueryParam = 1002
	CodeRequestBody       = 1004
	CodeNotLogin          = 1005
	CodeInternalServer    = 1006
	CodeNotAllow          = 1007
	CodeUnknown           = 1008
	CodeUpdating          = 1009
)

type Codes map[int]string

var (
	_codes = Codes{
		CodeSuccess:           "请求成功",
		CodeVersion:           "客户端版本错误，请升级客户端",
		CodeNotFound:          "找不到对于系统&模块",
		CodeRequestUrlParam:   "URL参数错误",
		CodeRequestQueryParam: "查询参数错误",
		CodeNotLogin:          "用户没有登录",
		CodeRequestBody:       "请求结构错误",
		CodeInternalServer:    "服务器错误",
		CodeNotAllow:          "权限校验失败",
		CodeUnknown:           "未知错误",
		CodeUpdating:          "服务正在升级",
	}
)

func GetMessage(code int) string {
	return _codes[code]
}

func CodeIter(fn func(key int, value string) bool) {
	for key, value := range _codes {
		if !fn(key, value) {
			return
		}
	}
}

func MergeCodes(codes Codes) {
	for code, describe := range codes {
		if _, exists := _codes[code]; exists {
			panic(fmt.Sprintf("code %d[%s] already exists", code, GetMessage(code)))
		}
		_codes[code] = describe
	}
}
