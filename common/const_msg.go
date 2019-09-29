package common

const (
	LenUserNameMin = 6
	LenUserNameMax = 24
	LenPasswordMax = 40
	Layout         = "2006-01-02 15:04:05"
	LenDesc        = 150
	LenAddr        = 50
	LenLimit       = 10
)

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	SERVER_ERROR:					 "服务器错误",
	INVALID_PARAMS:                  "请求参数错误",
	ERROR_SEARCH_USER:               "查找用户失败",
	ERROR_PAGED_USER_FAIL:           "获取分页用户信息失败",
	ERROR_NOT_EXIST_USER:            "该用户不存在",
	ERROR_COUNT_USER_FAIL:           "获取所有用户数量失败",
	ERROR_ADD_USER_FAIL:             "添加新用户失败",
	ERROR_UPDATE_PASSWORD_FAIL:      "修改用户密码失败",
	ERROR_EDIT_USER_FAIL:            "编辑用户信息失败",
	ERROR_ADD_USER_INFO_FAIL:        "添加用户信息失败",
	ERROR_DELETE_USER_FAIL:          "删除用户失败",
	ERROR_NOT_EXIST_ARTICLE:         "该文章不存在",
	ERROR_ADD_ARTICLE_FAIL:          "新增文章失败",
	ERROR_DELETE_ARTICLE_FAIL:       "删除文章失败",
	ERROR_CHECK_EXIST_ARTICLE_FAIL:  "检查文章是否存在失败",
	ERROR_EDIT_ARTICLE_FAIL:         "修改文章失败",
	ERROR_COUNT_ARTICLE_FAIL:        "统计文章失败",
	ERROR_GET_ARTICLES_FAIL:         "获取多个文章失败",
	ERROR_GET_ARTICLE_FAIL:          "获取单个文章失败",
	ERROR_GEN_ARTICLE_POSTER_FAIL:   "生成文章海报失败",
	ERROR_GET_VERIFY_CODE_FAIL:      "验证码错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
	ERROR_AUTH_TOKEN:                "Token生成失败",
	ERROR_AUTH:                      "Token错误",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok{
		return "获取信息失败"
	}
	return msg
}