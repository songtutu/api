/*
 * @Author: song.tutu
 * @Date: 2020-12-26 09:57:13
 * @Last Modified by:   song.tutu
 * @Last Modified time: 2020-12-26 09:57:13
 */
package errmsg

const (
	SUCCSE_CODE           = 200
	ERROR_CODE            = 500
	ERROR_PARAM_NULL_CODE = 600
	//code = 100.开头代表用户表错误
	ERROR_USERNAME_USED_CODE = 1001
)

var codeMsg = map[int]string{
	SUCCSE_CODE:              "OK",
	ERROR_CODE:               "服务器撸猫去了~",
	ERROR_PARAM_NULL_CODE:    "参数不能为空",
	ERROR_USERNAME_USED_CODE: "用户名已存在!",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
