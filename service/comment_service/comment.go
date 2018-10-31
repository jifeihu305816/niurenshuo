package comment_service

import (
	"encoding/json"
	"niurenshuo/models"
	"niurenshuo/pkg/gredis"
	"niurenshuo/pkg/logging"
	"niurenshuo/service/cache_service"
)

type Comment struct {
	ID        int
	CommentId int
	TopicId   int
	TopicType int
	Content   string
	FromUid   int
	ToUid     int
	WebId     int
	Status    int

	PageNum  int
	PageSize int
}

func (c *Comment) Add() error {
	comment := map[string]interface{}{
		"comment_id": c.CommentId,
		"topic_id":   c.TopicId,
		"topic_type": c.TopicType,
		"content":    c.Content,
		"from_uid":   c.FromUid,
		"to_uid":     c.ToUid,
		"web_id":     c.WebId,
		"status":     c.Status,
	}
	if err := models.AddComment(comment); err != nil {
		return err
	}

	return nil
}

func (c *Comment) Edit() error {
	return models.EditComment(c.ID, map[string]interface{}{
		"status": c.Status,
	})
}

func (c *Comment) GetAll() ([]*models.Comment, error) {
	var (
		comments, cacheComments []*models.Comment
	)

	cache := cache_service.Comment{
		ID: c.ID,

		TopicId:   c.TopicId,
		TopicType: c.TopicType,
		WebId:     c.WebId,

		PageNum:  c.PageNum,
		PageSize: c.PageSize,
	}

	key := cache.GetCommentsKey()

	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheComments)
			return cacheComments, nil
		}
	}

	comments, err := models.GetComments(c.PageNum, c.PageSize, c.getMaps())
	if err != nil {
		return nil, err
	}

	gredis.Set(key, comments, 3600)
	return comments, nil
}

func (c *Comment) Delete() error {
	return models.DeleteComment(c.ID)
}

func (c *Comment) ExistByID() (bool, error) {
	return models.ExistCommentByID(c.ID)
}

func (c *Comment) Count() (int, error) {
	return models.GetCommentTotal(c.getMaps())
}

func (c *Comment) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0

	if c.Status != -1 {
		maps["status"] = c.Status
	}

	maps["topic_id"] = c.TopicId
	maps["web_id"] = c.WebId
	maps["topic_type"] = c.TopicType

	return maps
}
