package util

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_DATA_NOT_EXIST = 10001
	ERROR_DATA_ADD       = 10002
	ERROR_DATA_DELETE    = 10003
	ERROR_DATA_EDIT      = 10004

	ERROR_USER_NAME_USED = 10011
	ERROR_CHANNEL_TITLE_USED = 10012

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
	ERROR_AUTH_NOUSER              = 20005
	ERROR_AUTH_PASSWORD            = 20006
	ERROR_AUTH_PERMISSION          = 20007
)

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	ERROR_DATA_NOT_EXIST: "数据库无此数据",
	ERROR_DATA_ADD:       "新增数据失败",
	ERROR_DATA_DELETE:    "删除数据失败",
	ERROR_DATA_EDIT:      "修改数据失败",

	ERROR_USER_NAME_USED: "用户名已存在",
	ERROR_CHANNEL_TITLE_USED: "栏目名已被使用",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	ERROR_AUTH_NOUSER:              "查无此人",
	ERROR_AUTH_PASSWORD:            "账号或密码错误",
	ERROR_AUTH_PERMISSION:          "权限不足",
}

func ErrMsg(code int) (err string) {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
