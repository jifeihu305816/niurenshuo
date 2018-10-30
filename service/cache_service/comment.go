package cache_service

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
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
