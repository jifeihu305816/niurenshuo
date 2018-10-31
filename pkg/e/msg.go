package e

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "请求参数错误",
	ERROR_NOT_EXIST_COMMENT:         "该评论不存在",
	ERROR_GET_COMMENTS_FAIL:         "获取多个评论失败",
	ERROR_ADD_COMMENT_FAIL:          "添加评论失败",
	ERROR_COUNT_COMMENT_FAIL:        "统计评论失败",
	ERROR_CHECK_EXIST_COMMENT_FAIL:  "检查评论是否存在失败",
	ERROR_EDIT_COMMENT_FAIL:         "修改评论失败",
	ERROR_DELETE_COMMENT_FAIL:       "删除评论失败",
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
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
