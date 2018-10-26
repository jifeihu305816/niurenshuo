package v1

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"niurenshuo/models"
	"niurenshuo/pkg/e"
	"niurenshuo/pkg/logging"
	"niurenshuo/pkg/setting"
	"niurenshuo/pkg/util"
)

// @Summary 获取评论
// @Product json
// @Param token query string true "token"
// @Param status query int false "状态 0或者1"
// @Param topic_id query int true "主题ID"
// @Param topic_type query int true "主题类型"
// @Param web_id query int true "网站id"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/comments [get]
func GetComments(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	valid := validation.Validation{}

	var status int = -1

	if arg := c.Query("status"); arg != "" {
		status = com.StrTo(arg).MustInt()
		maps["status"] = status
		valid.Range(status, 0, 1, "status").Message("状态只允许0或1")
	}

	var webId int = -1

	if arg := c.Query("web_id"); arg != "" {
		webId = com.StrTo(arg).MustInt()
		maps["web_id"] = webId
		valid.Min(webId, 1, "web_id").Message("网站ID必须大于0")
	}

	var topicId int = -1

	if arg := c.Query("topic_id"); arg != "" {
		topicId = com.StrTo(arg).MustInt()
		maps["topic_id"] = topicId
		valid.Min(topicId, 1, "topic_id").Message("主题ID必须大于0")
	}

	var topicType int = -1

	if arg := c.Query("topic_type"); arg != "" {
		topicType = com.StrTo(arg).MustInt()
		maps["topic_type"] = topicType
		valid.Min(topicType, 1, "topic_type").Message("主题类型必须大于0")
	}

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		data["lists"] = models.GetComments(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetCommentTotal(maps)
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

// @Summary 更新评论
// @Product json
// @Param id path int true "ID"
// @Param token query string true "token"
// @Param status query int false "status"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/comments/{id} [put]
func EditComment(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}

	var status int
	status = 0
	if arg := c.Query("status"); arg != "" {
		status = com.StrTo(arg).MustInt()
		valid.Range(status, 0, 1, "status").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistCommentByID(id) {
			data := make(map[string]interface{})
			data["status"] = status
			models.EditComment(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_COMMENT
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

// @Summary 新增评论
// @Accept  x-www-form-urlencoded
// @Product json
// @Param token query string true "token"
// @Param comment_id formData  int false "评论ID"
// @Param topic_id formData  int true "主题ID"
// @Param topic_type formData  int true "主题类型"
// @Param web_id formData  int true "网站ID"
// @Param from_uid formData  int true "评论用户"
// @Param to_uid formData  int true "目标用户"
// @Param content formData  string true "评论内容"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/comments [post]
func AddComment(c *gin.Context) {

	code := e.INVALID_PARAMS

	commentId := com.StrTo(c.PostForm("comment_id")).MustInt()
	topicId := com.StrTo(c.PostForm("topic_id")).MustInt()
	topicType := com.StrTo(c.PostForm("topic_type")).MustInt()
	webID := com.StrTo(c.PostForm("web_id")).MustInt()
	fromUid := com.StrTo(c.PostForm("from_uid")).MustInt()
	toUid := com.StrTo(c.PostForm("to_uid")).MustInt()

	content := c.PostForm("content")

	valid := validation.Validation{}
	valid.Required(topicId, "topic_id").Message("主题ID不能为空")
	valid.Required(webID, "web_id").Message("站点ID不能为空")
	valid.Required(fromUid, "from_uid").Message("评论uid不能为空")
	valid.Required(toUid, "to_uid").Message("评论所属uid不能为空")
	valid.Required(content, "content").Message("评论内容不能为空")
	valid.MaxSize(content, 1000, "content").Message("评论内容最长为1000字符")

	if !valid.HasErrors() {

		data := make(map[string]interface{})
		data["comment_id"] = commentId
		data["content"] = content
		data["topic_id"] = topicId
		data["topic_type"] = topicType
		data["web_id"] = webID
		data["from_uid"] = fromUid
		data["to_uid"] = toUid
		models.AddComment(data)
		code = e.SUCCESS
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// @Summary 删除评论
// @Product json
// @Param token query string true "token"
// @Param id path int true "ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/comments/{id} [delete]
func DeleteComment(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistCommentByID(id) {
			models.DeleteComment(id)
		} else {
			code = e.ERROR_NOT_EXIST_COMMENT
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
