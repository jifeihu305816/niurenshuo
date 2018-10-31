package cache_service

import (
	"niurenshuo/pkg/e"
	"strconv"
	"strings"
)

type Comment struct {
	ID int

	TopicId   int
	TopicType int
	WebId     int

	PageNum  int
	PageSize int
}

func (c *Comment) GetCommentKey() string {
	return e.CACHE_COMMENT + "_" + strconv.Itoa(c.ID)
}

func (c *Comment) GetCommentsKey() string {
	keys := []string{
		e.CACHE_COMMENT,
		"LIST",
	}

	if c.ID > 0 {
		keys = append(keys, strconv.Itoa(c.ID))
	}

	if c.WebId > 0 {
		keys = append(keys, strconv.Itoa(c.WebId))
	}

	if c.TopicType > 0 {
		keys = append(keys, strconv.Itoa(c.TopicType))
	}

	if c.TopicId > 0 {
		keys = append(keys, strconv.Itoa(c.TopicId))
	}

	if c.PageNum > 0 {
		keys = append(keys, strconv.Itoa(c.PageNum))
	}

	if c.PageSize > 0 {
		keys = append(keys, strconv.Itoa(c.PageSize))
	}

	return strings.Join(keys, "_")
}
