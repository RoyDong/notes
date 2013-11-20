package controller

import (
    "net/http"
    "github.com/roydong/potato"
    "github.com/roydong/notes/model"
)

type Comment struct {
    Base
}

func (c *Comment) New() {
    r := c.Request
    if r.Method != "POST" {
        potato.Panic(http.StatusBadRequest, "post allow only")
    }

    tid,_ := r.Int64("tid")
    user,_ := r.String("user")

    topic := model.TopicModel.Find(tid)
    if topic == nil {
        potato.Panic(http.StatusBadRequest, "topic not exists")
    }

    if len(user) > 0 {

    }

    comment := new(model.Comment)
    comment.Content,_ = r.String("content")
    comment.SetTopic(topic)

    if model.CommentModel.Save(comment) {
        c.RenderPartial("topic/_comment", comment)
    } else {
        potato.Panic(http.StatusInternalServerError, "cant save to db")
    }
}
