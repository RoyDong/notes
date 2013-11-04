package controller

import (
    "github.com/roydong/potato"
    "github.com/roydong/topic/model"
)

type Base struct {
    potato.Controller
    user *model.User
}

func (c *Base) User() *model.User {
    if c.user == nil {
        c.user,_ = c.Request.Session.Value("user").(*model.User)
    }

    return c.user
}

func Init() {
    potato.R.RegControllers([]interface{}{
        new(Error),
        new(User),
        new(Main),
        new(Topic),
    })
}
