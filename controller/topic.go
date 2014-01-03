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
    data := map[string]interface{} {
        "topics": model.TopicModel.Search("title", title),
        "q": title,
    }
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
