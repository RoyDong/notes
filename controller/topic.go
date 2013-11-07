package controller

import (
    "net/http"
    "github.com/roydong/potato"
    "github.com/roydong/notes/model"
)

type Topic struct {
    Base
}


func (c *Topic) Show() {
    id,_ := c.Request.Int("id")

    if topic := model.TopicModel.Find(id); topic != nil {
        c.Render("topic/show", topic)
    } else {
        potato.Panic(http.StatusNotFound, "topic not found")
    }
}
