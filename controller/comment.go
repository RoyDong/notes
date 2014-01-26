package controller

import (
    "github.com/roydong/notes/model"
    "time"
)

type Comment struct {
    Base
}

func (c *Comment) New() {
    r := c.Request
    if r.Method != "POST" {
        panic("post allow only")
    }

    tid, _ := r.Int64("tid")
    user, _ := r.String("user")

    topic := model.TopicModel.FindById(tid)
    if topic == nil {
        panic("topic not exists")
    }

    if len(user) > 0 {

    }

    now := time.Now()
    comment := new(model.Comment)
    comment.Content, _ = r.String("content")
    comment.UpdatedAt = now
    comment.CreatedAt = now
    comment.SetTopic(topic)

    if model.CommentModel.Save(comment) {
        c.RenderPartial("topic/_comment", comment)
    } else {
        panic("cant save to db")
    }
}
