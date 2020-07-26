package data

import (
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

func (thread *Thread) CreatedAtData() string {
	return thread.CreatedAt.Format("Jan 2, 2012 at 4:54pm")
}

func (post *Post) CreatedAtData() string {
	return post.CreatedAt.Format("Jan 2, 2012 at 4:54pm")
}
