package controller

import (
    "net/http"
    "github.com/roydong/potato"
    "github.com/roydong/notes/model"
)

type Topic struct {
    Base
}


func (c *Topic) List() {
    title,_ := c.Request.String("q")
    data := make(map[string]interface{}, 2)
    data["topics"] = model.TopicModel.SearchBy("title", title, "created_at DESC")
    data["q"] = title
    c.Render("topic/list", data)
}

func (c *Topic) Show() {
    id,_ := c.Request.Int64("id")

    if topic := model.TopicModel.FindById(id); topic != nil {
        c.Render("topic/show", topic)
    } else {
        potato.Panic(http.StatusNotFound, "topic not found")
    }
}
