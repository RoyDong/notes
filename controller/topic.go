package controller

import (
    "fmt"
    "net/http"
    "github.com/roydong/potato"
    "github.com/roydong/topic/model"
)

type Topic struct {
    Base
}

func (c *Topic) New() {
    form := new(model.TopicForm)
    if c.Request.Method == "POST" {
        form.LoadData(c.Request)
        if l := len(form.Title); l == 0 || l > 255 {
            form.Message = "title length must between 1 - 255"
            goto RENDER
        }

        if len(form.Content) == 0 {
            form.Message = "content is empty"
            goto RENDER
        }

        topic := new(model.Topic)
        topic.Title = form.Title
        topic.Content = form.Content

        if model.TopicModel.Save(topic) {
            c.Redirect(fmt.Sprintf("/topic/%d", topic.Id()))
            return
        }

        form.Message = "could not save to db"
    }

    RENDER:
        c.Render("topic/new", form)
}

func (c *Topic) Show() {
    id,_ := c.Request.Int("id")

    if topic := model.TopicModel.Find(id); topic != nil {
        c.Render("topic/show", topic)
    } else {
        potato.Panic(http.StatusNotFound, "topic not found")
    }
}
