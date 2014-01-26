package admin

import (
    "fmt"
    "github.com/roydong/notes/model"
    "time"
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

        now := time.Now()
        topic := new(model.Topic)
        topic.Title = form.Title
        topic.Content = form.Content
        topic.State = form.State
        topic.UpdatedAt = now
        topic.CreatedAt = now

        if model.TopicModel.Save(topic) {
            c.Redirect(fmt.Sprintf("/topic/%d", topic.Id), 302)
        }

        form.Message = "could not save to db"
    }

RENDER:
    c.Render("admin/topic/new", form)
}

func (c *Topic) Edit() {
    id, _ := c.Request.Int64("id")
    topic := model.TopicModel.FindById(id)
    if topic == nil {
        panic("topic not found")
    }

    if c.Request.Method == "POST" {
        form := new(model.TopicForm)
        form.LoadData(c.Request)
        if l := len(form.Title); l == 0 || l > 255 {
            form.Message = "title length must between 1 - 255"
            goto RENDER
        }

        if len(form.Content) == 0 {
            form.Message = "content is empty"
            goto RENDER
        }

        topic.Title = form.Title
        topic.Content = form.Content
        topic.State = form.State
        topic.UpdatedAt = time.Now()

        if !model.TopicModel.Save(topic) {
            form.Message = "could not save to db"
        }
    }

RENDER:
    c.Render("admin/topic/edit", topic)
}

func (c *Topic) List() {
    title, _ := c.Request.String("title")
    page, _ := c.Request.Int("page")
    if page < 1 {
        page = 1
    }
    size, _ := c.Request.Int("size")
    if size < 1 {
        size = 200
    }

    c.Render("admin/topic/list", &map[string]interface{}{
        "page":     page,
        "prevpage": page - 1,
        "nextpage": page + 1,
        "size":     size,
        "topics":   model.TopicModel.Search("title", title),
    })
}
