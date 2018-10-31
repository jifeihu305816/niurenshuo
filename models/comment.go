package models

import "github.com/jinzhu/gorm"

type Comment struct {
	Model
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
func GetComments(pageNum int, pageSize int, maps interface{}) ([]*Comment, error) {
	var comments []*Comment
	if err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

//获取评论总数
func GetCommentTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Comment{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

//添加评论
func AddComment(data map[string]interface{}) error {

	comment := Comment{
		CommentId: data["comment_id"].(int),
		Content:   data["content"].(string),
		TopicId:   data["topic_id"].(int),
		TopicType: data["topic_type"].(int),
		WebId:     data["web_id"].(int),
		FromUid:   data["from_uid"].(int),
		ToUid:     data["to_uid"].(int),
		Status:    data["status"].(int),
	}
	if err := db.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

//查询评论是否存在

func ExistCommentByID(id int) (bool, error) {
	var comment Comment
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&comment).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if comment.ID > 0 {
		return true, nil
	}

	return false, nil
}

//修改评论
func EditComment(id int, data interface{}) error {

	if err := db.Model(&Comment{}).Where("id = ? AND deleted_on = ?", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

//软删除评论
func DeleteComment(id int) error {
	if err := db.Where("id = ?", id).Delete(Comment{}).Error; err != nil {
		return err
	}

	return nil
}

//硬删除评论
func CleanAllComment() error {
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Comment{}).Error; err != nil {
		return err
	}
	return nil
}
