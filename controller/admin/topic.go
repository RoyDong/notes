package admin

import (
    "fmt"
    "net/http"
    "github.com/roydong/potato"
    "github.com/roydong/notes/model"
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
        }

        form.Message = "could not save to db"
    }

    RENDER:
        c.Render("admin/topic/new", form)
}

func (c *Topic) Edit() {
    id,_ := c.Request.Int("id")
    topic := model.TopicModel.Find(id)
    if topic == nil {
        potato.Panic(http.StatusNotFound, "topic not found")
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

        if !model.TopicModel.Save(topic) {
            form.Message = "could not save to db"
        }
    }

    RENDER:
        c.Render("admin/topic/edit", topic)
}


type topicListView struct {
    Page int
    PrevPage int
    NextPage int
    Size int
    Topics []*model.Topic
}

func (c *Topic) List() {
    q := make(map[string]string, 2)
    title,_ := c.Request.String("title")
    if len(title) > 0 {
        q["title"] = title
    }

    content,_ := c.Request.String("content")
    if len(content) > 0 {
        q["content"] = content
    }

    page,_ := c.Request.Int("page")
    if page < 1 { page = 1 }
    size,_ := c.Request.Int("size")
    if size < 1 { size = 200 }

    c.Render("admin/topic/list", &topicListView{
        Page: page,
        PrevPage: page - 1,
        NextPage: page + 1,
        Size: size,
        Topics: model.TopicModel.Search(q, page, size),
    })
}
