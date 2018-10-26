package models

import (
	"github.com/jinzhu/gorm"
	"niurenshuo/pkg/logging"
)

type Comment struct {
	gorm.Model
	CommentId int    `json:"comment_id" gorm:"index:idx_cid"`
	TopicId   int    `json:"topic_id" gorm:"index:idx_web_topic"`
	TopicType int    `json:"topic_type" gorm:"index:idx_web_topic"`
	Content   string `json:"content" gorm:"size:5000"`
	FromUid   int    `json:"from_uid"`
	ToUid     int    `json:"to_uid"`
	WebId     int    `json:"web_id" gorm:"index:idx_web_topic"`
	Status    int    `json:"status"`
}

//获取评论列表
func GetComments(pageNum int, pageSize int, maps interface{}) (comments []Comment) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&comments)
	return
}

//获取评论总数
func GetCommentTotal(maps interface{}) (count int) {
	db.Model(&Comment{}).Where(maps).Count(&count)
	return
}

//添加评论
func AddComment(data map[string]interface{}) bool {

	if err := db.Create(&Comment{
		CommentId: data["comment_id"].(int),
		Content:   data["content"].(string),
		TopicId:   data["topic_id"].(int),
		TopicType: data["topic_type"].(int),
		WebId:     data["web_id"].(int),
		FromUid:   data["from_uid"].(int),
		ToUid:     data["to_uid"].(int),
		Status:    1,
	}).Error; err != nil {
		logging.Fatal(err)
	}
	return true
}

//查询评论是否存在

func ExistCommentByID(id int) bool {
	var comment Comment
	db.First(&comment, id)
	if comment.ID > 0 {
		return true
	}

	return false
}

//修改评论
func EditComment(id int, data interface{}) bool {

	db.Model(&Comment{}).Where("id = ?", id).Updates(data)

	return true
}

//软删除评论
func DeleteComment(id int) bool {
	var comment Comment
	comment.ID = uint(id)
	db.Delete(&comment)
	return true
}

//硬删除评论
func CleanAllComment() bool {
	db.Unscoped().Where("deleted_at != ? ", 0).Delete(&Comment{})
	return true
}
