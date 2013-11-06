package controller

import (
    "github.com/roydong/potato"
    "github.com/roydong/topic/model"
)

type Topic struct {
    Base
}

func (c *Topic) New() {
    var message string
    if c.Request.Method == "POST" {
        topic := new(model.Topic)
        topic.Title,_ = c.Request.String("title")
        if l := len(topic.Title); l == 0 || l > 255 {
            message = "title length must between 1 - 255"
            goto RENDER
        }

        topic.Content,_ = c.Request.String("content")
        if len(topic.Content) == 0 {
            message = "content is empty"
            goto RENDER
        }

        potato.L.Println(topic)

        if model.TopicModel.Save(topic) {
            c.Redirect(fmt.Sprintf("/topic/%d", topic.Id()))
            return
        }

        message = "could not save to db"
    }

    RENDER:
        c.Render("topic/new", message)
}

func (c *Topic) Show() {
    id,_ := c.Request.Int("id")

    if topic := model.TopicModel.Find(id); topic == nil {
        c.Panic(http.StatusNotFound, "topic not found")
    }

    c.Render("topic/show", topic)
}
