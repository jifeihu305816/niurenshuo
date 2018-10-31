package v1

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"niurenshuo/pkg/app"
	"niurenshuo/pkg/e"
	"niurenshuo/pkg/setting"
	"niurenshuo/pkg/util"
	"niurenshuo/service/comment_service"
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
	appG := app.Gin{c}

	valid := validation.Validation{}

	var status int = -1
	if arg := c.Query("status"); arg != "" {
		status = com.StrTo(arg).MustInt()
		valid.Range(status, 0, 1, "status").Message("状态只允许0或1")
	}

	var webId int

	webId = com.StrTo(c.Query("web_id")).MustInt()
	valid.Min(webId, 1, "web_id").Message("网站ID必须大于0")

	var topicId int

	topicId = com.StrTo(c.Query("topic_id")).MustInt()
	valid.Min(topicId, 1, "topic_id").Message("主题ID必须大于0")

	var topicType int

	topicType = com.StrTo(c.Query("topic_type")).MustInt()
	valid.Min(topicType, 1, "topic_type").Message("主题类型必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	commentService := comment_service.Comment{
		WebId:     webId,
		TopicType: topicType,
		TopicId:   topicId,
		PageNum:   util.GetPage(c),
		PageSize:  setting.AppSetting.PageSize,
	}

	total, err := commentService.Count()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_COUNT_COMMENT_FAIL, nil)
		return
	}

	comments, err := commentService.GetAll()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_COMMENTS_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = comments
	data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// @Summary 更新评论
// @Product json
// @Param id path int true "ID"
// @Param token query string true "token"
// @Param status query int false "status"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/comments/{id} [put]
func EditComment(c *gin.Context) {

	appG := app.Gin{C: c}

	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}

	var status = 0
	if arg := c.Query("status"); arg != "" {
		status = com.StrTo(arg).MustInt()
		valid.Range(status, 0, 1, "status").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	commentService := comment_service.Comment{
		ID:     id,
		Status: status,
	}

	exists, err := commentService.ExistByID()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_CHECK_EXIST_COMMENT_FAIL, nil)
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_COMMENT, nil)
	}

	err = commentService.Edit()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_EDIT_COMMENT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
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

	appG := app.Gin{C: c}

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

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	commentService := comment_service.Comment{
		CommentId: commentId,
		Content:   content,
		TopicId:   topicId,
		TopicType: topicType,
		WebId:     webID,
		FromUid:   fromUid,
		ToUid:     toUid,
	}

	if err := commentService.Add(); err != nil {
		appG.Response(http.StatusOK, e.ERROR_ADD_COMMENT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary 删除评论
// @Product json
// @Param token query string true "token"
// @Param id path int true "ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/comments/{id} [delete]
func DeleteComment(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	commentService := comment_service.Comment{
		ID: id,
	}
	exists, err := commentService.ExistByID()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_CHECK_EXIST_COMMENT_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_COMMENT, nil)
		return
	}

	err = commentService.Delete()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_DELETE_COMMENT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
