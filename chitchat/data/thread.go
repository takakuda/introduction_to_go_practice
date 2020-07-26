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

func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")

	if err != nil {
		return
	}

	for rows.Next() {
		th := Thread{}

		if err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId, &th.CreatedAt); err != nil {
			return
		}

		threads = append(threads, th)
	}
	rows.Close()
	return
}
