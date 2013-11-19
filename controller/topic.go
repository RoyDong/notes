package controller

import (
    "fmt"
    "net/http"
    "github.com/roydong/potato"
    "github.com/roydong/notes/model"
)

type Topic struct {
    Base
}


func (c *Topic) List() {
    q := make(map[string]string, 2)
    q["state"] = fmt.Sprintf("%d", model.TopicStatePublished)

    title,_ := c.Request.String("title")
    if len(title) > 0 { q["title"] = title }

    content,_ := c.Request.String("content")
    if len(content) > 0 { q["content"] = content }

    page,_ := c.Request.Int("page")
    if page < 1 { page = 1 }

    size,_ := c.Request.Int("size")
    if size < 1 { size = 200 }

    c.Render("topic/list", map[string]interface{} {
        "page": page,
        "prevpage": page - 1,
        "nextpage": page + 1,
        "size": size,
        "topics": model.TopicModel.Search(q),
    })
}

func (c *Topic) Show() {
    id,_ := c.Request.Int64("id")

    if topic := model.TopicModel.Find(id); topic != nil {
        c.Render("topic/show", topic)
    } else {
        potato.Panic(http.StatusNotFound, "topic not found")
    }
}
